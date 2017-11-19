package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	p "github.com/somepackage/playlist"

	"gopkg.in/yaml.v2"
)

var (
	youtubeSearchString = flag.String("q", "google", "Search String")
	maxSearchResult     = flag.Int64("m", 4, "Max number of search results")
)

func main() {

	config := new(p.Config)
	configPath := os.Getenv("CONFIG")
	if configPath == "" {
		fmt.Println("CONFIG file path not specifed. Guided steps to get tokens")
		config = p.GuidedTourToGetToken(config)
	} else {
		yamlData, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Fatalf("Error:loadingConfigYaml:%v", err)
		}
		err = yaml.Unmarshal(yamlData, config)
		if err != nil {
			log.Fatalf("Error:Unmarshalling:%v", err)
		}
	}

	playlists, err := p.GetSpotifyPlaylist(config.Spotify.Username, config.Spotify.Token)
	if err != nil {
		log.Fatalf("Error:GetSpotifyPlaylist:%v", err)
	}

	for _, playlist := range playlists.Playlists {
		fmt.Println(fmt.Sprintf("playlistName: %s\n playlistUrl: %s\n===============================", playlist.PlaylistName, playlist.PlaylistUrl))
		trackItems, _ := p.GetTrackList(playlist.PlaylistName, playlist.PlaylistUrl, config.Spotify.Token)
		for _, track := range trackItems.Items {
			t := track.Track
			fmt.Println(fmt.Sprintf(" trackName: %s\n  \n albumName: %s\n ",
				t.TrackName, t.Album.Name))
			searchString := fmt.Sprintf("%s - %s", t.TrackName, t.Album.Name)
			p.SearchYoutubeByKey(searchString, 5, config.Youtube.Token)
		}
	}

}
