## gf datasources create-correlation

Adds correlation

### Synopsis

Body schema (CreateCorrelationCommand):
{
  "config": {
    "field": string,
    "target": any,
    "transformations": [object],
    "type": string
  },
  "description": string,
  "label": string,
  "provisioned": boolean,
  "targetUID": string,
  "type": string
}
  config.field             required
  config.target            required

```
gf datasources create-correlation [flags]
```

### Options

```
      --body string         The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help                help for create-correlation
      --source-uid string   SourceUID
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

* [gf datasources](gf_datasources.md)	 - 

