package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
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
	var (
		baseDir     string
		outGo       string
		outHelpJSON string
	)
	flag.StringVar(&baseDir, "base-dir", "", "path to grafana-openapi-client-go root directory")
	flag.StringVar(&outGo, "out-go", "", "destination path for the generated Go file")
	flag.StringVar(&outHelpJSON, "out-help-json", "", "destination path for the help JSON document")
	flag.Parse()
	if baseDir == "" {
		log.Fatal("-base-dir is required")
	}
	if outGo == "" {
		log.Fatal("-out-go is required")
	}
	if outHelpJSON == "" {
		log.Fatal("-out-help-json is required")
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
					if resp.HasPayload && resp.PayloadExpr != nil {
						title := m.ReturnTypeName + ".Payload"
						jsonSchema, annotations, rootDesc, err := BuildJSONSchemaFromExpr(baseDir, resp.PayloadExpr, title, resp.ContainsFrame)
						if err != nil {
							log.Printf("warning: could not build response JSON Schema for %s.%s: %v", entry.PkgName, m.Name, err)
						} else {
							act.Response.JSONSchema = jsonSchema
							act.Response.Annotations = annotations
							act.Response.RootDescription = rootDesc
						}
					}
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
						IsInterface: pf.ModelType == "interface{}" || pf.ModelType == "any",
					}
					if !bfi.IsInterface && strings.HasPrefix(pf.ModelType, "models.") {
						jsonSchema, annotations, err := BuildBodyJSONSchema(baseDir, pf.ModelType)
						if err != nil {
							log.Printf("warning: could not build JSON Schema for %s: %v", pf.ModelType, err)
						} else {
							bfi.JSONSchema = jsonSchema
							bfi.Annotations = annotations
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

	// Build the --help-json payload: a flat map keyed by "<service> <action>"
	// so agents can do O(1) lookup without recursing a tree.
	helpJSON, err := buildHelpJSON(services)
	if err != nil {
		log.Fatalf("building help JSON: %v", err)
	}

	// Phase 6: Execute template
	funcMap := template.FuncMap{
		"flagFunc":             flagFunc,
		"defaultValue":         defaultValue,
		"flagHelp":             flagHelp,
		"stringLiteral":        stringLiteral,
		"actionLongParts":      actionLongParts,
		"actionBodySchema":     actionBodySchema,
		"actionResponseSchema": actionResponseSchema,
	}

	t, err := template.New("gen").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		log.Fatalf("parsing template: %v", err)
	}

	data := struct {
		Services []*Service
	}{
		Services: services,
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		log.Fatalf("executing template: %v", err)
	}

	b, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stderr.Write(buf.Bytes())
		log.Fatalf("formatting: %v", err)
	}
	if err := os.WriteFile(outGo, b, 0o644); err != nil {
		log.Fatalf("writing %s: %v", outGo, err)
	}
	if err := os.WriteFile(outHelpJSON, []byte(helpJSON), 0o644); err != nil {
		log.Fatalf("writing %s: %v", outHelpJSON, err)
	}
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
		return `"Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{\"foo\": \"bar\"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema."`
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

// actionBodySchema returns the rendered "Body schema" help block for an action,
// or "" when the action has no body or its body schema could not be derived.
// The block is a header line followed by a typed field table; the strict
// Draft 2020-12 JSON Schema is available via --describe-body-jsonschema.
func actionBodySchema(act *Action) string {
	if act == nil || act.BodyField == nil || act.BodyField.JSONSchema == "" {
		return ""
	}
	if act.BodyField.Annotations == "" {
		return ""
	}
	typeName := strings.TrimPrefix(act.BodyField.ModelType, "models.")
	return "Body schema (" + typeName + "):\n" + act.BodyField.Annotations
}

// actionResponseSchema returns the rendered "Response schema" help block for an
// action, or "" when the response has no payload schema. The block is a header
// line, an optional "Note:" line carrying the schema's top-level description
// (used today only for the Frame-mismatch advisory), and a typed field table.
// The strict JSON Schema is available via --describe-response-jsonschema.
func actionResponseSchema(act *Action) string {
	if act == nil || act.Response == nil || act.Response.JSONSchema == "" {
		return ""
	}
	if act.Response.Annotations == "" && act.Response.RootDescription == "" {
		return ""
	}
	out := "Response schema (" + act.Response.TypeName + ".Payload):"
	if act.Response.RootDescription != "" {
		out += "\n  Note: " + act.Response.RootDescription
	}
	if act.Response.Annotations != "" {
		out += "\n" + act.Response.Annotations
	}
	return out
}

// helpDoc is the top-level shape of the --help-json output. It is a discovery
// index, not a full schema dump: body/response carry the model type names but
// not the JSON Schemas, which are available per-command via the
// --describe-body-jsonschema / --describe-response-jsonschema flags. The map
// is flat-keyed by "<service> <action>" (the invocation string minus "gf ")
// so agents can do O(1) lookup with .commands[name] instead of walking a tree.
type helpDoc struct {
	Version  string                       `json:"version"`
	Commands map[string]*helpActionOutput `json:"commands"`
}

type helpActionOutput struct {
	Service  string              `json:"service"`
	Action   string              `json:"action"`
	Short    string              `json:"short,omitempty"`
	Long     string              `json:"long,omitempty"`
	Flags    []helpFlagOutput    `json:"flags"`
	Body     *helpBodyOutput     `json:"body,omitempty"`
	Response *helpResponseOutput `json:"response,omitempty"`
}

type helpFlagOutput struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	Doc      string `json:"doc,omitempty"`
}

type helpBodyOutput struct {
	ModelType string `json:"modelType"`
}

type helpResponseOutput struct {
	TypeName      string `json:"typeName,omitempty"`
	HasPayload    bool   `json:"hasPayload"`
	ContainsFrame bool   `json:"containsFrame"`
}

// buildHelpJSON returns an indented JSON document summarising the whole CLI
// command tree. Indented form keeps the embedded help.json diff-friendly
// (one JSON line per source line); the runtime calls json.Compact before
// printing so --help-json output is single-line.
func buildHelpJSON(services []*Service) (string, error) {
	doc := &helpDoc{
		Version:  "gf-help-json/1",
		Commands: make(map[string]*helpActionOutput, 256),
	}
	for _, svc := range services {
		for _, act := range svc.Actions {
			key := svc.CmdName + " " + act.CmdName
			out := &helpActionOutput{
				Service: svc.CmdName,
				Action:  act.CmdName,
				Short:   act.Short,
				Long:    act.Long,
			}
			for _, fl := range act.Flags {
				// The synthetic "body" flag doesn't carry useful doc text for
				// agents; skip it here — body info lives under Body below.
				if fl.FieldName == "Body" {
					continue
				}
				out.Flags = append(out.Flags, helpFlagOutput{
					Name:     fl.Name,
					Type:     fl.Type,
					Required: fl.IsRequired,
					Doc:      fl.Doc,
				})
			}
			if out.Flags == nil {
				out.Flags = []helpFlagOutput{}
			}
			if act.BodyField != nil {
				out.Body = &helpBodyOutput{ModelType: act.BodyField.ModelType}
			}
			if act.Response != nil && act.Response.HasPayload {
				out.Response = &helpResponseOutput{
					TypeName:      act.Response.TypeName,
					HasPayload:    true,
					ContainsFrame: act.Response.ContainsFrame,
				}
			}
			doc.Commands[key] = out
		}
	}
	b, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
