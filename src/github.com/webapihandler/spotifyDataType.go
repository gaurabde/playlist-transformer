package main

type Playlist struct {
	PlaylistName string `json:"name"`
	PlaylistId   string `json:"id"`
	PlaylistUrl  string `json:"href"`
}

type Playlists struct {
	Playlists []Playlist `json:"items"`
}

type TracksItems struct {
	Items []Tracks `json:"items"`
}

type Tracks struct {
	Track Track `json:"track"`
}

type Track struct {
	TrackName string   `json:"name"`
	TrackUrl  string   `json:"href"`
	Album     Album    `json:"album"`
	Artists   []Artist `json:"artists"`
}

type Artist struct {
	Name string `json:"name"`
	Url  string `json:"href"`
}

type Album struct {
	AlbumType string `json:"album_type"`
	Name      string `json:"name"`
}
