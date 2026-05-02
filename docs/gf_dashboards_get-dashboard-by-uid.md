## gf dashboards get-dashboard-by-uid

Gets dashboard by uid

### Synopsis

Gets dashboard by uid

Optional query parameter `apiVersion` selects the Kubernetes API version used to load the dashboard first (for example `v1beta1`). If that request fails, the default version is used instead. When omitted, only the default is used.

Will return the dashboard given the dashboard unique identifier (uid).

Use: /apis/dashboards.grafana.app/v1/namespaces/{ns}/dashboards/{uid}

Response schema (GetDashboardByUIDOK.Payload):
  dashboard                                        object
  meta                                             object
  meta.annotationsPermissions                      object
  meta.annotationsPermissions.dashboard            object
  meta.annotationsPermissions.dashboard.canAdd     boolean
  meta.annotationsPermissions.dashboard.canDelete  boolean
  meta.annotationsPermissions.dashboard.canEdit    boolean
  meta.apiVersion                                  string
  meta.canAdmin                                    boolean
  meta.canDelete                                   boolean
  meta.canEdit                                     boolean
  meta.canSave                                     boolean
  meta.canStar                                     boolean
  meta.created                                     string
  meta.createdBy                                   string
  meta.expires                                     string
  meta.folderId                                    number   Deprecated: use FolderUID instead
  meta.folderTitle                                 string
  meta.folderUid                                   string
  meta.folderUrl                                   string
  meta.hasAcl                                      boolean
  meta.isFolder                                    boolean
  meta.isSnapshot                                  boolean
  meta.isStarred                                   boolean
  meta.provisioned                                 boolean
  meta.provisionedExternalId                       string
  meta.publicDashboardEnabled                      boolean
  meta.slug                                        string
  meta.type                                        string
  meta.updated                                     string
  meta.updatedBy                                   string
  meta.url                                         string
  meta.version                                     number

```
gf dashboards get-dashboard-by-uid [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for get-dashboard-by-uid
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --uid string                     UID [required]
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

* [gf dashboards](gf_dashboards.md)	 - Dashboards API

