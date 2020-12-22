package anton

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Method used to send the Master to Anton
func (master Master) SendToAnton(antonLocation, MongoURI, DevTelegramGroupID, AntonTelegramBot string) {

	// We want to send the account we're following, and not the Master Account, if the Account Type is "Agent"
	if master.AccountType == "Agent" {
		// We can just take the first Line's UserName and Pass because there shouldn't be multiple lines from different
		// users at the same time
		master.MasterName = master.MasterLines[0].MasterName
		master.MasterPass = master.MasterLines[0].MasterPass
	}

	// Send the process id and create hash to authenticate to the Slave server
	client := GetClient(MongoURI)
	currentProcess := GatherScrapingProcess(client, DevTelegramGroupID, AntonTelegramBot)
	currentProcessHash := CreateProcessHash(currentProcess.CurrentID, currentProcess.Salt)
	DisconnectClient(client)

	master.Hash = currentProcessHash

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

// Method to create hash with string and salt, returns the hash
func CreateProcessHash(processID, salt string) string {

	// Concatenate the string using salt with ProcessID
	saltedString := salt + processID

	// Now we will use SHA256 for the hashing this string
	hashByte := sha256.Sum256([]byte(saltedString))

	// Return the HashBytes as string
	return string(hashByte[:])

}
