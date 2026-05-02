## gf convert-prometheus convert-prometheus-post-rule-group

Converts a prometheus rule group into a grafana rule group and creates or updates it within the specified namespace

### Synopsis

Converts a prometheus rule group into a grafana rule group and creates or updates it within the specified namespace

If the group already exists and was not imported from a Prometheus-compatible source initially, it will not be replaced and an error will be returned.

Body schema (PrometheusRuleGroup):
  interval                 number
  labels                   map<string, string>
  limit                    number
  name                     string
  query_offset             string
  rules                    array<object>
  rules[].alert            string
  rules[].annotations      map<string, string>
  rules[].expr             string
  rules[].for              string
  rules[].keep_firing_for  string
  rules[].labels           map<string, string>
  rules[].record           string

Response schema (ConvertPrometheusPostRuleGroupAccepted.Payload):
  error      string
  errorType  string
  status     string

```
gf convert-prometheus convert-prometheus-post-rule-group [flags]
```

### Options

```
      --body string                                       Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema                          Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema                      Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                                              help for convert-prometheus-post-rule-group
      --namespace-title string                            NamespaceTitle [required]
      --raw                                               Print the raw HTTP response body instead of the decoded payload
      --x-grafana-alerting-alert-rules-paused             XGrafanaAlertingAlertRulesPaused
      --x-grafana-alerting-datasource-uid string          XGrafanaAlertingDatasourceUID
      --x-grafana-alerting-folder-uid string              XGrafanaAlertingFolderUID
      --x-grafana-alerting-notification-settings string   XGrafanaAlertingNotificationSettings
      --x-grafana-alerting-recording-rules-paused         XGrafanaAlertingRecordingRulesPaused
      --x-grafana-alerting-target-datasource-uid string   XGrafanaAlertingTargetDatasourceUID
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

* [gf convert-prometheus](gf_convert-prometheus.md)	 - Convert prometheus API

