package _select

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		expected := fastUrl
		actual, err := Racer(fastUrl, slowUrl)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if expected != actual {
			log.Fatalf("expected %v actual %v", expected, actual)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Millisecond)
		serverB := makeDelayedServer(25 * time.Millisecond)

		urlA := serverA.URL
		urlB := serverB.URL

		_, err := ConfigurableRacer(urlA, urlB, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(delay)
		writer.WriteHeader(http.StatusOK)
	}))
}
