## gf provisioning post-alert-rule



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
  condition                required
  data                     required
  execErrState             required, enum: OK | Alerting | Error
  folderUID                required
  for                      required
  noDataState              required, enum: Alerting | NoData | OK
  notification_settings.receiver required
  orgID                    required
  record.from              required
  record.metric            required
  ruleGroup                required
  title                    required

```
gf provisioning post-alert-rule [flags]
```

### Options

```
      --body string                   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema      Print the JSON Schema of the request body and exit without calling the API
  -h, --help                          help for post-alert-rule
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

