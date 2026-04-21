## gf datasources update-datasource-by-uid

Updates an existing data source

### Synopsis

Similar to creating a data source, `password` and `basicAuthPassword` should be defined under secureJsonData in order to be stored securely as an encrypted blob in the database. Then, the encrypted fields are listed under secureJsonFields section in the response.

If you are running Grafana Enterprise and have Fine-grained access control enabled you need to have a permission with action: `datasources:write` and scopes: `datasources:*`, `datasources:uid:*` and `datasources:uid:1` (single data source).

Body schema (UpdateDataSourceCommand):
{
  "access": string,
  "basicAuth": boolean,
  "basicAuthUser": string,
  "database": string,
  "isDefault": boolean,
  "jsonData": any,
  "name": string,
  "secureJsonData": {"key": string},
  "type": string,
  "uid": string,
  "url": string,
  "user": string,
  "version": number,
  "withCredentials": boolean
}

```
gf datasources update-datasource-by-uid [flags]
```

### Options

```
      --body string   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help          help for update-datasource-by-uid
      --uid string    UID
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

