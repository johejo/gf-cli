## gf datasources

Datasources API

```
gf datasources [flags]
```

### Options

```
  -h, --help   help for datasources
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

* [gf](gf.md)	 - CLI for Grafana API
* [gf datasources add-datasource](gf_datasources_add-datasource.md)	 - Creates a data source
* [gf datasources call-datasource-resource](gf_datasources_call-datasource-resource.md)	 - Fetches data source resources
* [gf datasources check-datasource-health](gf_datasources_check-datasource-health.md)	 - Sends a health check request to the plugin datasource identified by the UID
* [gf datasources create-correlation](gf_datasources_create-correlation.md)	 - Adds correlation
* [gf datasources delete-correlation](gf_datasources_delete-correlation.md)	 - Deletes a correlation
* [gf datasources delete-datasource-by-name](gf_datasources_delete-datasource-by-name.md)	 - Deletes an existing data source by name this function will be removed in the future
* [gf datasources delete-datasource-by-uid](gf_datasources_delete-datasource-by-uid.md)	 - Deletes an existing data source by UID
* [gf datasources get-correlation](gf_datasources_get-correlation.md)	 - Gets a correlation
* [gf datasources get-correlations](gf_datasources_get-correlations.md)	 - Gets all correlations
* [gf datasources get-correlations-by-source-uid](gf_datasources_get-correlations-by-source-uid.md)	 - Gets all correlations originating from the given data source
* [gf datasources get-datasource-by-name](gf_datasources_get-datasource-by-name.md)	 - Gets a single data source by name this function will be removed in the future
* [gf datasources get-datasource-by-uid](gf_datasources_get-datasource-by-uid.md)	 - Gets a single data source by UID
* [gf datasources get-datasource-id-by-name](gf_datasources_get-datasource-id-by-name.md)	 - Gets data source Id by name this function will be removed in the future
* [gf datasources get-datasources](gf_datasources_get-datasources.md)	 - Gets all data sources
* [gf datasources query-metrics-with-expressions](gf_datasources_query-metrics-with-expressions.md)	 - Data source query metrics with expressions
* [gf datasources update-correlation](gf_datasources_update-correlation.md)	 - Updates a correlation
* [gf datasources update-datasource-by-uid](gf_datasources_update-datasource-by-uid.md)	 - Updates an existing data source

