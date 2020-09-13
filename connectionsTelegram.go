package anton

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func SendTelegram(msg, groupID, tokenID string) {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// This is the URI:
	postURL := "https://api.telegram.org/bot{tokenID}/sendMessage"

	// Create New Post Body for the message to send
	postBody := new(TelegramMsg)

	// Format the body with the GroupID and the Message
	postBody.ChatID, _ = strconv.ParseInt(groupID, 10, 64)
	postBody.Text = msg

	// Format the URL with the static parameters
	finalURL := helper.ReplaceParameters(postURL, "{tokenID}", tokenID)

	// Fix encoding on the JSON body from the struct before sending
	postBodyEncoded, err := json.Marshal(postBody)

	// Send a post request with your token
	res, err := http.Post(finalURL, "application/json", bytes.NewBuffer(postBodyEncoded))
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("unexpected status" + res.Status)
	}
}
