package main

import (
	"encoding/base64"
	"net/http"
	"testing"
)

// TestProbeBlobMetadataIncludesSHA256 asserts that blob metadata (returned by
// POST /blobs and GET /blobs/{name}) carries the lowercase-canonical sha256 of
// the stored content, not an empty string.
func TestProbeBlobMetadataIncludesSHA256(t *testing.T) {
	_, h := newTestServer()
	code, body := doJSON(t, h, http.MethodPost, "/blobs", putBlobRequest{
		Name:    "a",
		Content: base64.StdEncoding.EncodeToString([]byte("abc")),
	})
	if code != http.StatusOK {
		t.Fatalf("put code = %d body=%v", code, body)
	}
	// sha256("abc") known vector.
	want := "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad"
	if got, _ := body["sha256"].(string); got != want {
		t.Fatalf("put sha256 = %q want %q", got, want)
	}
	code, body = doJSON(t, h, http.MethodGet, "/blobs/a", nil)
	if code != http.StatusOK {
		t.Fatalf("get code = %d", code)
	}
	if got, _ := body["sha256"].(string); got != want {
		t.Fatalf("get sha256 = %q want %q", got, want)
	}
}
