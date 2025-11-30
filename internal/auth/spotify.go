package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"spotui/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type CallbackResult struct {
	AccessToken  string
	Error 		 error
}

type AccessTokenData struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
} 


var scope = "user-read-playback-state+user-modify-playback-state+user-read-currently-playing+user-library-read+user-top-read+user-library-read+user-follow-read+user-read-private+playlist-read-private+user-read-currently-playing+user-read-recently-played"

func createOauthLink(clientID string, scope string, redirect_uri string) string {
	return fmt.Sprintf("https://accounts.spotify.com/authorize?client_id=%s&response_type=code&redirect_uri=%s&scope=%s", clientID, url.QueryEscape(redirect_uri), scope) 
}

func StartCallbackSever() (<-chan CallbackResult, func()) {
	utils.OpenURL(createOauthLink(utils.Current.Spotify.ClientID, scope, "http://127.0.0.1:8000/callback"))
	r := gin.New()
	resultChan := make(chan CallbackResult, 1)

	r.GET("/callback", func (c *gin.Context)  {
		err := c.Query("error")
		code := c.Query("code")

		if err != "" {
			resultChan <- CallbackResult{Error: fmt.Errorf("%s", err)}
			c.String(http.StatusBadRequest, "Authorization failed %s", err)
		}

		if code == "" {
			resultChan <- CallbackResult{Error: fmt.Errorf("no code provided")}
			c.String(http.StatusBadRequest, "No code provided")
		}

		if err == "" && code != "" {
			c.String(http.StatusOK, "Authentication successful. You can safely close this window.")
			resp := requestAccessToken(code)
			resultChan <- CallbackResult{AccessToken: resp.AccessToken}
		}
	})

	srv := &http.Server{
		Addr: ":8000",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %d\n", &err)
		}
	}()

	stop := func ()  {
		_ = srv.Shutdown(context.Background())
	}

	return resultChan, stop
}

func requestAccessToken(code string) AccessTokenData {
	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", "http://127.0.0.1:8000/callback")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp := utils.SendSpotifyReq(req)	
	
	var accessTokenData AccessTokenData
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&accessTokenData)
		if err != nil {
			log.Fatal(err)
		}	
		
		dataJson, _ := json.MarshalIndent(accessTokenData, "", "\t")
		err = os.WriteFile("config/spotify.json", dataJson, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	resp.Body.Close()
	return accessTokenData
	
}

func readRefreshToken() string {
	reader, _ := os.ReadFile("config/spotify.json")
	
	var accessTokenData AccessTokenData
	err := json.Unmarshal(reader, &accessTokenData)
	if err != nil {
		log.Fatal(err)
	}
	return accessTokenData.RefreshToken
}

func RefreshAccessToken() string {
	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", readRefreshToken())

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	resp := utils.SendSpotifyReq(req)
	
	var accessTokenData AccessTokenData
	if resp.StatusCode == http.StatusOK {
		err := json.NewDecoder(resp.Body).Decode(&accessTokenData)
		if err != nil {
			log.Fatal(err)
		}	
		
		dataJson, _ := json.MarshalIndent(accessTokenData, "", "\t")
		err = os.WriteFile("config/spotify.json", dataJson, 0644)
		if err != nil {
			log.Fatal(err)
		}
	}
	resp.Body.Close()
	return accessTokenData.AccessToken
}