package playlist

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const spotifyApiURL = "https://api.spotify.com/v1/users"

func GuidedTourToGetToken(config *Config) *Config {
	fmt.Println("Configs loaded...\n" +
		"\t Open following link in browser and get the Authorization Token with following access:\n" +
		"\t\t 1. playlist-read-private\n" +
		"\t\t 2. user-library-read \n\n")
	if config.Spotify.DevApiUrl == "" {
		config.Spotify.DevApiUrl = "https://developer.spotify.com/web-api/console/"
	}
	fmt.Println("\t\t", config.Spotify.DevApiUrl)

	fmt.Println("\n\n Please enter spotify token generated from above instructions: ")
	r := bufio.NewReader(os.Stdin)
	token, _ := r.ReadString('\n')
	config.Spotify.Token = strings.TrimSuffix(token, "\n")

	fmt.Println("Spotify Username: ")
	username, _ := r.ReadString('\n')
	config.Spotify.Username = strings.TrimSuffix(username, "\n")

	return config
}

func GetSpotifyPlaylist(userId string, token string) (*Playlists, error) {
	fmt.Println("Getting Playlists ...")
	playlistUrl := fmt.Sprintf("%s/%s/playlists", spotifyApiURL, userId)
	respData, err := GetUrlResponse(playlistUrl, token)
	var playlist = new(Playlists)

	if err != nil {
		fmt.Println(fmt.Sprintf("Playlist response error: %s", err.Error()))
		return playlist, err
	}

	err = json.Unmarshal(respData, &playlist)
	//fmt.Println(playlist, len(playlist.Playlists))
	return playlist, nil
}

func GetTrackList(playlistName, playlistUrl, token string) (*TracksItems, error) {

	fmt.Println("Getting Tracks from playlist: ", playlistName)
	tracksUrl := fmt.Sprintf("%s/tracks", playlistUrl)
	jsonData, err := GetUrlResponse(tracksUrl, token)
	var tracks = new(TracksItems)

	if err != nil {
		fmt.Println(fmt.Sprintf("Playlist response error: %s", err.Error()))
		return tracks, err
	}

	err = json.Unmarshal(jsonData, &tracks)
	//fmt.Println(tracks, len(tracks.Items))
	writePlaylistData(stripfySpace(fmt.Sprintf("/tmp/%s.json", playlistName)), tracks)
	return tracks, nil
}

func writePlaylistData(playlistName string, playListData *TracksItems) {
	trackJsonData, _ := json.Marshal(playListData)
	_ = ioutil.WriteFile(playlistName, trackJsonData, 0644)
	fmt.Println("Playlist:%s:jsonAdded", playlistName)
}

func stripfySpace(dataString string) string {
	return strings.Join(strings.Fields(dataString), "")
}
