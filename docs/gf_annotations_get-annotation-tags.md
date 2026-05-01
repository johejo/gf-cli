## gf annotations get-annotation-tags

Finds annotations tags

### Synopsis

Finds annotations tags

Find all the event tags created in the annotations.

```
gf annotations get-annotation-tags [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-annotation-tags
      --limit string                   Max limit for results returned. Default: "100"
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --tag string                     Tag is a string that you can use to filter tags.
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

* [gf annotations](gf_annotations.md)	 - Annotations API

