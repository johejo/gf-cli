package main

import (
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"path"
	"reflect"
	"strings"
	"sync"
	"text/template"

	"github.com/gertd/go-pluralize"
	"github.com/grafana/grafana-openapi-client-go/client"
	"github.com/grafana/grafana-openapi-client-go/client/access_control"
	"github.com/grafana/grafana-openapi-client-go/client/access_control_provisioning"
	"github.com/grafana/grafana-openapi-client-go/client/admin"
	"github.com/grafana/grafana-openapi-client-go/client/admin_ldap"
	"github.com/grafana/grafana-openapi-client-go/client/admin_provisioning"
	"github.com/grafana/grafana-openapi-client-go/client/admin_users"
	"github.com/grafana/grafana-openapi-client-go/client/annotations"
	"github.com/grafana/grafana-openapi-client-go/client/api_keys"
	"github.com/grafana/grafana-openapi-client-go/client/correlations"
	"github.com/grafana/grafana-openapi-client-go/client/dashboard_permissions"
	"github.com/grafana/grafana-openapi-client-go/client/dashboard_public"
	"github.com/grafana/grafana-openapi-client-go/client/dashboard_versions"
	"github.com/grafana/grafana-openapi-client-go/client/dashboards"
	"github.com/grafana/grafana-openapi-client-go/client/datasources"
	"github.com/grafana/grafana-openapi-client-go/client/devices"
	"github.com/grafana/grafana-openapi-client-go/client/ds"
	"github.com/grafana/grafana-openapi-client-go/client/enterprise"
	"github.com/grafana/grafana-openapi-client-go/client/folder_permissions"
	"github.com/grafana/grafana-openapi-client-go/client/folders"
	"github.com/grafana/grafana-openapi-client-go/client/get_current_org"
	"github.com/grafana/grafana-openapi-client-go/client/health"
	"github.com/grafana/grafana-openapi-client-go/client/ldap_debug"
	"github.com/grafana/grafana-openapi-client-go/client/library_elements"
	"github.com/grafana/grafana-openapi-client-go/client/licensing"
	"github.com/grafana/grafana-openapi-client-go/client/migrations"
	"github.com/grafana/grafana-openapi-client-go/client/org"
	"github.com/grafana/grafana-openapi-client-go/client/org_invites"
	"github.com/grafana/grafana-openapi-client-go/client/org_preferences"
	"github.com/grafana/grafana-openapi-client-go/client/orgs"
	"github.com/grafana/grafana-openapi-client-go/client/playlists"
	"github.com/grafana/grafana-openapi-client-go/client/provisioning"
	"github.com/grafana/grafana-openapi-client-go/client/query_history"
	"github.com/grafana/grafana-openapi-client-go/client/recording_rules"
	"github.com/grafana/grafana-openapi-client-go/client/reports"
	"github.com/grafana/grafana-openapi-client-go/client/saml"
	"github.com/grafana/grafana-openapi-client-go/client/search"
	"github.com/grafana/grafana-openapi-client-go/client/service_accounts"
	"github.com/grafana/grafana-openapi-client-go/client/signed_in_user"
	"github.com/grafana/grafana-openapi-client-go/client/signing_keys"
	"github.com/grafana/grafana-openapi-client-go/client/snapshots"
	"github.com/grafana/grafana-openapi-client-go/client/sso_settings"
	"github.com/grafana/grafana-openapi-client-go/client/sync_team_groups"
	"github.com/grafana/grafana-openapi-client-go/client/teams"
	"github.com/grafana/grafana-openapi-client-go/client/user"
	"github.com/grafana/grafana-openapi-client-go/client/user_preferences"
	"github.com/grafana/grafana-openapi-client-go/client/users"
	"github.com/iancoleman/strcase"
)

var (
	typeList = []reflect.Type{
		getType[access_control.ClientService](),
		getType[access_control_provisioning.ClientService](),
		getType[admin.ClientService](),
		getType[admin_ldap.ClientService](),
		getType[admin_provisioning.ClientService](),
		getType[admin_users.ClientService](),
		getType[annotations.ClientService](),
		getType[api_keys.ClientService](),
		getType[correlations.ClientService](),
		getType[dashboard_permissions.ClientService](),
		getType[dashboard_public.ClientService](),
		getType[dashboard_versions.ClientService](),
		getType[dashboards.ClientService](),
		getType[datasources.ClientService](),
		getType[devices.ClientService](),
		getType[ds.ClientService](),
		getType[enterprise.ClientService](),
		getType[folder_permissions.ClientService](),
		getType[folders.ClientService](),
		getType[get_current_org.ClientService](),
		getType[health.ClientService](),
		getType[ldap_debug.ClientService](),
		getType[library_elements.ClientService](),
		getType[licensing.ClientService](),
		getType[migrations.ClientService](),
		getType[org.ClientService](),
		getType[org_invites.ClientService](),
		getType[org_preferences.ClientService](),
		getType[orgs.ClientService](),
		getType[playlists.ClientService](),
		getType[provisioning.ClientService](),
		getType[query_history.ClientService](),
		getType[recording_rules.ClientService](),
		getType[reports.ClientService](),
		getType[saml.ClientService](),
		getType[search.ClientService](),
		getType[service_accounts.ClientService](),
		getType[signed_in_user.ClientService](),
		getType[signing_keys.ClientService](),
		getType[snapshots.ClientService](),
		getType[sso_settings.ClientService](),
		getType[sync_team_groups.ClientService](),
		getType[teams.ClientService](),
		getType[user.ClientService](),
		getType[user_preferences.ClientService](),
		getType[users.ClientService](),
	}
)

func getType[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}

func toKebab(s string) string {
	k := strcase.ToKebab(s)
	// workaround
	k = strings.ReplaceAll(k, "ui-ds", "uids")
	k = strings.ReplaceAll(k, "ap-ikey", "api-key")
	return k
}

func toSingular(s string) string {
	return sync.OnceValue(pluralize.NewClient)().Singular(s)
}

var (
	funcMap = map[string]any{
		"toCamel":    strcase.ToCamel,
		"toKebab":    toKebab,
		"toSingular": toSingular,
		"hasSuffix":  strings.HasSuffix,
		"zeroValue": func(v string) (string, error) {
			switch v {
			case "bool":
				return `false`, nil
			case "string":
				return `""`, nil
			case "int", "int64":
				return `0`, nil
			case "stringSlice":
				return `[]string{}`, nil
			case "int64Slice":
				return `[]int64{}`, nil
			}
			return "", fmt.Errorf("unexpected argument %s", v)
		},
		"formatGoType": func(v string) string {
			if s, ok := strings.CutSuffix(v, "Slice"); ok {
				return "[]" + s
			}
			return v
		},
	}
)

func main() {
	apiType := getType[client.GrafanaHTTPAPI]()
	apiFieldMap := make(map[string]string)
	for i := range apiType.NumField() {
		f := apiType.Field(i)
		if !f.IsExported() {
			continue
		}
		pkgPath := f.Type.PkgPath()
		apiFieldMap[pkgPath] = f.Name
	}
	var services []*Service
	for _, t := range typeList {
		var actions []*Action
		for i := range t.NumMethod() {
			m := t.Method(i)
			if !strings.HasSuffix(m.Type.In(0).String(), "Params") {
				continue
			}
			// skip deprecated API
			if strings.Contains(m.Name, "ByID") {
				continue
			}
			// skip datasource proxy API
			if strings.Contains(m.Name, "DatasourceProxy") {
				continue
			}

			cmdName := m.Name
			cmdName = strings.ReplaceAll(cmdName, "WithParams", "")
			cmdName = strings.ReplaceAll(cmdName, "WithUID", "")
			cmdName = strings.ReplaceAll(cmdName, "calls", "")
			cmdName = strings.ReplaceAll(cmdName, "DataSource", "Datasource")
			cmdName = strings.ReplaceAll(cmdName, "GET", "")
			cmdName = strings.ReplaceAll(cmdName, "POST", "")
			actionName := cmdName
			cmdName = toKebab(cmdName)
			var flgs []*Flag
			for i := range m.Type.NumIn() {
				in := m.Type.In(i)
				switch in.String() {
				case "[]dashboards.ClientOption":
					continue
				}
				flgs = append(flgs, getFlags(cmdName, in)...)
			}
			numOut := m.Type.NumOut()
			out0 := m.Type.Out(0)
			if out0.Kind() == reflect.Ptr {
				out0 = out0.Elem()
			}
			hasGetPaylobd := false
			if out0.Kind() == reflect.Struct {
				if out0.NumField() > 0 && out0.Field(0).Name == "Payload" {
					hasGetPaylobd = true
				}
			}
			var bodyTypeName string
			for _, flg := range flgs {
				if flg.GoParamName == "Body" {
					bodyTypeName = strings.TrimPrefix(flg.Type, "*")
				}
			}
			actions = append(actions, &Action{
				CmdName:         cmdName,
				ActionName:      actionName,
				GoMethodName:    m.Name,
				GoParamName:     apiFieldMap[m.Type.In(0).Elem().PkgPath()],
				NumOut:          numOut,
				GoParamTypeName: m.Type.In(0).Elem().Name(),
				HasGetPayload:   hasGetPaylobd,
				BodyTypeName:    bodyTypeName,
				Flags:           flgs,
			})
		}
		pkg := path.Base(t.PkgPath())
		services = append(services, &Service{
			Name:    strcase.ToLowerCamel(pkg),
			CmdName: toKebab(pkg),
			PkgName: pkg,
			Actions: actions,
		})
	}

	t, err := template.New("tmpl").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, services); err != nil {
		log.Fatal(err)
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		os.Stdout.Write(buf.Bytes())
		log.Fatal(err)
	}
	os.Stdout.Write(b)
}

func getFlags(name string, typ reflect.Type) []*Flag {
	isPtr := false
	if typ.Kind() == reflect.Ptr {
		isPtr = true
		typ = typ.Elem()
	}
	var flgs []*Flag
	switch typ.Kind() {
	case reflect.Struct:
		for i := range typ.NumField() {
			f := typ.Field(i)
			if !f.IsExported() {
				continue
			}
			switch f.Name {
			case "HTTPClient", "Context":
				continue
			case "Body":
				flgs = append(flgs, &Flag{
					Name:        "body",
					GoParamName: "Body",
					Type:        f.Type.String(),
					IsRequired:  true,
				})
			default:
				flgs = append(flgs, getFlags(f.Name, f.Type)...)
			}
		}
	case reflect.Slice:
		if typ.Elem().Comparable() {
			flgs = append(flgs, &Flag{
				Name:        name,
				GoParamName: name,
				Type:        typ.Elem().Kind().String() + "Slice",
			})
		}
	default:
		flgs = append(flgs, &Flag{
			Name:        name,
			GoParamName: name,
			Type:        typ.Kind().String(),
			IsPtr:       isPtr,
			IsRequired:  !isPtr,
		})
	}
	return flgs
}

type Service struct {
	Name    string
	PkgName string
	CmdName string
	Actions []*Action
}

type Action struct {
	CmdName         string
	ActionName      string
	GoMethodName    string
	NumOut          int
	GoParamName     string
	GoParamTypeName string
	HasGetPayload   bool
	BodyTypeName    string
	Flags           []*Flag
}

type Flag struct {
	Name        string
	GoParamName string
	Type        string
	IsPtr       bool
	IsRequired  bool
}

const tmpl = `// Code generated by internal/gen/main.go DO NOT EDIT.
package internal

import (
	"fmt"
	{{- range .}}
	"github.com/grafana/grafana-openapi-client-go/client/{{ .PkgName }}"
	{{- end }}
	"github.com/grafana/grafana-openapi-client-go/models"
	"github.com/spf13/cobra"
)

var (
	{{- range . }}
	{{ .Name }}Cmd = &cobra.Command{
		Use: "{{ .CmdName }}",
		DisableAutoGenTag: true,
	}
	{{- $serviceName := .Name }}
	{{- $pkgName := .PkgName }}
	{{- range .Actions }}
	{{- $actionName := .ActionName }}
	{{ $serviceName }}{{ .ActionName }}Cmd = &cobra.Command{
		Use: "{{ .CmdName }}",
		SilenceUsage: true,
		DisableAutoGenTag: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			api, err := gfClient()
			if err != nil {
				return err
			}
			{{- if .BodyTypeName }}
			var body {{ .BodyTypeName }}
			if err := getBodyParam(
				{{ $serviceName }}{{ $actionName }}Flag.Body,
				&body,
			); err != nil {
				return err
			}
			{{- end }}
			{{- if eq .NumOut 1 }}
			err = api.{{ .GoParamName }}.{{ .GoMethodName }}(
			{{- else if eq .NumOut 2 }}
			resp, err := api.{{ .GoParamName }}.{{ .GoMethodName }}(
			{{- else if eq .NumOut 3}}
			resp, _, err := api.{{ .GoParamName }}.{{ .GoMethodName }}(
			{{- end }}
				&{{ $pkgName }}.{{ .GoParamTypeName }}{
				{{- range .Flags }}
					{{- if eq .GoParamName "Body" }}
					Body: &body,
					{{- else }}
					{{ .GoParamName }}: {{ if .IsPtr }}&{{ end }}{{ $serviceName }}{{ $actionName }}Flag.{{ .GoParamName }},
					{{- end }}
				{{- end }}
				},
			)
			if err != nil {
				if pe, ok := err.(getPayloadError); ok {
					if err := printPayload(pe.GetPayload()); err != nil {
						return err
					}
					return err
				}
				return err
			}
			{{- if .HasGetPayload }}
			return printPayload(resp.GetPayload())
			{{- else }}
				{{- if eq .NumOut 1 }}
			fmt.Println("{}")
				{{- else }}
			fmt.Println(resp.String())
				{{- end }}
			return nil
			{{- end }}
		},
	}
	{{- end }}
	{{- range .Actions }}
	{{- if .Flags | len }}
	{{ $serviceName }}{{ .ActionName }}Flag = struct {
		{{- range .Flags }}
		{{- if eq .Name "body" }}
		{{ .Name | toCamel }} string
		{{- else }}
		{{ .GoParamName }} {{ .Type | formatGoType }}
		{{- end }}
		{{- end }}
	}{}
	{{- end }}
	{{- end }}
	{{- end }}
)

func init() {
	{{- range . }}
	{{- $serviceName := .Name }}
	rootCmd.AddCommand({{ $serviceName }}Cmd)
	{{- range .Actions }}
		{{- $actionName := .ActionName }}
		{{- range .Flags }}
		{{ $serviceName }}{{ $actionName }}Cmd.Flags().
			{{- if eq .Name "body" }}
			StringVar(
			{{- else }}
			{{ .Type | toCamel }}Var(
			{{- end }}
			&{{ $serviceName }}{{ $actionName }}Flag.{{ .GoParamName }},
			"{{ .Name | toKebab }}", 
			{{- if eq .Name "body" }}
			"", 
			"The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{\"foo\": \"bar\"}'",
			{{- else }}
			{{ .Type | zeroValue }},
			{{- if eq .Name "UID"}}
			"Unique identifier (uid)",
			{{- else if eq .Name "Name"}}
			"Name of the {{ $serviceName | toSingular }}",
			{{- else }}
			"{{ .Name }}",
			{{- end }}
			{{- end }}
		)
		{{ if .IsRequired }}
		{{ $serviceName }}{{ $actionName }}Cmd.MarkFlagRequired("{{ .Name | toKebab }}")
		{{- end }}
		{{- end }}
		{{ $serviceName }}Cmd.AddCommand({{ $serviceName }}{{ .ActionName }}Cmd)
	{{- end }}
	{{- end }}
}
`
