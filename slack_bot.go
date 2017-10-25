package main

import (
        "fmt"
	"net/http"

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
	_bot, err := reddit.NewBotFromAgentFile("slack_bot.agent", 0)
        if err != nil {
                fmt.Println("Failed to create bot handle: ", err)
                return
        } else {
		bot = _bot
	}

	// mux router
	r := mux.NewRouter()

	r.HandleFunc("/{subreddit}", getFirstUrlFromUrl).Methods("GET")

	http.ListenAndServe(":7777", r)
}
