## gf convert-prometheus convert-prometheus-post-rule-group



### Synopsis

Body schema (PrometheusRuleGroup):
{
  "interval": number,
  "labels": {"key": string},
  "limit": number,
  "name": string,
  "query_offset": string,
  "rules": [
    {
      "alert": string,
      "annotations": {"key": string},
      "expr": string,
      "for": string,
      "keep_firing_for": string,
      "labels": {"key": string},
      "record": string
    }
  ]
}

```
gf convert-prometheus convert-prometheus-post-rule-group [flags]
```

### Options

```
      --body string                                       Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block is a type reference; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema                          Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema                      Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                                              help for convert-prometheus-post-rule-group
      --namespace-title string                            NamespaceTitle
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

