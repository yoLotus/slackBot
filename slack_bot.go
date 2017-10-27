package main

import (
	"encoding/json"
        "fmt"
	"net/http"
	"os"

        "github.com/turnage/graw/reddit"
	"github.com/gorilla/mux"
)

var bot reddit.Bot

type imageURL struct {
	URL string `json:"image_url"`
}

type message struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
	Attachments  []imageURL `json:"attachments"`
}

func getFirstPostsFromSubreddit(w http.ResponseWriter, r *http.Request) {
	harvest, err := bot.Listing(fmt.Sprintf("/r/%s", r.FormValue("text")), "")

	if err != nil {
                fmt.Println(fmt.Sprintf("Failed to fetch /r/%s: ", r.FormValue("text")), err)
                return
        }

	w.Header().Set("Content-Type", "application/json")
	imagesURL := make([]imageURL, 5)
	for i, post := range harvest.Posts[:5] {
		imagesURL[i] = imageURL{URL: post.URL}
	}

	response := message{
		ResponseType: "in_channel",
		Text: fmt.Sprintf("All hot posts on %s subreddit", r.FormValue("text")),
		Attachments: imagesURL,
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	// bot initialization
	cfg := reddit.BotConfig{
		Agent: os.Getenv("user_agent"),
		// Your registered app info from following:
		// https://github.com/reddit/reddit/wiki/OAuth2
		App: reddit.App{
			ID: os.Getenv("client_id"),
			Secret: os.Getenv("client_secret"),
			Username: os.Getenv("username"),
			Password: os.Getenv("password"),
		},
	}
	_bot, err := reddit.NewBot(cfg)

        if err != nil {
                fmt.Println("Failed to create bot handle: ", err)
                return
        } else {
		bot = _bot
	}

	// mux router
	r := mux.NewRouter()

	r.HandleFunc("/", getFirstPostsFromSubreddit).Methods("POST")

	http.ListenAndServe(":" + os.Getenv("PORT"), r)
}
