package main

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractParamIn(t *testing.T) {
	t.Parallel()

	src := `package fake

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

type FakeParams struct {
	Body         *string
	TeamID       int64
	UserID       int64
	Page         *int64
	XHeader      *string
	XRequiredHdr string
}

func (o *FakeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if err := r.SetPathParam("team_id", swag.FormatInt64(o.TeamID)); err != nil {
		return err
	}

	if err := r.SetPathParam("user_id", swag.FormatInt64(o.UserID)); err != nil {
		return err
	}

	if o.Page != nil {
		var qrPage int64
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.XHeader != nil {
		if err := r.SetHeaderParam("X-Header", *o.XHeader); err != nil {
			return err
		}
	}

	if err := r.SetHeaderParam("X-Required-Hdr", o.XRequiredHdr); err != nil {
		return err
	}

	return nil
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "fake.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}

	got := extractParamIn(f, "FakeParams")

	want := map[string]string{
		"Body":         "body",
		"TeamID":       "path",
		"UserID":       "path",
		"Page":         "query",
		"XHeader":      "header",
		"XRequiredHdr": "header",
	}

	for k, v := range want {
		if got[k] != v {
			t.Errorf("In[%q] = %q, want %q", k, got[k], v)
		}
	}
	for k, v := range got {
		if _, ok := want[k]; !ok {
			t.Errorf("unexpected entry In[%q] = %q", k, v)
		}
	}
}

func TestExtractParamIn_NoWriteToRequest(t *testing.T) {
	t.Parallel()

	src := `package fake

type LonelyParams struct {
	Foo string
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "fake.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}

	got := extractParamIn(f, "LonelyParams")
	if len(got) != 0 {
		t.Errorf("expected empty map, got %v", got)
	}
}

func TestExtractParamIn_WrongStruct(t *testing.T) {
	t.Parallel()

	src := `package fake

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type OtherParams struct{}

func (o *OtherParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetPathParam("x", o.X); err != nil {
		return err
	}
	return nil
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "fake.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("ParseFile: %v", err)
	}

	got := extractParamIn(f, "TargetParams")
	if len(got) != 0 {
		t.Errorf("expected empty map for wrong struct, got %v", got)
	}
}
