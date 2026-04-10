## gf dashboards



```
gf dashboards [flags]
```

### Options

```
  -h, --help   help for dashboards
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

* [gf](gf.md)	 - CLI for Grafana API
* [gf dashboards create-dashboard-snapshot](gf_dashboards_create-dashboard-snapshot.md)	 - Whens creating a snapshot using the API you have to provide the full dashboard payload including the snapshot data this endpoint is designed for the grafana UI
* [gf dashboards create-public-dashboard](gf_dashboards_create-public-dashboard.md)	 - Create public dashboard for a dashboard
* [gf dashboards delete-dashboard-by-uid](gf_dashboards_delete-dashboard-by-uid.md)	 - Deletes dashboard by uid
* [gf dashboards delete-dashboard-snapshot](gf_dashboards_delete-dashboard-snapshot.md)	 - Deletes snapshot by key
* [gf dashboards delete-dashboard-snapshot-by-delete-key](gf_dashboards_delete-dashboard-snapshot-by-delete-key.md)	 - Deletes snapshot by delete key
* [gf dashboards delete-public-dashboard](gf_dashboards_delete-public-dashboard.md)	 - Delete public dashboard for a dashboard
* [gf dashboards get-dashboard-by-uid](gf_dashboards_get-dashboard-by-uid.md)	 - Gets dashboard by uid
* [gf dashboards get-dashboard-permissions-list-by-uid](gf_dashboards_get-dashboard-permissions-list-by-uid.md)	 - Gets all existing permissions for the given dashboard
* [gf dashboards get-dashboard-snapshot](gf_dashboards_get-dashboard-snapshot.md)	 - Gets snapshot by key
* [gf dashboards get-dashboard-tags](gf_dashboards_get-dashboard-tags.md)	 - Gets all dashboards tags of an organisation
* [gf dashboards get-dashboard-version-by-uid](gf_dashboards_get-dashboard-version-by-uid.md)	 - Gets a specific dashboard version using UID
* [gf dashboards get-dashboard-versions-by-uid](gf_dashboards_get-dashboard-versions-by-uid.md)	 - 
* [gf dashboards get-home-dashboard](gf_dashboards_get-home-dashboard.md)	 - Gets home dashboard
* [gf dashboards get-public-annotations](gf_dashboards_get-public-annotations.md)	 - Get annotations for a public dashboard
* [gf dashboards get-public-dashboard](gf_dashboards_get-public-dashboard.md)	 - Get public dashboard by dashboardUid
* [gf dashboards import-dashboard](gf_dashboards_import-dashboard.md)	 - Imports dashboard
* [gf dashboards interpolate-dashboard](gf_dashboards_interpolate-dashboard.md)	 - Interpolates dashboard this is an experimental endpoint under dashboard library or suggested dashboards feature flags and is subject to change
* [gf dashboards list-public-dashboards](gf_dashboards_list-public-dashboards.md)	 - Get list of public dashboards
* [gf dashboards post-dashboard](gf_dashboards_post-dashboard.md)	 - Creates update dashboard
* [gf dashboards query-public-dashboard](gf_dashboards_query-public-dashboard.md)	 - Get results for a given panel on a public dashboard
* [gf dashboards restore-dashboard-version-by-uid](gf_dashboards_restore-dashboard-version-by-uid.md)	 - Restores a dashboard to a given dashboard version using UID
* [gf dashboards search-dashboard-snapshots](gf_dashboards_search-dashboard-snapshots.md)	 - 
* [gf dashboards update-dashboard-permissions-by-uid](gf_dashboards_update-dashboard-permissions-by-uid.md)	 - Updates permissions for a dashboard
* [gf dashboards update-public-dashboard](gf_dashboards_update-public-dashboard.md)	 - 
* [gf dashboards view-public-dashboard](gf_dashboards_view-public-dashboard.md)	 - Get public dashboard for view

