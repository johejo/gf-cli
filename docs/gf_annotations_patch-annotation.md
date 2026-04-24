## gf annotations patch-annotation

Patches annotation

### Synopsis

Patches annotation

Updates one or more properties of an annotation that matches the specified ID. This operation currently supports updating of the `text`, `tags`, `time` and `timeEnd` properties. This is available in Grafana 6.0.0-beta2 and above.

Body schema (PatchAnnotationsCmd):
{
  "data": any,  // models.JSON
  "id": number,
  "tags": [string],
  "text": string,
  "time": number,
  "timeEnd": number
}

```
gf annotations patch-annotation [flags]
```

### Options

```
      --annotation-id string           AnnotationID
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block is a type reference; use --describe-body-jsonschema for a strict JSON Schema.
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for patch-annotation
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

* [gf annotations](gf_annotations.md)	 - Annotations API

