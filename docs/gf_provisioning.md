## gf provisioning

Provisioning API

```
gf provisioning [flags]
```

### Options

```
  -h, --help   help for provisioning
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
* [gf provisioning delete-alert-rule](gf_provisioning_delete-alert-rule.md)	 - Deletes a specific alert rule by UID
* [gf provisioning delete-alert-rule-group](gf_provisioning_delete-alert-rule-group.md)	 - Delete rule group
* [gf provisioning delete-contactpoints](gf_provisioning_delete-contactpoints.md)	 - Deletes a contact point
* [gf provisioning delete-mute-timing](gf_provisioning_delete-mute-timing.md)	 - Deletes a mute timing
* [gf provisioning delete-template](gf_provisioning_delete-template.md)	 - Deletes a notification template group
* [gf provisioning export-mute-timing](gf_provisioning_export-mute-timing.md)	 - Exports a mute timing in provisioning format
* [gf provisioning export-mute-timings](gf_provisioning_export-mute-timings.md)	 - Exports all mute timings in provisioning format
* [gf provisioning get-alert-rule](gf_provisioning_get-alert-rule.md)	 - Gets a specific alert rule by UID
* [gf provisioning get-alert-rule-export](gf_provisioning_get-alert-rule-export.md)	 - Exports an alert rule in provisioning file format
* [gf provisioning get-alert-rule-group](gf_provisioning_get-alert-rule-group.md)	 - Gets a rule group
* [gf provisioning get-alert-rule-group-export](gf_provisioning_get-alert-rule-group-export.md)	 - Exports an alert rule group in provisioning file format
* [gf provisioning get-alert-rules](gf_provisioning_get-alert-rules.md)	 - Gets all the alert rules
* [gf provisioning get-alert-rules-export](gf_provisioning_get-alert-rules-export.md)	 - Exports all alert rules in provisioning file format
* [gf provisioning get-contactpoints](gf_provisioning_get-contactpoints.md)	 - Gets all the contact points
* [gf provisioning get-contactpoints-export](gf_provisioning_get-contactpoints-export.md)	 - Exports all contact points in provisioning file format
* [gf provisioning get-mute-timing](gf_provisioning_get-mute-timing.md)	 - Gets a mute timing
* [gf provisioning get-mute-timings](gf_provisioning_get-mute-timings.md)	 - Gets all the mute timings
* [gf provisioning get-policy-tree](gf_provisioning_get-policy-tree.md)	 - Gets the notification policy tree
* [gf provisioning get-policy-tree-export](gf_provisioning_get-policy-tree-export.md)	 - Exports the notification policy tree in provisioning file format
* [gf provisioning get-template](gf_provisioning_get-template.md)	 - Gets a notification template group
* [gf provisioning get-templates](gf_provisioning_get-templates.md)	 - Gets all notification template groups
* [gf provisioning post-alert-rule](gf_provisioning_post-alert-rule.md)	 - Creates a new alert rule
* [gf provisioning post-contactpoints](gf_provisioning_post-contactpoints.md)	 - Creates a contact point
* [gf provisioning post-mute-timing](gf_provisioning_post-mute-timing.md)	 - Creates a new mute timing
* [gf provisioning put-alert-rule](gf_provisioning_put-alert-rule.md)	 - Updates an existing alert rule
* [gf provisioning put-alert-rule-group](gf_provisioning_put-alert-rule-group.md)	 - Creates or update alert rule group
* [gf provisioning put-contactpoint](gf_provisioning_put-contactpoint.md)	 - Updates an existing contact point
* [gf provisioning put-mute-timing](gf_provisioning_put-mute-timing.md)	 - Replaces an existing mute timing
* [gf provisioning put-policy-tree](gf_provisioning_put-policy-tree.md)	 - Sets the notification policy tree
* [gf provisioning put-template](gf_provisioning_put-template.md)	 - Updates an existing notification template group
* [gf provisioning reset-policy-tree](gf_provisioning_reset-policy-tree.md)	 - Clears the notification policy tree

