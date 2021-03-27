package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {

	const numberOfWins = 3

	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	for i := 0; i < numberOfWins; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	}

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), fmt.Sprintf("%v", numberOfWins))

}
