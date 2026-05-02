package main

type Service struct {
	PkgName   string // "access_control"
	FieldName string // "AccessControl" (field name on GrafanaHTTPAPI)
	CmdName   string // "access-control"
	VarName   string // "accessControl" (variable name prefix in generated code)
	Short     string // single-line summary from the Client struct doc comment
	Actions   []*Action
}

type Action struct {
	CmdName        string        // "add-team-role"
	VarName        string        // "AddTeamRole" (variable name suffix)
	GoMethodName   string        // "AddTeamRoleWithParams"
	ParamsTypeName string        // "AddTeamRoleParams"
	Short          string        // single-line summary from source doc comment
	Long           string        // detailed help text from source doc comment
	HTTPMethod     string        // "GET", "POST", ... from the WithParams runtime.ClientOperation; "" if not found
	HTTPPath       string        // OpenAPI path pattern, e.g. "/admin/ldap/{user_name}"; "" if not found
	Response       *ResponseInfo // nil if method returns only error
	BodyField      *BodyFieldInfo
	Flags          []*Flag
}

type ResponseInfo struct {
	TypeName        string // "AddTeamRoleOK"
	HasPayload      bool
	NumReturns      int    // 1=error, 2=resp+error, 3=resp+extra+error
	ContainsFrame   bool   // true if the payload type tree contains models.Frame (wire-format mismatch)
	JSONSchema      string // pretty-printed JSON Schema (draft 2020-12) for the Payload; empty when unavailable
	Annotations     string // typed field table mirroring JSONSchema; empty when the schema has no walkable properties
	RootDescription string // schema's top-level description; carries the Frame-mismatch advisory when ContainsFrame is true
}

type BodyFieldInfo struct {
	ModelType   string // "models.AddTeamRoleCommand"
	IsInterface bool   // true if ModelType is "interface{}"
	JSONSchema  string // pretty-printed JSON Schema (draft 2020-12); empty when unavailable
	Annotations string // typed field table mirroring JSONSchema; empty when the schema has no walkable properties
}

type Flag struct {
	Name       string // "team-id" (kebab-case CLI flag name)
	FieldName  string // "TeamID" (field name on Params struct)
	Type       string // "string", "int64", "bool", "[]string", "[]int64"
	IsPtr      bool
	IsRequired bool
	Doc        string // help text from source doc comment
	Default    string // Go-literal default extracted from the doc comment; empty if absent
	In         string // OpenAPI parameter location: "path", "query", "header", "body", "form", "file"
}
