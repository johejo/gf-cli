## gf signed-in-user update-user-preferences

Updates user preferences

### Synopsis

Updates user preferences

Omitting a key (`theme`, `homeDashboardUID`, `timezone`) will cause the current value to be replaced with the system default value.

Body schema (UpdatePrefsCmd):
{
  "cookies": [string],
  "homeDashboardId": number,
  "homeDashboardUID": string,
  "language": string,
  "navbar": {
    "bookmarkUrls": [string]
  },
  "queryHistory": {
    "homeTab": string
  },
  "regionalFormat": string,
  "theme": string,
  "timezone": string,
  "weekStart": string
}
  homeDashboardId          The numerical :id of a favorited dashboard
  theme                    enum: light | dark | system
  timezone                 enum: utc | browser

```
gf signed-in-user update-user-preferences [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block is a type reference; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-user-preferences
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

* [gf signed-in-user](gf_signed-in-user.md)	 - Signed in user API

