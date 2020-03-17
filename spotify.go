package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Robinnnnn/good-vibes-only-server.git/utils"
	"github.com/gorilla/mux"
	"github.com/zmb3/spotify"
)

const (
	stateCookieKey    = "spotify_auth_state"
	playlistCookieKey = "spotify_playlist_id"
)

type spotifyController struct {
	appConfig  *applicationConfig
	authClient *spotify.Authenticator
	httpClient *http.Client
}

type authGrantResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int16  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

func newSpotifyController(
	appConfig *applicationConfig,
	authClient *spotify.Authenticator,
	httpClient *http.Client,
) *spotifyController {
	return &spotifyController{
		appConfig,
		authClient,
		httpClient,
	}
}

func initSpotifyAuthClient(config *applicationConfig) *spotify.Authenticator {
	redirectURL := fmt.Sprintf("%s/oauth", config.BaseURL)
	auth := spotify.NewAuthenticator(redirectURL, spotify.ScopeUserReadPrivate)
	auth.SetAuthInfo(config.SpotifyClientID, config.SpotifySecretKey)
	return &auth
}

// Redirects the user to login with Spotify. State cookie is used to verify OAuth handshake.
func (s spotifyController) handleLogin(w http.ResponseWriter, r *http.Request) {
	stateHash := utils.RandomString(50)
	utils.AddCookie(w, stateCookieKey, stateHash)

	vars := mux.Vars(r)
	// TODO: bake playlistId into the state hash, to be decrypted
	if id, ok := vars["playlistId"]; ok {
		utils.AddCookie(w, playlistCookieKey, id)
	}

	// Link where user will sign into their Spotify account
	// e.g. https://accounts.spotify.com/authorize?client_id=48...
	authURL := s.authClient.AuthURL(stateHash)
	http.Redirect(w, r, authURL, http.StatusFound)
}

// Users are redirected here from a successful Spotify login.
// Spotify will attach an auth code which we'll use to authorize ourselves on behalf of the user.
func (s spotifyController) HandleOAuthRedirect(w http.ResponseWriter, r *http.Request) {
	// Ensure state cookie has persisted
	stateCookie, err := r.Cookie(stateCookieKey)
	if stateCookie == nil {
		http.Error(w, "State not found", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "State not found", http.StatusBadRequest)
		return
	}
	utils.RemoveCookie(w, stateCookieKey)

	vars := mux.Vars(r)
	code := vars["code"]
	state := vars["state"]

	// Verify state query param against cookie
	stateQueryValue := state
	if stateCookie.Value != stateQueryValue {
		http.Error(w, "State value mismatch", http.StatusBadRequest)
		return
	}

	// Get user auth tokems with code
	body, err := s.getUserAuthTokens(code)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	returnURL := fmt.Sprintf("%s/oauth", s.appConfig.WebBaseURL)
	utils.AddCookie(w, "access-token", body.AccessToken)
	utils.AddCookie(w, "refresh-token", body.RefreshToken)
	http.Redirect(w, r, returnURL, http.StatusFound)
}

func (s spotifyController) getUserAuthTokens(code string) (*authGrantResponse, error) {
	// Set form fields with auth info
	formData := url.Values{}
	formData.Set("code", code)
	formData.Set("redirect_uri", fmt.Sprintf("%s/oauth", s.appConfig.BaseURL))
	formData.Set("grant_type", "authorization_code")
	data := bytes.NewBufferString(formData.Encode())

	// Fire POST request
	req, err := http.NewRequest(http.MethodPost, spotify.TokenURL, data)
	authHeader := fmt.Sprintf("%s:%s", s.appConfig.SpotifyClientID, s.appConfig.SpotifySecretKey)
	encodedAuthHeader := base64.StdEncoding.EncodeToString([]byte(authHeader))
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", encodedAuthHeader))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		resStr := buf.String()
		return nil, errors.New(resStr)
	}

	resBody := &authGrantResponse{}
	if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
		return nil, err
	}

	return resBody, nil
}
