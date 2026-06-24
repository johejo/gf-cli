## gf provisioning get-contactpoints-export

Exports all contact points in provisioning file format

### Synopsis

Exports all contact points in provisioning file format

Response schema (GetContactpointsExportOK.Payload):
  apiVersion                                                        number
  contactPoints                                                     array<object>
  contactPoints[].name                                              string
  contactPoints[].orgId                                             number
  contactPoints[].receivers                                         array<object>
  contactPoints[].receivers[].disableResolveMessage                 boolean
  contactPoints[].receivers[].settings                              object
  contactPoints[].receivers[].type                                  string
  contactPoints[].receivers[].uid                                   string
  groups                                                            array<object>
  groups[].folder                                                   string
  groups[].interval                                                 string
  groups[].name                                                     string
  groups[].orgId                                                    number
  groups[].rules                                                    array<object>
  groups[].rules[].annotations                                      map<string, string>
  groups[].rules[].condition                                        string
  groups[].rules[].dashboardUid                                     string
  groups[].rules[].data                                             array<object>
  groups[].rules[].data[].datasourceUid                             string
  groups[].rules[].data[].model                                     map<string, object>
  groups[].rules[].data[].queryType                                 string
  groups[].rules[].data[].refId                                     string
  groups[].rules[].data[].relativeTimeRange                         object
  groups[].rules[].data[].relativeTimeRange.from                    number
  groups[].rules[].data[].relativeTimeRange.to                      number
  groups[].rules[].execErrState                                     string                enum: OK | Alerting | Error
  groups[].rules[].for                                              string
  groups[].rules[].isPaused                                         boolean
  groups[].rules[].keepFiringFor                                    string
  groups[].rules[].labels                                           map<string, string>
  groups[].rules[].missing_series_evals_to_resolve                  number
  groups[].rules[].noDataState                                      string                enum: Alerting | NoData | OK
  groups[].rules[].notification_settings                            object
  groups[].rules[].notification_settings.active_time_intervals      array<string>
  groups[].rules[].notification_settings.group_by                   array<string>
  groups[].rules[].notification_settings.group_interval             string
  groups[].rules[].notification_settings.group_wait                 string
  groups[].rules[].notification_settings.mute_time_intervals        array<string>
  groups[].rules[].notification_settings.receiver                   string
  groups[].rules[].notification_settings.repeat_interval            string
  groups[].rules[].panelId                                          number
  groups[].rules[].record                                           object
  groups[].rules[].record.from                                      string
  groups[].rules[].record.metric                                    string
  groups[].rules[].record.targetDatasourceUid                       string
  groups[].rules[].title                                            string
  groups[].rules[].uid                                              string
  muteTimes                                                         array<object>
  muteTimes[].name                                                  string
  muteTimes[].orgId                                                 number
  muteTimes[].time_intervals                                        array<object>
  muteTimes[].time_intervals[].name                                 string
  muteTimes[].time_intervals[].time_intervals                       array<object>
  muteTimes[].time_intervals[].time_intervals[].days_of_month       array<string>
  muteTimes[].time_intervals[].time_intervals[].location            string
  muteTimes[].time_intervals[].time_intervals[].months              array<string>
  muteTimes[].time_intervals[].time_intervals[].times               array<object>
  muteTimes[].time_intervals[].time_intervals[].times[].end_time    string
  muteTimes[].time_intervals[].time_intervals[].times[].start_time  string
  muteTimes[].time_intervals[].time_intervals[].weekdays            array<string>
  muteTimes[].time_intervals[].time_intervals[].years               array<string>
  policies                                                          array<object>
  policies[].active_time_intervals                                  array<string>
  policies[].continue                                               boolean
  policies[].group_by                                               array<string>
  policies[].group_interval                                         string
  policies[].group_wait                                             string
  policies[].match                                                  map<string, string>   Deprecated. Remove before v1.0 release.
  policies[].match_re                                               map<string, string>
  policies[].matchers                                               array<object>
  policies[].matchers[].Name                                        string
  policies[].matchers[].Type                                        number
  policies[].matchers[].Value                                       string
  policies[].mute_time_intervals                                    array<string>
  policies[].object_matchers                                        array<array<string>>
  policies[].orgId                                                  number
  policies[].receiver                                               string
  policies[].repeat_interval                                        string
  policies[].routes                                                 array<object>
  policies[].routes[].active_time_intervals                         array<string>
  policies[].routes[].continue                                      boolean
  policies[].routes[].group_by                                      array<string>
  policies[].routes[].group_interval                                string
  policies[].routes[].group_wait                                    string
  policies[].routes[].match                                         map<string, string>   Deprecated. Remove before v1.0 release.
  policies[].routes[].match_re                                      map<string, string>
  policies[].routes[].matchers                                      array<object>
  policies[].routes[].matchers[].Name                               string
  policies[].routes[].matchers[].Type                               number
  policies[].routes[].matchers[].Value                              string
  policies[].routes[].mute_time_intervals                           array<string>
  policies[].routes[].object_matchers                               array<array<string>>
  policies[].routes[].receiver                                      string
  policies[].routes[].repeat_interval                               string
  policies[].routes[].routes                                        array<object>

```
gf provisioning get-contactpoints-export [flags]
```

### Options

```
      --decrypt                        Whether any contained secure settings should be decrypted or left redacted. Redacted settings will contain RedactedValue instead. Currently, only org admin can view decrypted secure settings.
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --download                       Whether to initiate a download of the file or not.
      --format string                  Format of the downloaded file. Supported yaml, json or hcl. Accept header can also be used, but the query parameter will take precedence. (default "yaml")
  -h, --help                           help for get-contactpoints-export
      --name string                    Filter by name
      --raw                            Print the raw HTTP response body instead of the decoded payload
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

