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

	if token == "" || chat == "" {
		log.Fatal("one or more of the required parameters is empty")
	}

	client := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	link = fmt.Sprintf("https://github.com/%s/commit/%s/checks", repo, commit)
	if status != "" {
		msg = fmt.Sprintf(`*%s*: %s [%s](%s)`, strings.ToUpper(status), message, commit, link)
	} else {
		msg = fmt.Sprintf(`%s [%s](%s)`, message, commit, link)
	}

	_, err := client.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
