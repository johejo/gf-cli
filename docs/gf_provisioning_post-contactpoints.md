## gf provisioning post-contactpoints

Creates a contact point

### Synopsis

Creates a contact point

Body schema (EmbeddedContactPoint):
  disableResolveMessage  boolean
  name                   string   Name is used as grouping key in the UI. Contact points with the
                                  same name will be grouped in the UI.
  provenance             string
  settings               object   REQUIRED
  type                   string   REQUIRED
                                  enum: alertmanager | dingding | discord | email | googlechat | kafka | line | opsgenie | pagerduty | pushover | sensugo | slack | teams | telegram | threema | victorops | webhook | wecom
  uid                    string   UID is the unique identifier of the contact point. The UID can be
                                  set by the user.

Response schema (PostContactpointsAccepted.Payload):
  disableResolveMessage  boolean
  name                   string   Name is used as grouping key in the UI. Contact points with the
                                  same name will be grouped in the UI.
  provenance             string
  settings               object   REQUIRED
  type                   string   REQUIRED
                                  enum: alertmanager | dingding | discord | email | googlechat | kafka | line | opsgenie | pagerduty | pushover | sensugo | slack | teams | telegram | threema | victorops | webhook | wecom
  uid                    string   UID is the unique identifier of the contact point. The UID can be
                                  set by the user.

```
gf provisioning post-contactpoints [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for post-contactpoints
      --raw                            Print the raw HTTP response body instead of the decoded payload
      --x-disable-provenance string    XDisableProvenance
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf provisioning](gf_provisioning.md)	 - Provisioning API

