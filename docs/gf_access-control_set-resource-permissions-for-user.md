## gf access-control set-resource-permissions-for-user

Sets resource permissions for a user

### Synopsis

Sets resource permissions for a user

Assigns permissions for a resource by a given type (`:resource`) and `:resourceID` to a user or a service account. Allowed resources are `datasources`, `teams`, `dashboards`, `folders`, and `serviceaccounts`. Refer to the `/access-control/{resource}/description` endpoint for allowed Permissions.

Body schema (SetPermissionCommand):
  permission  string

Response schema (SetResourcePermissionsForUserOK.Payload):
  message  string

```
gf access-control set-resource-permissions-for-user [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for set-resource-permissions-for-user
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --resource string                Resource [required]
      --resource-id string             ResourceID [required]
      --user-id int                    UserID [required]
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

