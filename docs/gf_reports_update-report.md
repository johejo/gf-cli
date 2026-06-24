## gf reports update-report

Updates a report

### Synopsis

Updates a report

Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.admin:write` with scope `reports:id:<report ID>`.

Requesting reports using the internal id will stop workgin in the future Use the reporting apiserver to manage reports.  See: /apis/reporting.grafana.app/

Body schema (CreateOrUpdateReport):
  dashboards                        array<object>
  dashboards[].dashboard            object
  dashboards[].dashboard.id         number
  dashboards[].dashboard.name       string
  dashboards[].dashboard.uid        string
  dashboards[].reportVariables      object
  dashboards[].timeRange            object
  dashboards[].timeRange.from       string
  dashboards[].timeRange.to         string
  enableCsv                         boolean
  enableDashboardUrl                boolean
  formats                           array<string>
  message                           string
  name                              string
  options                           object
  options.csvEncoding               string
  options.layout                    string
  options.orientation               string
  options.pdfCombineOneFile         boolean
  options.pdfShowTemplateVariables  boolean
  options.timeRange                 object
  options.timeRange.from            string
  options.timeRange.to              string
  recipients                        string
  replyTo                           string
  scaleFactor                       number
  schedule                          object
  schedule.dayOfMonth               string
  schedule.endDate                  string
  schedule.frequency                string
  schedule.intervalAmount           number
  schedule.intervalFrequency        string
  schedule.startDate                string
  schedule.timeZone                 string
  schedule.workdaysOnly             boolean
  state                             string
  subject                           string

Response schema (UpdateReportOK.Payload):
  message  string

```
gf reports update-report [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-report
      --id int                         ID [required]
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

* [gf reports](gf_reports.md)	 - Reports API

