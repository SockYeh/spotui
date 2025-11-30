package utils

import (
	"encoding/base64"
	"log"
	"net/http"
)

func SendSpotifyReq(req *http.Request) *http.Response {
	spotifyB64 := base64.StdEncoding.EncodeToString([]byte(Current.Spotify.ClientID + ":" + Current.Spotify.ClientSecret))
	
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic " + spotifyB64)
	resp, err := http.DefaultClient.Do(req)
	
	if err != nil {
		log.Fatal(err)
	}
	return resp
}