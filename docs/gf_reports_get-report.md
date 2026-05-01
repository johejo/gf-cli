## gf reports get-report

Gets a report

### Synopsis

Gets a report

Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports:read` with scope `reports:id:<report ID>`.

Requesting reports using the internal id will stop workgin in the future Use the reporting apiserver to manage reports.  See: /apis/reporting.grafana.app/

Response schema (GetReportOK.Payload):
  created                           string
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
  id                                number
  message                           string
  name                              string
  options                           object
  options.layout                    string
  options.orientation               string
  options.pdfCombineOneFile         boolean
  options.pdfShowTemplateVariables  boolean
  options.timeRange                 object
  options.timeRange.from            string
  options.timeRange.to              string
  orgId                             number
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
  uid                               string
  updated                           string
  userId                            number

```
gf reports get-report [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-report
      --id int                         ID
      --raw                            Print the raw HTTP response body instead of the decoded payload
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

* [gf reports](gf_reports.md)	 - Reports API

