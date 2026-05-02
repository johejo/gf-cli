## gf annotations get-annotations

Finds annotations

### Synopsis

Finds annotations

Starting in Grafana v6.4 regions annotations are now returned in one entity that now includes the timeEnd property.

```
gf annotations get-annotations [flags]
```

### Options

```
      --alert-id int                   Find annotations for a specified alert rule by its ID. deprecated: AlertID is deprecated and will be removed in future versions. Please use AlertUID instead.
      --alert-uid string               Find annotations for a specified alert rule by its UID.
      --dashboard-id int               Find annotations that are scoped to a specific dashboard
      --dashboard-uid string           Find annotations that are scoped to a specific dashboard
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --from int                       Find annotations created after specific epoch datetime in milliseconds.
  -h, --help                           help for get-annotations
      --limit int                      Max limit for results returned.
      --match-any                      Match any or all tags
      --panel-id int                   Find annotations that are scoped to a specific panel
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --tags strings                   Use this to filter organization annotations. Organization annotations are annotations from an annotation data source that are not connected specifically to a dashboard or panel. You can filter by multiple tags. [required]
      --to int                         Find annotations created before specific epoch datetime in milliseconds.
      --type string                    Return alerts or user created annotations
      --user-id int                    Limit response to annotations created by specific user.
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

