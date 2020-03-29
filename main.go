package main

import (
	"fmt"
	"net/http"

	"github.com/Robinnnnn/good-vibes-only-server.git/utils"
)

func main() {
	appConfig := getApplicationConfig()

	httpClient := utils.CreateHTTPClient()

	authClient := initSpotifyAuthClient(appConfig)

	controller := newSpotifyController(
		appConfig,
		authClient,
		httpClient,
	)

	router := initializeRouter(controller)

	http.ListenAndServe(
		fmt.Sprintf(":%s", appConfig.Port),
		router,
	)
}
