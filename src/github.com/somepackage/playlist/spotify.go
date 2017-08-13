package playlist

import (
	"encoding/json"
	"fmt"
)

const spotifyApiURL = "https://api.spotify.com/v1/users"

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
	return tracks, nil
}
