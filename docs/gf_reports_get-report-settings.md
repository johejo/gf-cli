## gf reports get-report-settings

Gets report settings

### Synopsis

Gets report settings

Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.settings:read`x.

Response schema (GetReportSettingsOK.Payload):
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

```
gf reports get-report-settings [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-report-settings
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

