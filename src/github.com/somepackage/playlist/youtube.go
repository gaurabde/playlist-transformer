package playlist

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

func SearchYoutubeByKey(searchString string, maxSearchResult int64, token string) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: token},
	}

	service, err := youtube.New(client)
	handleError(err)

	call := service.Search.List("id,snippet").
		Q(searchString).
		MaxResults(maxSearchResult)

	resp, err := call.Do()
	handleError(err)

	//fmt.Println(resp)
	fmt.Println("\n    ------youtube- ", searchString)
	for _, item := range resp.Items {
		fmt.Println(item.Snippet.Title)
		if item.Id.Kind == "youtube#video" {
			fmt.Println(item.Id.VideoId, item.Snippet.Title)
		}
	}

}

func handleError(err error) {
	if err != nil {
		log.Fatal(fmt.Sprintf("Error creating client: %v", err))
	}
}
