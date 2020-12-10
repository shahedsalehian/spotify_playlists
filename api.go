package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func getPlaylists(accessToken string) []Playlist {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.spotify.com/v1/users/shahedsal/playlists?limit=50", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var result PlaylistItems
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result.PlaylistItems
}

func getPlaylist(playlistURI string, accessToken string) []Track {
	client := &http.Client{}

	req, err := http.NewRequest("GET", playlistURI, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var result Tracks
	err = json.Unmarshal(bodyBytes, &result)
	if err != nil {
		log.Fatal(err)
	}

	return result.Tracks
}

func login() string {
	secrets := readSecrets()
	clientID := secrets[0]
	clientSecret := secrets[1]

	data := url.Values{"grant_type": {"client_credentials"}}
	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	secret := fmt.Sprintf("%s:%s", clientID, clientSecret)
	base64Secret := base64.StdEncoding.EncodeToString([]byte(secret))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64Secret))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var m auth
	err = json.Unmarshal(bodyBytes, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m.AccessToken
}

func readSecrets() []string {
	clientID, err := ioutil.ReadFile("client_id")
	if err != nil {
		log.Fatal(err)
	}
	clientSecret, err := ioutil.ReadFile("client_secret")
	if err != nil {
		log.Fatal(err)
	}

	return []string{string(clientID), string(clientSecret)}
}
