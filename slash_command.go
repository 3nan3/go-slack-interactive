package main

import (
	"log"
	"net/http"

	"github.com/slack-go/slack"
)

type slashCommandHandler struct {
	acceptableTokens []string
}

func (h slashCommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slashCommand, err := slack.SlashCommandParse(r)
	if err != nil {
		log.Printf("[ERROR] Failed to parse request: %v", r)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !slashCommand.ValidateToken(h.acceptableTokens...) {
		log.Printf("[ERROR] Invalid token: %s", slashCommand.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	attachment := *newAttachment()
	attachment.Text = "Which beer do you want? :beer:"
	attachment.Actions = []slack.AttachmentAction{
		{
			Name: "select",
			Type: "select",
			Options: []slack.AttachmentActionOption{
				{
					Text:  "Asahi Super Dry",
					Value: "Asahi Super Dry",
				},
				{
					Text:  "Kirin Lager Beer",
					Value: "Kirin Lager Beer",
				},
				{
					Text:  "Sapporo Black Label",
					Value: "Sapporo Black Label",
				},
				{
					Text:  "Suntory Malts",
					Value: "Suntory Malts",
				},
				{
					Text:  "Yona Yona Ale",
					Value: "Yona Yona Ale",
				},
			},
		},
		{
			Name:  "cancel",
			Text:  "Cancel",
			Type:  "button",
			Style: "danger",
		},
	}
	responseMessage(w, attachment)
}