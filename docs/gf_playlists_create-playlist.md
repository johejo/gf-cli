## gf playlists create-playlist

Creates playlist

### Synopsis

Creates playlist

Please refer to [new API](?api=playlist.grafana.app-v1).

Body schema (CreatePlaylistCommand):
  interval            string
  items               array<object>
  items[].Id          number
  items[].PlaylistId  number
  items[].order       number
  items[].title       string
  items[].type        string
  items[].value       string
  name                string

Response schema (CreatePlaylistOK.Payload):
  id        number
  interval  string
  name      string
  uid       string

```
gf playlists create-playlist [flags]
```

### Options

```
      --body string                    Request body JSON or path to a JSON file (e.g. --body=/path/to/body.json, --body='{"foo": "bar"}'). The 'Body schema' block below lists fields and types; use --describe-body-jsonschema for a strict JSON Schema. [required]
      --describe-body-jsonschema       Print the JSON Schema of the request body and exit without calling the API
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for create-playlist
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

* [gf playlists](gf_playlists.md)	 - Playlists API

