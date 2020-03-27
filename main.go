package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/yanzay/tbot/v2"
)

func main() {
	var msg, link string
	token := os.Getenv("INPUT_TOKEN")
	chat := os.Getenv("INPUT_CHAT")
	status := os.Getenv("INPUT_STATUS")
	message := os.Getenv("INPUT_MESSAGE")
	repo := os.Getenv("GITHUB_REPOSITORY")
	commit := os.Getenv("GITHUB_SHA")
	actor := os.Getenv("GITHUB_ACTOR")

	if token == "" || chat == "" || status == "" {
		log.Fatal("token input is required")
	}

	client := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	link = fmt.Sprintf("https://github.com/%s/commit/%s/checks", repo, commit)

	if message != "" {
		msg = fmt.Sprintf(`*%s*: %s, %s ([%s](%s))`, strings.ToUpper(status), actor, message, repo, link)
	} else {
		msg = fmt.Sprintf(`*%s*: %s ([%s](%s))`, strings.ToUpper(status), actor, repo, link)
	}

	_, err := client.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
