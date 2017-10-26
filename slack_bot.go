package main

import (
        "fmt"
	"net/http"
	"os"

        "github.com/turnage/graw/reddit"
	"github.com/gorilla/mux"
)

var bot reddit.Bot

func getFirstUrlFromUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	harvest, err := bot.Listing(fmt.Sprintf("/r/%s", vars["subreddit"]), "")

	if err != nil {
                fmt.Println(fmt.Sprintf("Failed to fetch /r/%s: ", vars["subreddit"]), err)
                return
        }

        for _, post := range harvest.Posts[:1] {
		fmt.Fprintf(w, "%s", post.URL)
        }
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

	r.HandleFunc("/{subreddit}", getFirstUrlFromUrl).Methods("GET")

	http.ListenAndServe(":" + os.Getenv("PORT"), r)
}
