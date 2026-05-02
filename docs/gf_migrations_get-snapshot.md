## gf migrations get-snapshot

Gets metadata about a snapshot including where it is in its processing and final results

### Synopsis

Gets metadata about a snapshot including where it is in its processing and final results

Response schema (GetSnapshotOK.Payload):
  created               string
  finished              string
  results               array<object>
  results[].errorCode   string               enum: ALERT_RULES_QUOTA_REACHED | ALERT_RULES_GROUP_QUOTA_REACHED | DATASOURCE_NAME_CONFLICT | DATASOURCE_INVALID_URL | DATASOURCE_ALREADY_MANAGED | FOLDER_NAME_CONFLICT | DASHBOARD_ALREADY_MANAGED | LIBRARY_ELEMENT_NAME_CONFLICT | UNSUPPORTED_DATA_TYPE | RESOURCE_CONFLICT | UNEXPECTED_STATUS_CODE | INTERNAL_SERVICE_ERROR | GENERIC_ERROR
  results[].message     string
  results[].name        string
  results[].parentName  string
  results[].refId       string               REQUIRED
  results[].status      string               REQUIRED
                                             enum: OK | WARNING | ERROR | PENDING | UNKNOWN
  results[].type        string               REQUIRED
                                             enum: DASHBOARD | DATASOURCE | FOLDER | LIBRARY_ELEMENT | ALERT_RULE | ALERT_RULE_GROUP | CONTACT_POINT | NOTIFICATION_POLICY | NOTIFICATION_TEMPLATE | MUTE_TIMING | PLUGIN
  sessionUid            string
  stats                 object
  stats.statuses        map<string, number>
  stats.total           number
  stats.types           map<string, number>
  status                string               enum: INITIALIZING | CREATING | PENDING_UPLOAD | UPLOADING | PENDING_PROCESSING | PROCESSING | FINISHED | CANCELED | ERROR | UNKNOWN
  uid                   string

```
gf migrations get-snapshot [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --errors-only                    ErrorsOnly is used to only return resources with error statuses
  -h, --help                           help for get-snapshot
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --result-limit int               Max limit for snapshot results returned. Default: 100
      --result-page int                ResultPage is used for pagination with ResultLimit Default: 1
      --result-sort-column string      ResultSortColumn can be used to override the default system sort. Valid values are "name", "resource_type", and "status". Default: "default"
      --result-sort-order string       ResultSortOrder is used with ResultSortColumn. Valid values are ASC and DESC. Default: "ASC"
      --snapshot-uid string            UID of a snapshot [required]
      --uid string                     Session UID of a session [required]
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

* [gf migrations](gf_migrations.md)	 - Migrations API

