## gf access-control get-role-assignments

Gets role assignments

### Synopsis

Gets role assignments

Get role assignments for the role with the given UID. Does not include role assignments mapped through group attribute sync.

You need to have a permission with action `teams.roles:list` and scope `teams:id:*` and `users.roles:list` and scope `users:id:*`.

Response schema (GetRoleAssignmentsOK.Payload):
  role_uid          string
  service_accounts  array<number>
  teams             array<number>
  users             array<number>

```
gf access-control get-role-assignments [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-role-assignments
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --role-uid string                RoleUID [required]
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

