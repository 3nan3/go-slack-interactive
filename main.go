package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/slack-go/slack"
)

var (
	verificationToken = os.Getenv("VERIFICATION_TOKEN")

	baseAttachment = slack.Attachment{
		Color:      "#f9a41b",
		CallbackID: "beer",
	}
)

func main() {
	http.Handle("/suimasen", slashCommandHandler{acceptableTokens: []string{verificationToken}})
	http.Handle("/interactive-endpoint", interactionHandler{acceptableTokens: []string{verificationToken}})

	log.Println("[INFO] Server listening")
	http.ListenAndServe(":3000", nil)
}

func newAttachment() *slack.Attachment {
	a := baseAttachment
	return &a
}

func responseMessage(w http.ResponseWriter, attachment slack.Attachment) {
	params := &slack.Msg{
		ReplaceOriginal: true,
		Attachments: []slack.Attachment{attachment},
	}
	b, err := json.Marshal(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
