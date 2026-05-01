## gf teams update-team-preferences

Updates team preferences

### Synopsis

Updates team preferences

Body schema (UpdatePrefsCmd):
  homeDashboardId       number         The numerical :id of a favorited dashboard
  homeDashboardUID      string
  language              string
  navbar                object
  navbar.bookmarkUrls   array<string>
  queryHistory          object
  queryHistory.homeTab  string
  regionalFormat        string
  theme                 string         enum: light | dark | system
  timezone              string         Any IANA timezone string (e.g. America/New_York), 'utc', 'browser', or empty string
  weekStart             string

Response schema (UpdateTeamPreferencesOK.Payload):
  message  string

```
gf teams update-team-preferences [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-team-preferences
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --team-id string                 TeamID
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

* [gf teams](gf_teams.md)	 - Teams API

