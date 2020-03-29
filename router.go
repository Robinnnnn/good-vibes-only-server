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
	router.HandleFunc("/oauth", s.handleOAuthRedirect).Methods("GET").Queries("code", "{code}", "state", "{state}")
	router.HandleFunc("/refresh", s.handleTokenRefresh).Methods("POST")

	routerWithLogs := handlers.LoggingHandler(os.Stdout, router)
	routerWithCORS := handlers.CORS(
		handlers.AllowedOrigins([]string{s.appConfig.WebBaseURL, s.appConfig.BaseURL}),
		handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(routerWithLogs)

	return routerWithCORS
}

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("💙")
	utils.RespondWithBody(w, "ok")
}
