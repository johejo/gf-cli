## gf access-control get-role

Gets a role

### Synopsis

Gets a role

Get a role for the given UID.

You need to have a permission with action `roles:read` and scope `roles:*`.

Response schema (GetRoleOK.Payload):
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
gf access-control get-role [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-role
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --role-uid string                RoleUID
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

* [gf access-control](gf_access-control.md)	 - Access control API

