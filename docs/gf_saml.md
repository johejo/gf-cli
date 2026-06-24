## gf saml

Saml API

```
gf saml [flags]
```

### Options

```
  -h, --help   help for saml
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

* [gf](gf.md)	 - CLI for Grafana API
* [gf saml get-metadata](gf_saml_get-metadata.md)	 - Its exposes the s p grafana s metadata for the Id p s consumption
* [gf saml get-saml-logout](gf_saml_get-saml-logout.md)	 - Gets logout initiates single logout process
* [gf saml get-slo](gf_saml_get-slo.md)	 - Its performs single logout s l o callback
* [gf saml post-acs](gf_saml_post-acs.md)	 - Its performs assertion consumer service a c s
* [gf saml post-slo](gf_saml_post-slo.md)	 - Its performs single logout s l o callback

