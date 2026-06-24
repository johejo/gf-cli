## gf

CLI for Grafana API

### Synopsis

Grafana API Client for command line operations

```
gf [flags]
```

### Options

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
  -h, --help                         help for gf
      --help-json                    Print the CLI schema index (commands, flags, body/response model types) as JSON and exit. It covers every command and is large; since it prints to stdout, filter it with jq/grep rather than reading it whole, e.g. gf --help-json | jq '.commands | keys'. Use --describe-body-jsonschema / --describe-response-jsonschema on a subcommand for the full JSON Schema.
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf access-control](gf_access-control.md)	 - Access control API
* [gf admin](gf_admin.md)	 - Admin API
* [gf admin-ldap](gf_admin-ldap.md)	 - Admin ldap API
* [gf admin-provisioning](gf_admin-provisioning.md)	 - Admin provisioning API
* [gf admin-users](gf_admin-users.md)	 - Admin users API
* [gf annotations](gf_annotations.md)	 - Annotations API
* [gf convert-prometheus](gf_convert-prometheus.md)	 - Convert prometheus API
* [gf dashboards](gf_dashboards.md)	 - Dashboards API
* [gf datasources](gf_datasources.md)	 - Datasources API
* [gf devices](gf_devices.md)	 - Devices API
* [gf enterprise](gf_enterprise.md)	 - Enterprise API
* [gf folders](gf_folders.md)	 - Folders API
* [gf group-attribute-sync](gf_group-attribute-sync.md)	 - Group attribute sync API
* [gf health](gf_health.md)	 - Health API
* [gf library-elements](gf_library-elements.md)	 - Library elements API
* [gf licensing](gf_licensing.md)	 - Licensing API
* [gf migrations](gf_migrations.md)	 - Migrations API
* [gf org](gf_org.md)	 - Org API
* [gf orgs](gf_orgs.md)	 - Orgs API
* [gf playlists](gf_playlists.md)	 - Playlists API
* [gf provisioning](gf_provisioning.md)	 - Provisioning API
* [gf query-history](gf_query-history.md)	 - Query history API
* [gf quota](gf_quota.md)	 - Quota API
* [gf recording-rules](gf_recording-rules.md)	 - Recording rules API
* [gf reports](gf_reports.md)	 - Reports API
* [gf saml](gf_saml.md)	 - Saml API
* [gf search](gf_search.md)	 - Search API
* [gf service-accounts](gf_service-accounts.md)	 - Service accounts API
* [gf signed-in-user](gf_signed-in-user.md)	 - Signed in user API
* [gf signing-keys](gf_signing-keys.md)	 - Signing keys API
* [gf snapshots](gf_snapshots.md)	 - Snapshots API
* [gf sso-settings](gf_sso-settings.md)	 - Sso settings API
* [gf sync-team-groups](gf_sync-team-groups.md)	 - Sync team groups API
* [gf teams](gf_teams.md)	 - Teams API
* [gf user](gf_user.md)	 - User API
* [gf users](gf_users.md)	 - Users API

