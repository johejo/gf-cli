## gf access-control update-role

Updates a custom role

### Synopsis

Updates a custom role

You need to have a permission with action `roles:write` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only create custom roles with the same, or a subset of permissions which the user has.

Body schema (UpdateRoleCommand):
{
  "description": string,
  "displayName": string,
  "global": boolean,
  "group": string,
  "hidden": boolean,
  "name": string,
  "permissions": [
    {
      "action": string,
      "created": string,
      "scope": string,
      "updated": string
    }
  ],
  "version": number
}
  description              required
  displayName              required
  group                    required

```
gf access-control update-role [flags]
```

### Options

```
      --body string       The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help              help for update-role
      --role-uid string   RoleUID
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

* [gf access-control](gf_access-control.md)	 - Access control API

