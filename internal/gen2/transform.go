package main

import (
	"strings"

	"github.com/iancoleman/strcase"
)

func toKebab(s string) string {
	k := strcase.ToKebab(s)
	k = strings.ReplaceAll(k, "ui-ds", "uids")
	k = strings.ReplaceAll(k, "ap-ikey", "api-key")
	return k
}

func toLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

// stripMethodSuffixes removes go-swagger method name decorations.
func stripMethodSuffixes(name string) string {
	name = strings.ReplaceAll(name, "WithParams", "")
	name = strings.ReplaceAll(name, "WithUID", "")
	name = strings.ReplaceAll(name, "calls", "")
	name = strings.ReplaceAll(name, "DataSource", "Datasource")
	name = strings.ReplaceAll(name, "GET", "")
	name = strings.ReplaceAll(name, "POST", "")
	return name
}

// methodToCommandName converts "AddTeamRoleWithParams" -> "add-team-role".
func methodToCommandName(name string) string {
	return toKebab(stripMethodSuffixes(name))
}

// methodToVarName converts "AddTeamRoleWithParams" -> "AddTeamRole".
func methodToVarName(name string) string {
	return stripMethodSuffixes(name)
}
