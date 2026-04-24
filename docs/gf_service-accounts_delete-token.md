## gf service-accounts delete-token

Deletes token deletes service account tokens

### Synopsis

Deletes token deletes service account tokens

Required permissions (See note in the [introduction](https://grafana.com/docs/grafana/latest/developers/http_api/serviceaccount/#service-account-api) for an explanation): action: `serviceaccounts:write` scope: `serviceaccounts:id:1` (single service account)

Requires basic authentication and that the authenticated user is a Grafana Admin.

```
gf service-accounts delete-token [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for delete-token
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --service-account-id int         ServiceAccountID
      --token-id int                   TokenID
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

* [gf service-accounts](gf_service-accounts.md)	 - Service accounts API

