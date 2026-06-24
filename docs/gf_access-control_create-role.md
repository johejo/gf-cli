## gf access-control create-role

Creates a new custom role

### Synopsis

Creates a new custom role and maps given permissions to that role. Note that roles with the same prefix as Fixed Roles can’t be created.

You need to have a permission with action `roles:write` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only create custom roles with the same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to create a custom role which allows to do that. This is done to prevent escalation of privileges.

Body schema (CreateRoleForm):
  description            string
  displayName            string
  global                 boolean
  group                  string
  hidden                 boolean
  name                   string
  permissions            array<object>
  permissions[].action   string
  permissions[].created  string
  permissions[].scope    string
  permissions[].updated  string
  uid                    string

Response schema (CreateRoleCreated.Payload):
  created                string         REQUIRED
  delegatable            boolean
  description            string         REQUIRED
  displayName            string         REQUIRED
  global                 boolean
  group                  string         REQUIRED
  hidden                 boolean
  mapped                 boolean
  name                   string         REQUIRED
  permissions            array<object>
  permissions[].action   string
  permissions[].created  string
  permissions[].scope    string
  permissions[].updated  string
  uid                    string         REQUIRED
  updated                string         REQUIRED
  version                number         REQUIRED

```
gf access-control create-role [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for create-role
      --raw                            Print the raw HTTP response body instead of the decoded payload
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf access-control](gf_access-control.md)	 - Access control API

