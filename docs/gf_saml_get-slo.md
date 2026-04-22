## gf saml get-slo

Its performs single logout s l o callback

### Synopsis

Its performs single logout s l o callback

There might be two possible requests: 1. Logout response (callback) when Grafana initiates single logout and IdP returns response to logout request. 2. Logout request when another SP initiates single logout and IdP sends logout request to the Grafana, or in case of IdP-initiated logout.

```
gf saml get-slo [flags]
```

### Options

```
  -h, --help   help for get-slo
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

* [gf saml](gf_saml.md)	 - 

