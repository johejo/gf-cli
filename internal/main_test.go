package internal

import "testing"

func Test_parseSchemeAndHost(t *testing.T) {
	tests := []struct {
		arg    string
		scheme string
		host   string
	}{
		{
			arg:    "https://play.grafana.org",
			scheme: "https",
			host:   "play.grafana.org",
		},
		{
			arg:    "https://localhost:443",
			scheme: "https",
			host:   "localhost:443",
		},
		{
			arg:    "http://localhost:3000",
			scheme: "http",
			host:   "localhost:3000",
		},
		{
			arg:    "localhost:8080",
			scheme: "http",
			host:   "localhost:8080",
		},
		{
			arg:    "play.grafana.org",
			scheme: "https",
			host:   "play.grafana.org",
		},
		{
			arg:    "play.grafana.org:1443",
			scheme: "https",
			host:   "play.grafana.org:1443",
		},
		{
			arg:    "192.168.100.11:8000",
			scheme: "http",
			host:   "192.168.100.11:8000",
		},
	}
	for _, test := range tests {
		t.Run(test.arg, func(t *testing.T) {
			scheme, host, err := parseSchemeAndHost(test.arg)
			if err != nil {
				t.Fatal(err)
			}
			if scheme != test.scheme {
				t.Errorf("unexpected scheme: want=%s, got=%s", test.scheme, scheme)
			}
			if host != test.host {
				t.Errorf("unexpected host: want=%s, got=%s", test.host, host)
			}
		})
	}
}
