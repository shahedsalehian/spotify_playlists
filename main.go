package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	accessToken := login()
	playlists := getPlaylists(accessToken)
	if _, err := os.Stat("playlists_output"); os.IsNotExist(err) {
		os.Mkdir("playlists_output", os.ModeDir)
	}

	for _, playlist := range playlists {
		playlistData := getPlaylist(playlist.PlaylistTracksDetail.Href, accessToken)
		file, err := os.Create(fmt.Sprintf("playlists_output/%s.csv", playlist.Name))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Creating %s.csv\n", playlist.Name)
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.Write([]string{"Artist", "Track"})
		for _, v := range playlistData {
			err := writer.Write([]string{v.Track.Artists[0].Name, v.Track.Name})
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
