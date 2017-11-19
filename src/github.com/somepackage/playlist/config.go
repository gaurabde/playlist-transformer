package playlist

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config struct {
	Spotify struct {
		DevApiUrl string `yaml:"dev_api_url"`
		Username  string `yaml:"username"`
		Token     string `yaml:"token"`
	}
	Youtube struct {
		Username string `yaml:"username"`
		Token    string `yaml:"token"`
	}
}

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
