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
	Response       *ResponseInfo // nil if method returns only error
	BodyField      *BodyFieldInfo
	Flags          []*Flag
}

type ResponseInfo struct {
	TypeName   string // "AddTeamRoleOK"
	HasPayload bool
	NumReturns int // 1=error, 2=resp+error, 3=resp+extra+error
}

type BodyFieldInfo struct {
	ModelType   string          // "models.AddTeamRoleCommand"
	IsInterface bool            // true if ModelType is "interface{}"
	Schema      *BodySchemaInfo // nil for interface{} bodies
	JSONSchema  string          // pretty-printed JSON Schema (draft 2020-12); empty when unavailable
}

type BodySchemaInfo struct {
	TypeName string        // e.g. "CreateTeamCommand"
	Fields   []*ModelField
}

type ModelField struct {
	JSONName     string        // from json tag, e.g. "email"
	GoType       string        // display type, e.g. "string", "[]Permission"
	JSONType     string        // resolved JSON type: "string", "number", "boolean", "any", "object"
	IsRequired   bool          // from "// Required: true" comment
	IsArray      bool          // true if field is a slice/array
	IsMap        bool          // true if field is map[K]V
	MapValueType string        // resolved JSON type for map values
	EnumValues   []string      // from "// Enum: [val1 val2]" comment
	NestedFields []*ModelField // 1-level expansion for struct types
}

type Flag struct {
	Name       string // "team-id" (kebab-case CLI flag name)
	FieldName  string // "TeamID" (field name on Params struct)
	Type       string // "string", "int64", "bool", "[]string", "[]int64"
	IsPtr      bool
	IsRequired bool
	Doc        string // help text from source doc comment
}
