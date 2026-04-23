package main

import (
	"bytes"
	_ "embed"
	"flag"
	"fmt"
	"go/format"
	"log"
	"os"
	"strings"
	"text/template"
)

//go:embed gen.gotmpl
var tmpl string

func main() {
	var baseDir string
	flag.StringVar(&baseDir, "base-dir", "", "path to grafana-openapi-client-go root directory")
	flag.Parse()
	if baseDir == "" {
		log.Fatal("-base-dir is required")
	}

	// Phase 1: Parse facade to discover services
	entries, err := ParseFacade(baseDir)
	if err != nil {
		log.Fatalf("parsing facade: %v", err)
	}

	// Phase 2-5: Parse each service
	var services []*Service
	for _, entry := range entries {
		short, methods, err := ParseService(baseDir, entry.PkgName)
		if err != nil {
			log.Printf("warning: skipping service %s: %v", entry.PkgName, err)
			continue
		}
		if len(methods) == 0 {
			continue
		}

		svc := &Service{
			PkgName:   entry.PkgName,
			FieldName: entry.FieldName,
			CmdName:   toKebab(entry.PkgName),
			VarName:   toLowerCamel(entry.PkgName),
			Short:     short,
		}

		for _, m := range methods {
			act := &Action{
				CmdName:        methodToCommandName(m.Name),
				VarName:        methodToVarName(m.Name),
				GoMethodName:   m.Name,
				ParamsTypeName: m.ParamsTypeName,
				Short:          m.Short,
				Long:           m.Long,
				Response: &ResponseInfo{
					NumReturns: m.NumReturns,
				},
			}

			if m.NumReturns == 0 {
				continue
			}

			// Parse response info
			if m.ReturnTypeName != "" {
				act.Response.TypeName = m.ReturnTypeName
				resp, err := ParseResponse(baseDir, entry.PkgName, m.ReturnTypeName)
				if err != nil {
					log.Printf("warning: could not parse response for %s.%s: %v", entry.PkgName, m.Name, err)
				} else {
					act.Response.HasPayload = resp.HasPayload
					act.Response.ContainsFrame = resp.ContainsFrame
				}
			}

			// Parse params
			paramFields, err := ParseParams(baseDir, entry.PkgName, m.ParamsTypeName)
			if err != nil {
				log.Printf("warning: could not parse params for %s.%s: %v", entry.PkgName, m.Name, err)
				continue
			}

			for _, pf := range paramFields {
				if pf.IsBody {
					bfi := &BodyFieldInfo{
						ModelType:   pf.ModelType,
						IsInterface: pf.ModelType == "interface{}",
					}
					if !bfi.IsInterface && strings.HasPrefix(pf.ModelType, "models.") {
						schema, err := ParseModelSchema(baseDir, pf.ModelType)
						if err != nil {
							log.Printf("warning: could not parse body schema for %s: %v", pf.ModelType, err)
						} else {
							bfi.Schema = schema
						}
						jsonSchema, err := BuildBodyJSONSchema(baseDir, pf.ModelType)
						if err != nil {
							log.Printf("warning: could not build JSON Schema for %s: %v", pf.ModelType, err)
						} else {
							bfi.JSONSchema = jsonSchema
						}
					}
					act.BodyField = bfi
					act.Flags = append(act.Flags, &Flag{
						Name:       "body",
						FieldName:  "Body",
						Type:       "string",
						IsRequired: true,
					})
				} else {
					act.Flags = append(act.Flags, &Flag{
						Name:       toKebab(pf.FieldName),
						FieldName:  pf.FieldName,
						Type:       pf.Type,
						IsPtr:      pf.IsPtr,
						IsRequired: !pf.IsPtr,
						Doc:        pf.Doc,
					})
				}
			}

			svc.Actions = append(svc.Actions, act)
		}

		if len(svc.Actions) > 0 {
			services = append(services, svc)
		}
	}

	// Phase 6: Execute template
	funcMap := template.FuncMap{
		"flagFunc":         flagFunc,
		"defaultValue":     defaultValue,
		"flagHelp":         flagHelp,
		"stringLiteral":    stringLiteral,
		"actionLongParts":  actionLongParts,
		"actionBodySchema": actionBodySchema,
	}

	t, err := template.New("gen").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, services); err != nil {
		log.Fatalf("executing template: %v", err)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stderr.Write(buf.Bytes())
		log.Fatalf("formatting: %v", err)
	}
	os.Stdout.Write(b)
}

// flagFunc returns the cobra flag registration function name for a type.
func flagFunc(typ string, fieldName string) string {
	if fieldName == "Body" {
		return "StringVar"
	}
	switch typ {
	case "string":
		return "StringVar"
	case "int64":
		return "Int64Var"
	case "int":
		return "IntVar"
	case "bool":
		return "BoolVar"
	case "float64":
		return "Float64Var"
	case "[]string":
		return "StringSliceVar"
	case "[]int64":
		return "Int64SliceVar"
	default:
		return "StringVar"
	}
}

// defaultValue returns the default value literal for a flag type.
func defaultValue(typ string, fieldName string) string {
	if fieldName == "Body" {
		return `""`
	}
	if fieldName == "Perpage" {
		return "1000"
	}
	switch typ {
	case "string":
		return `""`
	case "int64", "int":
		return "0"
	case "bool":
		return "false"
	case "float64":
		return "0.0"
	case "[]string":
		return "[]string{}"
	case "[]int64":
		return "[]int64{}"
	default:
		return `""`
	}
}

// flagHelp returns a quoted help string for a flag.
// If doc is non-empty it takes precedence over the default.
func flagHelp(fieldName string, doc string) string {
	if fieldName == "Body" {
		return `"The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{\"foo\": \"bar\"}'"` //nolint:goconst
	}
	if doc != "" {
		return fmt.Sprintf(`%q`, doc)
	}
	return fmt.Sprintf(`%q`, fieldName)
}

// stringLiteral returns a readable Go string literal for generated help text.
// Multi-line strings use a raw string literal; embedded backticks are emitted
// as "`" + concatenated segments so the multi-line form is preserved.
func stringLiteral(s string) string {
	if !strings.Contains(s, "\n") {
		return fmt.Sprintf("%q", s)
	}
	if !strings.Contains(s, "`") {
		return "`" + s + "`"
	}
	var b strings.Builder
	for i, part := range strings.Split(s, "`") {
		if i > 0 {
			b.WriteString(` + "` + "`" + `" + `)
		}
		b.WriteByte('`')
		b.WriteString(part)
		b.WriteByte('`')
	}
	return b.String()
}

func actionLongParts(act *Action) []string {
	var longParts []string
	if act.Long != "" {
		longParts = append(longParts, strings.Split(act.Long, "\n\n")...)
	}
	if len(longParts) == 0 {
		return nil
	}
	if act.Short == "" {
		return longParts
	}
	return append([]string{act.Short}, longParts...)
}

func actionBodySchema(act *Action) string {
	if act.BodyField == nil || act.BodyField.Schema == nil {
		return ""
	}
	return formatBodySchema(act.BodyField.Schema)
}
