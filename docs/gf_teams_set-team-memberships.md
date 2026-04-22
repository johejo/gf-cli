## gf teams set-team-memberships

Sets team memberships

### Synopsis

Sets team memberships

Takes user emails, and updates team members and admins to the provided lists of users. Any current team members and admins not in the provided lists will be removed.

Body schema (SetTeamMembershipsCommand):
{
  "admins": [string],
  "members": [string]
}

```
gf teams set-team-memberships [flags]
```

### Options

```
      --body string                The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema   Print the JSON Schema of the request body and exit without calling the API
  -h, --help                       help for set-team-memberships
      --team-id string             TeamID
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

