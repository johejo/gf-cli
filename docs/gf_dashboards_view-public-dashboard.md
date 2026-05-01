## gf dashboards view-public-dashboard

Get public dashboard for view

### Synopsis

Get public dashboard for view

Response schema (ViewPublicDashboardOK.Payload):
  dashboard                                           object
  meta                                                object
  meta.annotationsPermissions                         object
  meta.annotationsPermissions.dashboard               object
  meta.annotationsPermissions.dashboard.canAdd        boolean
  meta.annotationsPermissions.dashboard.canDelete     boolean
  meta.annotationsPermissions.dashboard.canEdit       boolean
  meta.annotationsPermissions.organization            object
  meta.annotationsPermissions.organization.canAdd     boolean
  meta.annotationsPermissions.organization.canDelete  boolean
  meta.annotationsPermissions.organization.canEdit    boolean
  meta.apiVersion                                     string
  meta.canAdmin                                       boolean
  meta.canDelete                                      boolean
  meta.canEdit                                        boolean
  meta.canSave                                        boolean
  meta.canStar                                        boolean
  meta.created                                        string
  meta.createdBy                                      string
  meta.expires                                        string
  meta.folderId                                       number   Deprecated: use FolderUID instead
  meta.folderTitle                                    string
  meta.folderUid                                      string
  meta.folderUrl                                      string
  meta.hasAcl                                         boolean
  meta.isFolder                                       boolean
  meta.isSnapshot                                     boolean
  meta.isStarred                                      boolean
  meta.provisioned                                    boolean
  meta.provisionedExternalId                          string
  meta.publicDashboardEnabled                         boolean
  meta.slug                                           string
  meta.type                                           string
  meta.updated                                        string
  meta.updatedBy                                      string
  meta.url                                            string
  meta.version                                        number

```
gf dashboards view-public-dashboard [flags]
```

### Options

```
      --access-token string            AccessToken
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for view-public-dashboard
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

* [gf dashboards](gf_dashboards.md)	 - Dashboards API

