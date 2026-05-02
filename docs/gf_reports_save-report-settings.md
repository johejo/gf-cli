## gf reports save-report-settings

Saves settings

### Synopsis

Saves settings

Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.settings:write`xx.

Body schema (ReportSettings):
  branding                  object
  branding.emailFooterLink  string
  branding.emailFooterMode  string
  branding.emailFooterText  string
  branding.emailLogoUrl     string
  branding.reportLogoUrl    string
  embeddedImageTheme        string
  footerFontFamily          string
  footerItems               array<object>
  footerItems[].color       string
  footerItems[].fontSize    string
  footerItems[].fontStyle   string
  footerItems[].fontWeight  string
  footerItems[].type        string
  footerItems[].value       string
  id                        number
  orgId                     number
  pdfDashboardTitleEnabled  boolean
  pdfHeaderEnabled          boolean
  pdfTheme                  string
  pdfTimeRangeEnabled       boolean
  userId                    number

Response schema (SaveReportSettingsOK.Payload):
  message  string

```
gf reports save-report-settings [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for save-report-settings
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

