package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const spotifyApiURL = "https://api.spotify.com/v1/users"

var token = "BQD5JWcL3ZokSPeGB4q8bzs8Hs4EORwQdgN2A43pA6y9INtj2aY-Bzx_LqiekJY2sjw6RAuKwANEFvdn65Cq95rQLu2kU0uVq0f57PrGVbSmP20D4dtatqEIDd38wf-pPtaKXYRkoLyPZOAfoAvAOJ0j5Cuq"

func getPlaylist(userId string, token string) (*Playlists, error) {
	fmt.Println("Getting Playlists ...")
	playlistUrl := fmt.Sprintf("%s/%s/playlists", spotifyApiURL, userId)
	jsonData, err := GetUrlResponse(playlistUrl, token)
	var playlist = new(Playlists)

	if err != nil {
		fmt.Errorf(fmt.Sprintf("Playlist response error: %s", err.Error()))
		return playlist, err
	}

	err = json.Unmarshal(jsonData, &playlist)
	//fmt.Println(playlist, len(playlist.Playlists))
	return playlist, nil
}

func getTrackList(playlistName, playlistUrl, token string) (*TracksItems, error) {

	fmt.Println("Getting Tracks from playlist: ", playlistName)
	tracksUrl := fmt.Sprintf("%s/tracks", playlistUrl)
	jsonData, err := GetUrlResponse(tracksUrl)
	var tracks = new(TracksItems)

	if err != nil {
		fmt.Errorf(fmt.Sprintf("Playlist response error: %s", err.Error()))
		return tracks, err
	}

	err = json.Unmarshal(jsonData, &tracks)
	//fmt.Println(tracks, len(tracks.Items))
	return tracks, nil
}

func main() {
	playlists, err := getPlaylist("gaurabde", token)
	if err != nil {
		log.Fatal(err)
	}

	for _, playlist := range playlists.Playlists {
		fmt.Println(fmt.Sprintf("playlistName: %s\n playlistUrl: %s\n===============================", playlist.PlaylistName, playlist.PlaylistUrl))
		trackItems, _ := getTrackList(playlist.PlaylistName, playlist.PlaylistUrl, token)
		for _, track := range trackItems.Items {
			t := track.Track
			fmt.Println(fmt.Sprintf("trackName: %s\n  trackUrl: %s\n albumName: %s\n artistList: %s",
				t.TrackName, t.TrackUrl, t.Album.Name, t.Artists))
		}
	}

}
