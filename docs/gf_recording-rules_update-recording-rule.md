## gf recording-rules update-recording-rule

Updates the active status of a rule

### Synopsis

Updates the active status of a rule

Body schema (RecordingRuleJSON):
{
  "active": boolean,
  "count": boolean,
  "description": string,
  "dest_data_source_uid": string,
  "id": string,
  "interval": number,
  "name": string,
  "prom_name": string,
  "queries": [any],
  "range": number,
  "target_ref_id": string
}

```
gf recording-rules update-recording-rule [flags]
```

### Options

```
      --body string                The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema   Print the JSON Schema of the request body and exit without calling the API
  -h, --help                       help for update-recording-rule
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

* [gf recording-rules](gf_recording-rules.md)	 - Recording rules API

