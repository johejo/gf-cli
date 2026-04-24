## gf reports update-report

Updates a report

### Synopsis

Updates a report

Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.admin:write` with scope `reports:id:<report ID>`.

Requesting reports using the internal id will stop workgin in the future Use the reporting apiserver to manage reports.  See: /apis/reporting.grafana.app/

Body schema (CreateOrUpdateReport):
{
  "dashboards": [
    {
      "dashboard": object,  // models.ReportDashboardID
      "reportVariables": any,
      "timeRange": object  // models.ReportTimeRange
    }
  ],
  "enableCsv": boolean,
  "enableDashboardUrl": boolean,
  "formats": [string],
  "message": string,
  "name": string,
  "options": {
    "layout": string,
    "orientation": string,
    "pdfCombineOneFile": boolean,
    "pdfShowTemplateVariables": boolean,
    "timeRange": object  // models.ReportTimeRange
  },
  "recipients": string,
  "replyTo": string,
  "scaleFactor": number,
  "schedule": {
    "dayOfMonth": string,
    "endDate": string,
    "frequency": string,
    "intervalAmount": number,
    "intervalFrequency": string,
    "startDate": string,
    "timeZone": string,
    "workdaysOnly": boolean
  },
  "state": string,
  "subject": string
}

```
gf reports update-report [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block is a type reference; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for update-report
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

