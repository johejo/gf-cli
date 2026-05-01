## gf licensing post-license-token

Creates license token

### Synopsis

Creates license token

You need to have a permission with action `licensing:write`.

Body schema (DeleteTokenCommand):
  instance  string

Response schema (PostLicenseTokenOK.Payload):
  account                       string
  anonymousRatio                number
  company                       string
  details_url                   string
  exp                           number
  iat                           number
  included_users                number
  iss                           string
  jti                           string
  lexp                          number
  lic_exp_warn_days             number
  lid                           string
  limit_by                      string
  max_concurrent_user_sessions  number
  nbf                           number
  prod                          array<string>
  slug                          string
  status                        number
  sub                           string
  tok_exp_warn_days             number
  trial                         boolean
  trial_exp                     number
  update_days                   number
  usage_billing                 boolean

```
gf licensing post-license-token [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for post-license-token
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

* [gf licensing](gf_licensing.md)	 - Licensing API

