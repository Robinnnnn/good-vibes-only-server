package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Robinnnnn/good-vibes-only-server.git/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func initializeRouter(s *spotifyController) http.Handler {
	router := mux.NewRouter()

	// Standard app routes
	router.HandleFunc("/alive", handleHealthcheck)

	// Spotify routes
	router.HandleFunc("/login", s.handleLogin).Methods("GET").Queries("playlistId", "{playlistId}")
	router.HandleFunc("/login", s.handleLogin).Methods("GET")
	router.HandleFunc("/oauth", s.HandleOAuthRedirect).Methods("GET").Queries("code", "{code}", "state", "{state}")

	loggedRouter := handlers.LoggingHandler(os.Stdout, router)

	return loggedRouter
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("💙")
	utils.RespondWithBody(w, "ok")
}
