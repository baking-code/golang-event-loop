package apiclient

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/endpoint1" {
			t.Errorf("Expected to request '/endpoint1', got: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"value":"fixed"}`))
	}))
	defer server.Close()
	var resp, err = Get(server.URL + "/endpoint1")
	assert.Equal(t, map[string]interface{}{"value": "fixed"}, resp)
	assert.Nil(t, err)
}

func TestPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/endpoint1" {
			t.Errorf("Expected to request '/endpoint1', got: %s", r.URL.Path)
		}
		resp, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(resp))
		}
	}))
	defer server.Close()
	var resp, err = Post(server.URL+"/endpoint1", map[string]interface{}{"value": "fixed"})
	assert.Equal(t, map[string]interface{}{"value": "fixed"}, resp)
	assert.Nil(t, err)
}
