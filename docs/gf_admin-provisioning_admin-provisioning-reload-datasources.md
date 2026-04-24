## gf admin-provisioning admin-provisioning-reload-datasources

Reloads datasource provisioning configurations

### Synopsis

Reloads datasource provisioning configurations

Reloads the provisioning config files for datasources again. It won’t return until the new provisioned entities are already stored in the database. In case of dashboards, it will stop polling for changes in dashboard files and then restart it with new configurations after returning. If you are running Grafana Enterprise and have Fine-grained access control enabled, you need to have a permission with action `provisioning:reload` and scope `provisioners:datasources`.

```
gf admin-provisioning admin-provisioning-reload-datasources [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for admin-provisioning-reload-datasources
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

* [gf admin-provisioning](gf_admin-provisioning.md)	 - Admin provisioning API

