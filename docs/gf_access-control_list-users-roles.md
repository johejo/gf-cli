## gf access-control list-users-roles

Lists roles assigned to multiple users

### Synopsis

Lists the roles that have been directly assigned to the given users. The list does not include built-in roles (Viewer, Editor, Admin or Grafana Admin), and it does not include roles that have been inherited from a team.

You need to have a permission with action `users.roles:read` and scope `users:id:*`.

Body schema (RolesSearchQuery):
{
  "includeHidden": boolean,
  "orgId": number,
  "teamIds": [number],
  "userIds": [number]
}

```
gf access-control list-users-roles [flags]
```

### Options

```
      --body string   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help          help for list-users-roles
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

* [gf access-control](gf_access-control.md)	 - 

