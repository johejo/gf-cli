## gf provisioning put-alert-rule



### Synopsis

Body schema (ProvisionedAlertRule):
{
  "annotations": {"key": string},
  "condition": string,
  "data": [
    {
      "datasourceUid": string,
      "model": any,
      "queryType": string,
      "refId": string,
      "relativeTimeRange": object
    }
  ],
  "execErrState": string,
  "folderUID": string,
  "for": string,
  "id": number,
  "isPaused": boolean,
  "keep_firing_for": string,
  "labels": {"key": string},
  "missingSeriesEvalsToResolve": number,
  "noDataState": string,
  "notification_settings": {
    "active_time_intervals": [string],
    "group_by": [string],
    "group_interval": string,
    "group_wait": string,
    "mute_time_intervals": [string],
    "receiver": string,
    "repeat_interval": string
  },
  "orgID": number,
  "provenance": string,
  "record": {
    "from": string,
    "metric": string,
    "target_datasource_uid": string
  },
  "ruleGroup": string,
  "title": string,
  "uid": string,
  "updated": string
}
  condition                REQUIRED
  data                     REQUIRED
  data[].datasourceUid     Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
  data[].model             JSON is the raw JSON query and includes the above properties as well as custom properties.
  data[].queryType         QueryType is an optional identifier for the type of query.
                           It can be used to distinguish different types of queries.
  data[].refId             RefID is the unique identifier of the query, set by the frontend call.
  execErrState             REQUIRED
                           enum: OK | Alerting | Error
  folderUID                REQUIRED
  for                      REQUIRED
  noDataState              REQUIRED
                           enum: Alerting | NoData | OK
  notification_settings.active_time_intervals Override the times when notifications should not be muted. These must match the name of a mute time interval defined
                           in the alertmanager configuration time_intervals section. All notifications will be suppressed unless they are sent
                           at the time that matches any interval.
  notification_settings.group_by Override the labels by which incoming alerts are grouped together. For example, multiple alerts coming in for
                           cluster=A and alertname=LatencyHigh would be batched into a single group. To aggregate by all possible labels
                           use the special value '...' as the sole label name.
                           This effectively disables aggregation entirely, passing through all alerts as-is. This is unlikely to be what
                           you want, unless you have a very low alert volume or your upstream notification system performs its own grouping.
                           Must include 'alertname' and 'grafana_folder' if not using '...'.
  notification_settings.group_interval Override how long to wait before sending a notification about new alerts that are added to a group of alerts for
                           which an initial notification has already been sent. (Usually ~5m or more.)
  notification_settings.group_wait Override how long to initially wait to send a notification for a group of alerts. Allows to wait for an
                           inhibiting alert to arrive or collect more initial alerts for the same group. (Usually ~0s to few minutes.)
  notification_settings.mute_time_intervals Override the times when notifications should be muted. These must match the name of a mute time interval defined
                           in the alertmanager configuration time_intervals section. When muted it will not send any notifications, but
                           otherwise acts normally.
  notification_settings.receiver REQUIRED
                           Name of the receiver to send notifications to.
  notification_settings.repeat_interval Override how long to wait before sending a notification again if it has already been sent successfully for an
                           alert. (Usually ~3h or more).
                           Note that this parameter is implicitly bound by Alertmanager's `--data.retention` configuration flag.
                           Notifications will be resent after either repeat_interval or the data retention period have passed, whichever
                           occurs first. `repeat_interval` should not be less than `group_interval`.
  orgID                    REQUIRED
  record.from              REQUIRED
                           Which expression node should be used as the input for the recorded metric.
  record.metric            REQUIRED
                           Name of the recorded metric.
  record.target_datasource_uid Which data source should be used to write the output of the recording rule, specified by UID.
  ruleGroup                REQUIRED
                           rule group
                           Max Length: 190
                           Min Length: 1
  title                    REQUIRED
                           title
                           Max Length: 190
                           Min Length: 1
  uid                      uid
                           Max Length: 40
                           Min Length: 1
  updated                  updated
                           Read Only: true

```
gf provisioning put-alert-rule [flags]
```

### Options

```
      --body string                   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema      Print the JSON Schema of the request body and exit without calling the API
  -h, --help                          help for put-alert-rule
      --raw                           Print the raw HTTP response body instead of the decoded payload
      --uid string                    Alert rule UID
      --x-disable-provenance string   XDisableProvenance
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

