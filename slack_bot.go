package main

import (
        "fmt"

        "github.com/turnage/graw/reddit"
)

func main() {
        bot, err := reddit.NewBotFromAgentFile("slack_bot.agent", 0)
        if err != nil {
                fmt.Println("Failed to create bot handle: ", err)
                return
        }

        harvest, err := bot.Listing("/r/pic", "")
        if err != nil {
                fmt.Println("Failed to fetch /r/pic: ", err)
                return
        }

        for _, post := range harvest.Posts[:2] {
		fmt.Printf("[%s] posted [%s] -> %s\n", post.Author, post.Title, post.URL)
        }
}
