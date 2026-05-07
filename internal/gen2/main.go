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
	"unicode"
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
				HTTPMethod:     m.HTTPMethod,
				HTTPPath:       m.HTTPPath,
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
						In:         "body",
					})
				} else {
					// Slice types ignore parsed defaults (defaultValue cannot
					// splice an array literal in verbatim), so don't carry one
					// into Flag.Default either — keeps cobra's actual default
					// and help.json's advertised default in agreement.
					def := pf.Default
					if pf.Type == "[]string" || pf.Type == "[]int64" {
						def = ""
					}
					act.Flags = append(act.Flags, &Flag{
						Name:       toKebab(pf.FieldName),
						FieldName:  pf.FieldName,
						Type:       pf.Type,
						IsPtr:      pf.IsPtr,
						IsRequired: pf.IsRequired,
						Doc:        pf.Doc,
						Default:    def,
						In:         pf.In,
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

// defaultValue returns the default value literal for a flag type. When parsed
// is non-empty, it is the Go-literal default extracted from the source doc
// comment (e.g. `1000`, `"View"`, `true`) and is preferred over the type-based
// zero value, so the cobra flag's default matches what the doc advertises.
//
// Slice types intentionally ignore parsed: go-swagger does not emit array
// defaults in the form we can splice in verbatim, so falling back to the empty
// slice keeps the generated code compilable.
func defaultValue(typ string, fieldName string, parsed string) string {
	if fieldName == "Body" {
		return `""`
	}
	if parsed != "" && typ != "[]string" && typ != "[]int64" {
		return parsed
	}
	// Perpage fallback: at least one upstream parameter
	// (service_accounts.SearchOrgServiceAccountsWithPagingParams) describes its
	// default in prose ("The default value is 1000.") instead of emitting a
	// structured `Default: 1000` line, so the parser cannot recover it. The
	// Grafana convention for Perpage is 1000, so default to that when no
	// literal was extracted to avoid regressing those commands to 0.
	if fieldName == "Perpage" && (typ == "int64" || typ == "int") {
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
// When isRequired is true, " [required]" is appended so the textual --help
// output distinguishes required flags (cobra/pflag does not show this by default).
func flagHelp(fieldName string, doc string, isRequired bool) string {
	var s string
	switch {
	case fieldName == "Body":
		s = `Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.`
	case doc != "":
		s = doc
	default:
		s = fieldName
	}
	if isRequired {
		s += " [required]"
	}
	return fmt.Sprintf(`%q`, s)
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
	if nearlyDuplicateHelp(act.Short, longParts[0]) {
		return longParts
	}
	return append([]string{act.Short}, longParts...)
}

// nearlyDuplicateHelp reports whether a and b carry essentially the same help
// content. Catches two flavours of upstream redundancy: Short truncated as a
// prefix of the first Long paragraph, and OpenAPI summary/description pairs
// that differ by an interior synonym ("new user" vs "global user").
//
// The Jaccard threshold was calibrated against a sweep of internal/help.json:
// it accepts every clear duplicate down to 0.78 and rejects the closest
// genuine non-duplicate at 0.54 (signed-in-user get-signed-in-user-team-list).
// One known borderline case (library-elements create-library-element, 0.60)
// stays un-deduped; lowering further would start eating real content.
const helpDupJaccardThreshold = 0.7

func nearlyDuplicateHelp(a, b string) bool {
	na, nb := normalizeHelpText(a), normalizeHelpText(b)
	if na == "" || nb == "" {
		return false
	}
	// Pad with spaces so substring containment respects word boundaries
	// (otherwise "set up" would match inside "reset upstream").
	pa, pb := " "+na+" ", " "+nb+" "
	if strings.Contains(pa, pb) || strings.Contains(pb, pa) {
		return true
	}
	return jaccardWords(na, nb) >= helpDupJaccardThreshold
}

func normalizeHelpText(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	prevSpace := true
	for _, r := range s {
		switch {
		case unicode.IsSpace(r), unicode.IsPunct(r):
			if !prevSpace {
				b.WriteByte(' ')
				prevSpace = true
			}
		default:
			b.WriteRune(unicode.ToLower(r))
			prevSpace = false
		}
	}
	return strings.TrimSpace(b.String())
}

func jaccardWords(a, b string) float64 {
	set := func(s string) map[string]struct{} {
		m := make(map[string]struct{})
		for w := range strings.FieldsSeq(s) {
			m[w] = struct{}{}
		}
		return m
	}
	A, B := set(a), set(b)
	if len(A) == 0 || len(B) == 0 {
		return 0
	}
	inter := 0
	for w := range A {
		if _, ok := B[w]; ok {
			inter++
		}
	}
	union := len(A) + len(B) - inter
	return float64(inter) / float64(union)
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
//
// CommonFlags lists the universal flags that gen.gotmpl wires onto every
// (or schema-bearing) action; per-action Flags only carries OpenAPI params,
// so without this block an agent reading --help-json would never see --raw
// or the --describe-*-jsonschema flags. AppliesWhen ties each common flag
// to a presence rule the agent can check against the per-action JSON:
//   - "always":   registered on every command.
//   - "body":     registered iff command.body != null && command.body.hasJSONSchema.
//   - "response": registered iff command.response != null && command.response.hasJSONSchema.
type helpDoc struct {
	Version     string                       `json:"version"`
	CommonFlags []helpCommonFlag             `json:"commonFlags"`
	Commands    map[string]*helpActionOutput `json:"commands"`
}

type helpActionOutput struct {
	Service    string              `json:"service"`
	Action     string              `json:"action"`
	Short      string              `json:"short,omitempty"`
	Long       string              `json:"long,omitempty"`
	HTTPMethod string              `json:"httpMethod,omitempty"`
	HTTPPath   string              `json:"httpPath,omitempty"`
	Flags      []helpFlagOutput    `json:"flags"`
	Body       *helpBodyOutput     `json:"body,omitempty"`
	Response   *helpResponseOutput `json:"response,omitempty"`
}

type helpFlagOutput struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Required bool   `json:"required"`
	// In is the OpenAPI parameter location ("path", "query", "header", "form",
	// "file"). Body fields are surfaced via the top-level "body" object and
	// excluded from the per-action Flags, so "body" never appears here.
	In  string `json:"in,omitempty"`
	Doc string `json:"doc,omitempty"`
	// Default is the Go-literal default extracted from the source doc comment.
	// Strings are JSON-quoted (e.g. "\"View\""), ints/bools/floats are bare
	// (e.g. "1000", "true"). Empty when the upstream comment had no Default:.
	Default string `json:"default,omitempty"`
}

// helpCommonFlag describes a flag wired onto commands by gen.gotmpl rather
// than by the OpenAPI param model. AppliesWhen is one of "always", "body",
// "response" — see helpDoc's comment for the agent contract. DefaultRule,
// when set, is a JSON-path-style expression evaluated against the per-action
// block ("response.containsFrame" → bool); when empty, the default is false.
// When DefaultRule fires (e.g. response.containsFrame is true and --raw flips
// default-on), the per-action response block also carries a human-readable
// "note" explaining why; agents should surface that string rather than
// re-deriving the wording.
type helpCommonFlag struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Doc         string `json:"doc,omitempty"`
	AppliesWhen string `json:"appliesWhen"`
	DefaultRule string `json:"defaultRule,omitempty"`
}

type helpBodyOutput struct {
	ModelType string `json:"modelType"`
	// HasJSONSchema is true iff --describe-body-jsonschema is registered for
	// this action. The flag is gated on a non-empty BodyField.JSONSchema, which
	// is empty for interface{}/any-typed bodies — body block is still emitted
	// in that case so agents can see modelType, but the flag is absent.
	HasJSONSchema bool `json:"hasJSONSchema"`
}

type helpResponseOutput struct {
	TypeName      string `json:"typeName,omitempty"`
	HasPayload    bool   `json:"hasPayload"`
	ContainsFrame bool   `json:"containsFrame"`
	// Note mirrors the schema's top-level "description" (set by
	// BuildJSONSchemaFromExpr when ContainsFrame is true) so JSON consumers see
	// the same advisory the human --help renders above the field table. 1:1
	// with ContainsFrame today; sourced from Response.RootDescription so the
	// two surfaces stay in sync from a single field.
	Note string `json:"note,omitempty"`
	// HasJSONSchema is true iff --describe-response-jsonschema is registered
	// for this action. Mirrors HasPayload in practice but tracks the flag's
	// real gate (Response.JSONSchema != "") in case schema construction failed.
	HasJSONSchema bool `json:"hasJSONSchema"`
}

// commonFlags is the canonical descriptor for the universal flags wired in
// gen.gotmpl. MUST stay in sync with the BoolVar registrations in
// gen.gotmpl (the --describe-*-jsonschema and --raw lines). Adding a flag
// here without adding it there (or vice versa) will mislead agent consumers.
var commonFlags = []helpCommonFlag{
	{
		Name:        "raw",
		Type:        "bool",
		Doc:         "Print the raw HTTP response body instead of the decoded payload",
		AppliesWhen: "always",
		DefaultRule: "response.containsFrame",
	},
	{
		Name:        "describe-body-jsonschema",
		Type:        "bool",
		Doc:         "Print the JSON Schema of the request body and exit without calling the API",
		AppliesWhen: "body",
	},
	{
		Name:        "describe-response-jsonschema",
		Type:        "bool",
		Doc:         "Print the JSON Schema of the response payload and exit without calling the API",
		AppliesWhen: "response",
	},
}

// buildHelpJSON returns an indented JSON document summarising the whole CLI
// command tree. Indented form keeps the embedded help.json diff-friendly
// (one JSON line per source line); the runtime calls json.Compact before
// printing so --help-json output is single-line.
func buildHelpJSON(services []*Service) (string, error) {
	doc := &helpDoc{
		Version:     "gf-help-json/1",
		CommonFlags: commonFlags,
		Commands:    make(map[string]*helpActionOutput, 256),
	}
	for _, svc := range services {
		for _, act := range svc.Actions {
			key := svc.CmdName + " " + act.CmdName
			out := &helpActionOutput{
				Service:    svc.CmdName,
				Action:     act.CmdName,
				Short:      act.Short,
				Long:       act.Long,
				HTTPMethod: act.HTTPMethod,
				HTTPPath:   act.HTTPPath,
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
					In:       fl.In,
					Doc:      fl.Doc,
					Default:  fl.Default,
				})
			}
			if out.Flags == nil {
				out.Flags = []helpFlagOutput{}
			}
			if act.BodyField != nil {
				out.Body = &helpBodyOutput{
					ModelType:     act.BodyField.ModelType,
					HasJSONSchema: act.BodyField.JSONSchema != "",
				}
			}
			if act.Response != nil && act.Response.HasPayload {
				out.Response = &helpResponseOutput{
					TypeName:      act.Response.TypeName,
					HasPayload:    true,
					ContainsFrame: act.Response.ContainsFrame,
					Note:          act.Response.RootDescription,
					HasJSONSchema: act.Response.JSONSchema != "",
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
