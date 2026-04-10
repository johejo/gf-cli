package main

type Service struct {
	PkgName   string // "access_control"
	FieldName string // "AccessControl" (field name on GrafanaHTTPAPI)
	CmdName   string // "access-control"
	VarName   string // "accessControl" (variable name prefix in generated code)
	Actions   []*Action
}

type Action struct {
	CmdName        string        // "add-team-role"
	VarName        string        // "AddTeamRole" (variable name suffix)
	GoMethodName   string        // "AddTeamRoleWithParams"
	ParamsTypeName string        // "AddTeamRoleParams"
	Short          string        // short help text from source doc comment
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
	ModelType   string // "models.AddTeamRoleCommand"
	IsInterface bool   // true if ModelType is "interface{}"
}

type Flag struct {
	Name       string // "team-id" (kebab-case CLI flag name)
	FieldName  string // "TeamID" (field name on Params struct)
	Type       string // "string", "int64", "bool", "[]string", "[]int64"
	IsPtr      bool
	IsRequired bool
	Doc        string // help text from source doc comment
}
