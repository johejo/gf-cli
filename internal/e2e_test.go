package internal_test

import (
	"strings"
	"testing"

	"github.com/johejo/gf-cli/internal"
)

func Test(t *testing.T) {
	cmd := internal.RootCmd()
	setArgs := func(t *testing.T) {
		name := t.Name()
		name = strings.ReplaceAll(name, "Test/", "")
		name = strings.ReplaceAll(name, "_", " ")
		t.Helper()
		cmd.SetArgs(strings.Split(name, " "))
		t.Cleanup(func() { cmd.SetArgs([]string{}) })
	}

	t.Run("health get-health", func(t *testing.T) {
		setArgs(t)
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("dashboards get-home-dashboard", func(t *testing.T) {
		t.Setenv("GF_BASIC_AUTH_USERNAME", "admin")
		t.Setenv("GF_BASIC_AUTH_PASSWORD", "asdf1234")
		setArgs(t)
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("users search-users", func(t *testing.T) {
		t.Setenv("GF_BASIC_AUTH_USERNAME", "admin")
		t.Setenv("GF_BASIC_AUTH_PASSWORD", "asdf1234")
		setArgs(t)
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("admin admin-get-settings", func(t *testing.T) {
		t.Setenv("GF_BASIC_AUTH_USERNAME", "admin")
		t.Setenv("GF_BASIC_AUTH_PASSWORD", "asdf1234")
		setArgs(t)
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("admin admin-get-stats", func(t *testing.T) {
		t.Setenv("GF_BASIC_AUTH_USERNAME", "admin")
		t.Setenv("GF_BASIC_AUTH_PASSWORD", "asdf1234")
		setArgs(t)
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("org-preferences get-org-preferences", func(t *testing.T) {
		t.Setenv("GF_BASIC_AUTH_USERNAME", "admin")
		t.Setenv("GF_BASIC_AUTH_PASSWORD", "asdf1234")
		setArgs(t)
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}
	})
}
