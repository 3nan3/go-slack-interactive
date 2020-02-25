package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/slack-go/slack"
)

type interactionHandler struct {
	acceptableTokens []string
}

func (h interactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("[ERROR] Invalid method: %s", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("[ERROR] Failed to read request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonStr, err := url.QueryUnescape(string(buf)[8:])
	if err != nil {
		log.Printf("[ERROR] Failed to unespace request body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message slack.InteractionCallback
	if err := json.Unmarshal([]byte(jsonStr), &message); err != nil {
		log.Printf("[ERROR] Failed to decode json message from slack: %s", jsonStr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !h.validateToken(message.Token) {
		log.Printf("[ERROR] Invalid token: %s", message.Token)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	attachment := *newAttachment()

	action := *message.ActionCallback.AttachmentActions[0]
	switch action.Name {
	case "select":
		value := action.SelectedOptions[0].Value
		attachment.Title = fmt.Sprintf("OK to order %s ?", strings.Title(value))
		attachment.Actions = []slack.AttachmentAction{
			{
				Name:  "start",
				Text:  "Yes",
				Type:  "button",
				Value: "start",
				Style: "primary",
			},
			{
				Name:  "cancel",
				Text:  "No",
				Type:  "button",
				Style: "danger",
			},
		}
		responseMessage(w, attachment)
		return
	case "start":
		title := ":ok: your order was submitted! yay!"
		attachment.Fields = []slack.AttachmentField{
			{
				Title: title,
				Value: "",
				Short: false,
			},
		}
		responseMessage(w, attachment)
		return
	case "cancel":
		title := fmt.Sprintf(":x: @%s canceled the request", message.User.Name)
		attachment.Fields = []slack.AttachmentField{
			{
				Title: title,
				Value: "",
				Short: false,
			},
		}
		responseMessage(w, attachment)
		return
	default:
		log.Printf("[ERROR] ]Invalid action was submitted: %s", action.Name)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h interactionHandler) validateToken(token string) bool {
	for _, t := range h.acceptableTokens {
		if token == t {
			return true
		}
	}
	return false
}
