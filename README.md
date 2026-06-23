# gf-cli

Grafana API Client for command line operations.

Auto-generated from [grafana-openapi-client-go](https://github.com/grafana/grafana-openapi-client-go),
so it mirrors the **entire** Grafana HTTP API — every endpoint gets a typed
subcommand with flags and JSON in/out.

Originally built as a personal tool for poking at Grafana instances from the shell.

## vs. the official `gcx`

Grafana ships an official CLI, [`gcx`](https://github.com/grafana/gcx). For most
operational and Grafana Cloud workflows — querying metrics/logs/traces, alerting,
SLOs, IRM, observability-as-code — use that.

`gf-cli` stays useful where it doesn't overlap:

- **Full, typed API coverage** — including the long-tail admin/org/provisioning/
  licensing/SSO endpoints that `gcx` leaves to its raw `api` passthrough.
- **Older & OSS self-hosted Grafana** — no Grafana 12+ requirement.
- **Machine-readable schemas** — `--help-json` and
  `--describe-body-jsonschema` / `--describe-response-jsonschema` expose the full
  command surface and exact request/response JSON Schema, handy for generating
  agent tool definitions.

## Install

```
go install github.com/johejo/gf-cli/cmd/gf@latest
```

## Example

```
$ GF_HOST=play.grafana.org gf search search --limit 1 --tag grafanacloud
[
  {
    "folderId": 1298,
    "folderTitle": "GrafanaCloud",
    "folderUid": "tOFItmonk",
    "folderUrl": "/dashboards/f/tOFItmonk/GrafanaCloud",
    "id": 2252,
    "permanentlyDeleteDate": "0001-01-01T00:00:00.000Z",
    "tags": [
      "cardinality-management",
      "grafanacloud"
    ],
    "title": "Cardinality management - 1 - overview",
    "type": "dash-db",
    "uid": "cardinality-management",
    "uri": "db/cardinality-management-1-overview",
    "url": "/d/cardinality-management/cardinality-management-1-overview"
  }
]
```

## Documentation

[docs](./docs/gf.md)
