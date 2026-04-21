## gf access-control create-role

Creates a new custom role

### Synopsis

Creates a new custom role and maps given permissions to that role. Note that roles with the same prefix as Fixed Roles can’t be created.

You need to have a permission with action `roles:write` and scope `permissions:type:delegate`. `permissions:type:delegate` scope ensures that users can only create custom roles with the same, or a subset of permissions which the user has. For example, if a user does not have required permissions for creating users, they won’t be able to create a custom role which allows to do that. This is done to prevent escalation of privileges.

Body schema (CreateRoleForm):
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
  "uid": string,
  "version": number
}

```
gf access-control create-role [flags]
```

### Options

```
      --body string   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
  -h, --help          help for create-role
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

