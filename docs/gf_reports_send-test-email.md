## gf reports send-test-email

Sends test report via email

### Synopsis

Sends test report via email

Available to org admins only and with a valid license.

You need to have a permission with action `reports:send`.

Body schema (CreateOrUpdateReport):
{
  "dashboards": [
    {
      "dashboard": object,
      "reportVariables": any,
      "timeRange": object
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
    "timeRange": object
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
gf reports send-test-email [flags]
```

### Options

```
      --body string   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help          help for send-test-email
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

* [gf reports](gf_reports.md)	 - 

