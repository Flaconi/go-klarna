package go_klarna

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	testingMux    *http.ServeMux
	testingServer *httptest.Server
)

func TestClient_Get(t *testing.T) {

}

func setupServer() {
	testingMux = http.NewServeMux()
	testingServer = httptest.NewServer(testingMux)
}

func setupMux(
	assertion *assert.Assertions,
	path string,
	request interface{},
	expectedMethod string,
	mockedResponse interface{},
) {
	testingMux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		assertion.Equal(expectedMethod, r.Method, "HTTP Method is not as expected")

		receivedReqBytes := make([]byte, r.ContentLength)
		r.Body.Read(receivedReqBytes)
		if nil != request {
			sentRequestBytes, _ := json.Marshal(request)
			assertion.Equal(sentRequestBytes, receivedReqBytes)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockedResponse)
	})
}

func tearDown() {
	testingServer.Close()
}

func testingClient() *client {
	uri, _ := url.Parse(testingServer.URL)
	return &client{
		Config{
			BaseURL:     uri,
			APIPassword: "somePass",
			APIUsername: "someUser",
		},
		http.DefaultClient,
	}
}
