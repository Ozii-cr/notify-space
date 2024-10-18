package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// Create a test server to mock the Space webhook
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	// Set environment variables
	os.Setenv("GITHUB_EVENT_NAME", "push")
	os.Setenv("GITHUB_REPOSITORY", "test/repo")
	os.Setenv("INPUT_SPACE_WEBHOOK_URL", server.URL)
	os.Setenv("INPUT_MESSAGE_TYPE", "plain")

	// Run the main function
	main()

}
