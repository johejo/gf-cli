## gf provisioning put-alert-rule-group



### Synopsis

Body schema (AlertRuleGroup):
{
  "folderUid": string,
  "interval": number,
  "rules": [
    {
      "annotations": {"key": string},
      "condition": string,
      "data": [object],
      "execErrState": string,
      "folderUID": string,
      "for": string,
      "id": number,
      "isPaused": boolean,
      "keep_firing_for": string,
      "labels": {"key": string},
      "missingSeriesEvalsToResolve": number,
      "noDataState": string,
      "notification_settings": object,
      "orgID": number,
      "provenance": string,
      "record": object,
      "ruleGroup": string,
      "title": string,
      "uid": string,
      "updated": string
    }
  ],
  "title": string
}
  rules[].condition        required
  rules[].data             required
  rules[].execErrState     required, enum: OK | Alerting | Error
  rules[].folderUID        required
  rules[].for              required
  rules[].noDataState      required, enum: Alerting | NoData | OK
  rules[].orgID            required
  rules[].ruleGroup        rule group Max Length: 190 Min Length: 1, required
  rules[].title            title Max Length: 190 Min Length: 1, required
  rules[].uid              uid Max Length: 40 Min Length: 1
  rules[].updated          updated Read Only: true

```
gf provisioning put-alert-rule-group [flags]
```

### Options

```
      --body string                   The path to the body json file or json string. For example, --body=/path/to/body.json or --body='{"foo": "bar"}'
      --describe-body-jsonschema      Print the JSON Schema of the request body and exit without calling the API
      --folder-uid string             FolderUID
      --group string                  Group
  -h, --help                          help for put-alert-rule-group
      --raw                           Print the raw HTTP response body instead of the decoded payload
      --x-disable-provenance string   XDisableProvenance
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

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

