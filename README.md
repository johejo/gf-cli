# gf-cli

Grafana API Client for command line operations with shell completions

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
