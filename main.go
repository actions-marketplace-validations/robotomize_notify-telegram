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
	var msg string
	token := os.Getenv("INPUT_TOKEN")
	chat := os.Getenv("INPUT_CHAT")
	status := os.Getenv("INPUT_STATUS")
	message := os.Getenv("INPUT_MESSAGE")

	if token == "" || chat == "" {
		log.Fatal("one or more of the required parameters is empty")
	}

	client := tbot.NewClient(token, http.DefaultClient, "https://api.telegram.org")

	if status != "" {
		msg = fmt.Sprintf(`*%s*: %s `, strings.ToUpper(status), message)
	} else {
		msg = fmt.Sprintf(`%s `, message)
	}

	_, err := client.SendMessage(chat, msg, tbot.OptParseModeMarkdown)
	if err != nil {
		log.Fatalf("unable to send message: %v", err)
	}
}
