## gf folders update-folder-permissions

Updates permissions for a folder this operation will remove existing permissions if they re not included in the request

### Synopsis

Updates permissions for a folder this operation will remove existing permissions if they re not included in the request

Body schema (UpdateDashboardACLCommand):
  items               array<object>
  items[].permission  number
  items[].role        string         enum: None | Viewer | Editor | Admin
  items[].teamId      number
  items[].userId      number

Response schema (UpdateFolderPermissionsOK.Payload):
  message  string

```
gf folders update-folder-permissions [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --folder-uid string              FolderUID [required]
  -h, --help                           help for update-folder-permissions
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

* [gf folders](gf_folders.md)	 - Folders API

