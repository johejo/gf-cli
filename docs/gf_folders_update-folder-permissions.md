## gf folders update-folder-permissions

Updates permissions for a folder this operation will remove existing permissions if they re not included in the request

### Synopsis

Updates permissions for a folder this operation will remove existing permissions if they re not included in the request

Body schema (UpdateDashboardACLCommand):
{
  "items": [
    {
      "permission": number,
      "role": string,
      "teamId": number,
      "userId": number
    }
  ]
}
  items[].role             enum: None | Viewer | Editor | Admin

```
gf folders update-folder-permissions [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block is a type reference; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --folder-uid string              FolderUID
  -h, --help                           help for update-folder-permissions
      --raw                            Print the raw HTTP response body instead of the decoded payload
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

* [gf folders](gf_folders.md)	 - Folders API

