## gf recording-rules



```
gf recording-rules [flags]
```

### Options

```
  -h, --help   help for recording-rules
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
* [gf recording-rules create-recording-rule](gf_recording-rules_create-recording-rule.md)	 - Creates a recording rule that is then registered and started
* [gf recording-rules create-recording-rule-write-target](gf_recording-rules_create-recording-rule-write-target.md)	 - Creates a remote write target
* [gf recording-rules delete-recording-rule](gf_recording-rules_delete-recording-rule.md)	 - Deletes removes the rule from the registry and stops it
* [gf recording-rules delete-recording-rule-write-target](gf_recording-rules_delete-recording-rule-write-target.md)	 - Deletes the remote write target
* [gf recording-rules get-recording-rule-write-target](gf_recording-rules_get-recording-rule-write-target.md)	 - Returns the prometheus remote write target
* [gf recording-rules list-recording-rules](gf_recording-rules_list-recording-rules.md)	 - Lists all rules in the database active or deleted
* [gf recording-rules test-create-recording-rule](gf_recording-rules_test-create-recording-rule.md)	 - Tests a recording rule
* [gf recording-rules update-recording-rule](gf_recording-rules_update-recording-rule.md)	 - Updates the active status of a rule

