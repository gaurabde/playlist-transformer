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

	configPath := os.Getenv("CONFIG")
	if configPath == "" {
		log.Fatal("Please set CONFIG path for yaml config")
	}

	yamlData, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalln("Error: loading file")
	}

	config := new(p.Config)
	err = yaml.Unmarshal(yamlData, config)
	if err != nil {
		log.Fatalln("Error: Unmarshalling")
	}

	playlists, err := p.GetSpotifyPlaylist("gaurabde", config.SpotifyToken)
	if err != nil {
		log.Fatal(err)
	}

	for _, playlist := range playlists.Playlists {
		fmt.Println(fmt.Sprintf("playlistName: %s\n playlistUrl: %s\n===============================", playlist.PlaylistName, playlist.PlaylistUrl))
		trackItems, _ := p.GetTrackList(playlist.PlaylistName, playlist.PlaylistUrl, config.SpotifyToken)
		for _, track := range trackItems.Items {
			t := track.Track
			fmt.Println(fmt.Sprintf(" trackName: %s\n  \n albumName: %s\n ",
				t.TrackName, t.Album.Name))
			searchString := fmt.Sprintf("%s - %s", t.TrackName, t.Album.Name)
			p.SearchYoutubeByKey(searchString, 5, config.YoutubeToken)
		}
	}

}
