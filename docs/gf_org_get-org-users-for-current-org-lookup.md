## gf org get-org-users-for-current-org-lookup

Gets all users within the current organization lookup

### Synopsis

Gets all users within the current organization lookup

Returns all org users within the current organization, but with less detailed information. Accessible to users with org admin role, admin in any folder or admin of any team. Mainly used by Grafana UI for providing list of users when adding team members and when editing folder/dashboard permissions.

```
gf org get-org-users-for-current-org-lookup [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-org-users-for-current-org-lookup
      --limit int                      Limit
      --query string                   Query
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

* [gf org](gf_org.md)	 - Org API

