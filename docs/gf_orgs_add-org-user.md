## gf orgs add-org-user

Adds a new user to the current organization

### Synopsis

Adds a new user to the current organization

Adds a global user to the current organization.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `org.users:add` with scope `users:*`.

Body schema (AddOrgUserCommand):
  loginOrEmail  string
  role          string  enum: None | Viewer | Editor | Admin

Response schema (AddOrgUserOK.Payload):
  message  string

```
gf orgs add-org-user [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for add-org-user
      --org-id int                     OrgID [required]
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
```

### SEE ALSO

* [gf orgs](gf_orgs.md)	 - Orgs API

