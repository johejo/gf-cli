## gf provisioning put-alert-rule-group

Creates or update alert rule group

### Synopsis

Creates or update alert rule group

Body schema (AlertRuleGroup):
  folderUid                                            string
  interval                                             number
  rules                                                array<object>
  rules[].annotations                                  map<string, string>
  rules[].condition                                    string               REQUIRED
  rules[].data                                         array<object>        REQUIRED
  rules[].data[].datasourceUid                         string               Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
  rules[].data[].model                                 object               JSON is the raw JSON query and includes the above properties as well as custom properties.
  rules[].data[].queryType                             string               QueryType is an optional identifier for the type of query.
                                                                            It can be used to distinguish different types of queries.
  rules[].data[].refId                                 string               RefID is the unique identifier of the query, set by the frontend call.
  rules[].data[].relativeTimeRange                     object
  rules[].data[].relativeTimeRange.from                number
  rules[].data[].relativeTimeRange.to                  number
  rules[].execErrState                                 string               REQUIRED
                                                                            enum: OK | Alerting | Error
  rules[].folderUID                                    string               REQUIRED
  rules[].for                                          string               REQUIRED
  rules[].id                                           number
  rules[].isPaused                                     boolean
  rules[].keep_firing_for                              string
  rules[].labels                                       map<string, string>
  rules[].missingSeriesEvalsToResolve                  number
  rules[].noDataState                                  string               REQUIRED
                                                                            enum: Alerting | NoData | OK
  rules[].notification_settings                        object
  rules[].notification_settings.active_time_intervals  array<string>        Override the times when notifications should not be muted. These must match the name of a mute time interval defined
                                                                            in the alertmanager configuration time_intervals section. All notifications will be suppressed unless they are sent
                                                                            at the time that matches any interval.
  rules[].notification_settings.group_by               array<string>        Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
                                                                            cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
                                                                            use the special value '...' as the sole label name.
                                                                            This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
                                                                            you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
                                                                            Must include 'alertname' and 'grafana_folder' if not using '...'.
  rules[].notification_settings.group_interval         string               Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
                                                                            which an initial notification has already been sent. (Usually ~5m or more.)
  rules[].notification_settings.group_wait             string               Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
                                                                            inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
  rules[].notification_settings.mute_time_intervals    array<string>        Override the times when notifications should be muted. These must match the name of a mute time interval defined
                                                                            in the alertmanager configuration time_intervals section. When muted it will not send any notifications, but
                                                                            otherwise acts normally.
  rules[].notification_settings.receiver               string               REQUIRED
                                                                            Name of the receiver to send notifications to.
  rules[].notification_settings.repeat_interval        string               Override how long to wait before sending a notification again if it has already been sent successfully for an
                                                                            alert. (Usually ~3h or more).
                                                                            Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
                                                                            Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
                                                                            occurs first. `repeat_interval` should not be less than `group_interval`.
  rules[].orgID                                        number               REQUIRED
  rules[].provenance                                   string
  rules[].record                                       object
  rules[].record.from                                  string               REQUIRED
                                                                            Which expression node should be used as the input for the recorded metric.
  rules[].record.metric                                string               REQUIRED
                                                                            Name of the recorded metric.
  rules[].record.target_datasource_uid                 string               Which data source should be used to write the output of the recording rule, specified by UID.
  rules[].ruleGroup                                    string               REQUIRED
                                                                            rule group
                                                                            Max Length: 190
                                                                            Min Length: 1
  rules[].title                                        string               REQUIRED
                                                                            title
                                                                            Max Length: 190
                                                                            Min Length: 1
  rules[].uid                                          string               uid
                                                                            Max Length: 40
                                                                            Min Length: 1
  rules[].updated                                      string               updated
                                                                            Read Only: true
  title                                                string

Response schema (PutAlertRuleGroupOK.Payload):
  folderUid                                            string
  interval                                             number
  rules                                                array<object>
  rules[].annotations                                  map<string, string>
  rules[].condition                                    string               REQUIRED
  rules[].data                                         array<object>        REQUIRED
  rules[].data[].datasourceUid                         string               Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
  rules[].data[].model                                 object               JSON is the raw JSON query and includes the above properties as well as custom properties.
  rules[].data[].queryType                             string               QueryType is an optional identifier for the type of query.
                                                                            It can be used to distinguish different types of queries.
  rules[].data[].refId                                 string               RefID is the unique identifier of the query, set by the frontend call.
  rules[].data[].relativeTimeRange                     object
  rules[].data[].relativeTimeRange.from                number
  rules[].data[].relativeTimeRange.to                  number
  rules[].execErrState                                 string               REQUIRED
                                                                            enum: OK | Alerting | Error
  rules[].folderUID                                    string               REQUIRED
  rules[].for                                          string               REQUIRED
  rules[].id                                           number
  rules[].isPaused                                     boolean
  rules[].keep_firing_for                              string
  rules[].labels                                       map<string, string>
  rules[].missingSeriesEvalsToResolve                  number
  rules[].noDataState                                  string               REQUIRED
                                                                            enum: Alerting | NoData | OK
  rules[].notification_settings                        object
  rules[].notification_settings.active_time_intervals  array<string>        Override the times when notifications should not be muted. These must match the name of a mute time interval defined
                                                                            in the alertmanager configuration time_intervals section. All notifications will be suppressed unless they are sent
                                                                            at the time that matches any interval.
  rules[].notification_settings.group_by               array<string>        Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
                                                                            cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
                                                                            use the special value '...' as the sole label name.
                                                                            This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
                                                                            you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
                                                                            Must include 'alertname' and 'grafana_folder' if not using '...'.
  rules[].notification_settings.group_interval         string               Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
                                                                            which an initial notification has already been sent. (Usually ~5m or more.)
  rules[].notification_settings.group_wait             string               Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
                                                                            inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
  rules[].notification_settings.mute_time_intervals    array<string>        Override the times when notifications should be muted. These must match the name of a mute time interval defined
                                                                            in the alertmanager configuration time_intervals section. When muted it will not send any notifications, but
                                                                            otherwise acts normally.
  rules[].notification_settings.receiver               string               REQUIRED
                                                                            Name of the receiver to send notifications to.
  rules[].notification_settings.repeat_interval        string               Override how long to wait before sending a notification again if it has already been sent successfully for an
                                                                            alert. (Usually ~3h or more).
                                                                            Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
                                                                            Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
                                                                            occurs first. `repeat_interval` should not be less than `group_interval`.
  rules[].orgID                                        number               REQUIRED
  rules[].provenance                                   string
  rules[].record                                       object
  rules[].record.from                                  string               REQUIRED
                                                                            Which expression node should be used as the input for the recorded metric.
  rules[].record.metric                                string               REQUIRED
                                                                            Name of the recorded metric.
  rules[].record.target_datasource_uid                 string               Which data source should be used to write the output of the recording rule, specified by UID.
  rules[].ruleGroup                                    string               REQUIRED
                                                                            rule group
                                                                            Max Length: 190
                                                                            Min Length: 1
  rules[].title                                        string               REQUIRED
                                                                            title
                                                                            Max Length: 190
                                                                            Min Length: 1
  rules[].uid                                          string               uid
                                                                            Max Length: 40
                                                                            Min Length: 1
  rules[].updated                                      string               updated
                                                                            Read Only: true
  title                                                string

```
gf provisioning put-alert-rule-group [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block above lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
      --folder-uid string              FolderUID
      --group string                   Group
  -h, --help                           help for put-alert-rule-group
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --x-disable-provenance string    XDisableProvenance
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

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

