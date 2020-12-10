package main

type auth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
	Scope       string `json:"scope"`
}

// PlaylistItems are all Playlists
type PlaylistItems struct {
	PlaylistItems []Playlist `json:"items"`
}

// Playlist describes the Playlist object
type Playlist struct {
	Description          string               `json:"description"`
	Name                 string               `json:"name"`
	PlaylistTracksDetail PlaylistTracksDetail `json:"tracks"`
}

// Tracks contains the links and the total number of tracks
type PlaylistTracksDetail struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

type Tracks struct {
	Tracks []Track `json:"items"`
}

type Track struct {
	Track TrackDetail `json:"track"`
}

type TrackDetail struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artists"`
}

type Artist struct {
	Name string `json:"name"`
}
