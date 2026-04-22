## gf reports save-report-settings

Saves settings

### Synopsis

Saves settings

Available to org admins only and with a valid or expired license.

You need to have a permission with action `reports.settings:write`xx.

Body schema (ReportSettings):
{
  "branding": {
    "emailFooterLink": string,
    "emailFooterMode": string,
    "emailFooterText": string,
    "emailLogoUrl": string,
    "reportLogoUrl": string
  },
  "embeddedImageTheme": string,
  "id": number,
  "orgId": number,
  "pdfTheme": string,
  "userId": number
}

```
gf reports save-report-settings [flags]
```

### Options

```
      --body string                The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema   Print the JSON Schema of the request body and exit without calling the API
  -h, --help                       help for save-report-settings
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

