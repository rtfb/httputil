package httputil

import (
	"net/http"
	"testing"
)

func TestExtractReferer(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com", nil)
	if err != nil {
		t.Error(err)
	}
	refURL := ExtractReferer(req)
	if refURL != "" {
		t.Fatalf("Empty string expected, but <%s> found!", refURL)
	}
	req.Header.Add("Referer", "foo/bar/baz")
	refURL = ExtractReferer(req)
	if refURL != "baz" {
		t.Fatalf("<baz> expected, but <%s> found!", refURL)
	}
}

func TestHostAndPort(t *testing.T) {
	data := []struct {
		host string
		port string
		exp  string
	}{
		{"host", "666", "host:666"},
		{"host:", "111", "host:111"},
		{"hoost", ":1234", "hoost:1234"},
		{"hoist:", ":987", "hoist:987"},
		{"h", "", "h"},
		{"h", "::", "h"},
	}
	for _, d := range data {
		got := JoinHostAndPort(d.host, d.port)
		if got != d.exp {
			t.Errorf("JoinHostAndPort(%q, %q) = %q, but expected %q",
				d.host, d.port, got, d.exp)
		}
	}
}
