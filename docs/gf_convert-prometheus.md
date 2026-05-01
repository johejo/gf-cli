## gf convert-prometheus

Convert prometheus API

```
gf convert-prometheus [flags]
```

### Options

```
  -h, --help   help for convert-prometheus
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

* [gf](gf.md)	 - CLI for Grafana API
* [gf convert-prometheus convert-prometheus-cortex-delete-namespace](gf_convert-prometheus_convert-prometheus-cortex-delete-namespace.md)	 - Deletes all rule groups that were imported from prometheus compatible sources within the specified namespace
* [gf convert-prometheus convert-prometheus-cortex-delete-rule-group](gf_convert-prometheus_convert-prometheus-cortex-delete-rule-group.md)	 - Deletes a specific rule group if it was imported from a prometheus compatible source
* [gf convert-prometheus convert-prometheus-cortex-get-namespace](gf_convert-prometheus_convert-prometheus-cortex-get-namespace.md)	 - Gets grafana managed alert rules that were imported from prometheus compatible sources for a specified namespace folder
* [gf convert-prometheus convert-prometheus-cortex-get-rule-group](gf_convert-prometheus_convert-prometheus-cortex-get-rule-group.md)	 - Gets a single rule group in prometheus compatible format if it was imported from a prometheus compatible source
* [gf convert-prometheus convert-prometheus-cortex-get-rules](gf_convert-prometheus_convert-prometheus-cortex-get-rules.md)	 - Gets all grafana managed alert rules that were imported from prometheus compatible sources grouped by namespace
* [gf convert-prometheus convert-prometheus-cortex-post-rule-group](gf_convert-prometheus_convert-prometheus-cortex-post-rule-group.md)	 - Converts a prometheus rule group into a grafana rule group and creates or updates it within the specified namespace
* [gf convert-prometheus convert-prometheus-cortex-post-rule-groups](gf_convert-prometheus_convert-prometheus-cortex-post-rule-groups.md)	 - Converts the submitted rule groups into grafana managed rules
* [gf convert-prometheus convert-prometheus-delete-namespace](gf_convert-prometheus_convert-prometheus-delete-namespace.md)	 - Deletes all rule groups that were imported from prometheus compatible sources within the specified namespace
* [gf convert-prometheus convert-prometheus-delete-rule-group](gf_convert-prometheus_convert-prometheus-delete-rule-group.md)	 - Deletes a specific rule group if it was imported from a prometheus compatible source
* [gf convert-prometheus convert-prometheus-get-namespace](gf_convert-prometheus_convert-prometheus-get-namespace.md)	 - Gets grafana managed alert rules that were imported from prometheus compatible sources for a specified namespace folder
* [gf convert-prometheus convert-prometheus-get-rule-group](gf_convert-prometheus_convert-prometheus-get-rule-group.md)	 - Gets a single rule group in prometheus compatible format if it was imported from a prometheus compatible source
* [gf convert-prometheus convert-prometheus-get-rules](gf_convert-prometheus_convert-prometheus-get-rules.md)	 - Gets all grafana managed alert rules that were imported from prometheus compatible sources grouped by namespace
* [gf convert-prometheus convert-prometheus-post-rule-group](gf_convert-prometheus_convert-prometheus-post-rule-group.md)	 - Converts a prometheus rule group into a grafana rule group and creates or updates it within the specified namespace
* [gf convert-prometheus convert-prometheus-post-rule-groups](gf_convert-prometheus_convert-prometheus-post-rule-groups.md)	 - Converts the submitted rule groups into grafana managed rules

