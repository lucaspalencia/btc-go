package httpclient

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHttpClient(t *testing.T) {
	t.Run("With successfull request", func(t *testing.T) {
		expected := []byte(`{"bitcoinc":{"usd": 1000, "usd_24h_change": 25.05, "last_updated_at": 500432}}`)

		testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Write(expected)
		}))

		defer testServer.Close()

		httpClient := New(testServer.URL)

		body, err := httpClient.Get()

		if err != nil {
			t.Error("httpClient.Get() returns error")
		}

		if reflect.DeepEqual(expected, body) != true {
			t.Error("httpClient.Get() returns body error")
		}
	})

	t.Run("With error request", func(t *testing.T) {
		httpClient := New("http://localhost")

		_, err := httpClient.Get()

		if err == nil {
			t.Error("httpClient.Get() not return error")
		}
	})
}
