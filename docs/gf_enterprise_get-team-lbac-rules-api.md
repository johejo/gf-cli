## gf enterprise get-team-lbac-rules-api

Retrieves l b a c rules for a team

### Synopsis

Retrieves l b a c rules for a team

Response schema (GetTeamLBACRulesAPIOK.Payload):
  rules            array<object>
  rules[].rules    array<string>
  rules[].teamId   string
  rules[].teamUid  string

```
gf enterprise get-team-lbac-rules-api [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-team-lbac-rules-api
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --uid string                     UID [required]
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using sever behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_USERNAME)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_PASSWORD)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
```

### SEE ALSO

* [gf enterprise](gf_enterprise.md)	 - Enterprise API

