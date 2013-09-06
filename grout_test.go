package grout

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRouteMux(t *testing.T) {
	r := NewRouteMux()
	if r == nil {
		t.Errorf("NewRouteMux should not have returned nil")
	}
}

func TestRoute(t *testing.T) {
	r := NewRouteMux()
	handler := func(w http.ResponseWriter, r *http.Request, m map[string]string) {}
	r.Route("hello", handler)
	if r.routes[0] == nil {
		t.Errorf("Handler did not get registered properly")
	}
}

func TestServeHTTP(t *testing.T) {
	r := NewRouteMux()
	var b bytes.Buffer
	handler := func(w http.ResponseWriter, r *http.Request, m map[string]string) {
		for k, v := range m {
			fmt.Fprintf(&b, "%s: %s\n", k, v)
		}
	}
	r.Route(`/blogs/(?P<name>[a-z][a-z_-]+[a-z])/(?P<othername>[0-9]+)`, handler)
	req, _ := http.NewRequest("GET", "http://localhost/blogs/some-crummy-blog/235", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if rr.Code == 404 {
		t.Errorf("We actually found the page...")
	}
	if b.String() != "name: some-crummy-blog\nothername: 235\n" {
		t.Errorf("Did not get variable name nor value")
	}
}

func Test404(t *testing.T) {
	r := NewRouteMux()
	handler := func(w http.ResponseWriter, r *http.Request, m map[string]string) {}
	r.Route(`/blogs/(?P<name>[a-z][a-z_-]+[a-z])/(?P<othername>[0-9]+)`, handler)
	req, _ := http.NewRequest("GET", "http://localhost/no+match/here", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if rr.Code != 404 {
		t.Errorf("response should be 404")
	}
}
