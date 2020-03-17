package main

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	appEnv           = "APP_ENV"
	baseURL          = "BASE_URL"
	webBaseURL       = "WEB_BASE_URL"
	port             = "PORT"
	spotifyClientID  = "SPOTIFY_CLIENT_ID"
	spotifySecretKey = "SPOTIFY_SECRET_KEY"
)

type applicationConfig struct {
	AppEnv           string
	BaseURL          string
	WebBaseURL       string
	Port             string
	SpotifyClientID  string
	SpotifySecretKey string
}

func getApplicationConfig() *applicationConfig {
	setDefaultConfig()

	loadEnv()

	config := &applicationConfig{
		AppEnv:           viper.GetString(appEnv),
		BaseURL:          viper.GetString(baseURL),
		WebBaseURL:       viper.GetString(webBaseURL),
		Port:             viper.GetString(port),
		SpotifyClientID:  viper.GetString(spotifyClientID),
		SpotifySecretKey: viper.GetString(spotifySecretKey),
	}

	if config.SpotifyClientID == "" {
		panic("Spotify client ID is missing!")
	}
	if config.SpotifySecretKey == "" {
		panic("Spotify secret key is missing!")
	}

	return config
}

func setDefaultConfig() {
	viper.SetDefault(appEnv, "local")
	viper.SetDefault(baseURL, "http://localhost:3001")
	viper.SetDefault(webBaseURL, "http://localhost:3000")
	viper.SetDefault(port, "3001")
}

func loadEnv() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Error reading in env file: %s", err))
	}

	if _, ok := viper.Get(spotifyClientID).(string); !ok {
		panic("Spotify client ID is missing!")
	}
	if _, ok := viper.Get(spotifySecretKey).(string); !ok {
		panic("Spotify secret key is missing!")
	}
}
