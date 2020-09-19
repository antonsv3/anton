package anton

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Method used to send the Master to Anton
func (master Master) SendToAnton(antonLocation string) {

	// We want to send the account we're following, and not the Master Account, if the Account Type is "Agent"
	if master.AccountType == "Agent" {
		// We can just take the first Line's UserName and Pass because there shouldn't be multiple lines from different
		// users at the same time
		master.MasterName = master.MasterLines[0].MasterName
		master.MasterPass = master.MasterLines[0].MasterPass
	}

	requestBody, err := json.Marshal(master)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(antonLocation, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
