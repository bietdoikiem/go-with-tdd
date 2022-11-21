package select_case

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return url of the server with the faster response time", func(t *testing.T) {

		slowServer := makeDelayTestServer(12 * time.Millisecond)
		fastServer := makeDelayTestServer(1 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q ", got, want)
		}
	})

	t.Run("return an error if a server doesn't respond within 10 seconds", func(t *testing.T) {
		slowServer := makeDelayTestServer(20 * time.Millisecond)
		fastServer := makeDelayTestServer(15 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		_, err := ConfigurableRacer(slowURL, fastURL, 10*time.Millisecond)

		if err == nil {
			t.Error("expect an error but didn't get one")
		}

	})
}

// makeDelayTestServers make a server with delay for unit testing purpose
func makeDelayTestServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
