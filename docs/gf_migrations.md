## gf migrations

Migrations API

```
gf migrations [flags]
```

### Options

```
  -h, --help   help for migrations
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
* [gf migrations cancel-snapshot](gf_migrations_cancel-snapshot.md)	 - Cancels a snapshot wherever it is in its processing chain
* [gf migrations create-cloud-migration-token](gf_migrations_create-cloud-migration-token.md)	 - Creates gcom access token
* [gf migrations create-session](gf_migrations_create-session.md)	 - Creates a migration session
* [gf migrations create-snapshot](gf_migrations_create-snapshot.md)	 - Triggers the creation of an instance snapshot associated with the provided session
* [gf migrations delete-cloud-migration-token](gf_migrations_delete-cloud-migration-token.md)	 - Deletes a cloud migration token
* [gf migrations delete-session](gf_migrations_delete-session.md)	 - Deletes a migration session by its uid
* [gf migrations get-cloud-migration-token](gf_migrations_get-cloud-migration-token.md)	 - Fetches the cloud migration token if it exists
* [gf migrations get-resource-dependencies](gf_migrations_get-resource-dependencies.md)	 - Gets the resource dependencies graph for the current set of migratable resources
* [gf migrations get-session](gf_migrations_get-session.md)	 - Gets a cloud migration session by its uid
* [gf migrations get-session-list](gf_migrations_get-session-list.md)	 - Gets a list of all cloud migration sessions that have been created
* [gf migrations get-shapshot-list](gf_migrations_get-shapshot-list.md)	 - Gets a list of snapshots for a session
* [gf migrations get-snapshot](gf_migrations_get-snapshot.md)	 - Gets metadata about a snapshot including where it is in its processing and final results
* [gf migrations upload-snapshot](gf_migrations_upload-snapshot.md)	 - Uploads a snapshot to the grafana migration service for processing

