## gf teams get-team-preferences

Gets team preferences

### Synopsis

Gets team preferences

Response schema (GetTeamPreferencesOK.Payload):
  homeDashboardUID      string         UID for the home dashboard
  language              string         Selected language (beta)
  navbar                object
  navbar.bookmarkUrls   array<string>
  queryHistory          object
  queryHistory.homeTab  string         one of: '' | 'query' | 'starred';
  regionalFormat        string         Selected locale (beta)
  theme                 string         light, dark, empty is default
  timezone              string         The timezone selection
                                       TODO: this should use the timezone defined in common
  weekStart             string         day of the week (sunday, monday, etc)

```
gf teams get-team-preferences [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-team-preferences
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

