package anton

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
)

// Get the Client, will pass this Client into our other functions
func GetClient(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func DisconnectClient(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}

func GatherMasters(client *mongo.Client, filter bson.M) []Master {
	var masters []Master
	collection := client.Database("Anton").Collection("Masters")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var master Master
		err = cur.Decode(&master)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		masters = append(masters, master)
	}

	// Close the Cursor
	cur.Close(context.TODO())

	// Return Results
	return masters

}

func GatherSlaves(client *mongo.Client, filter bson.M) []Slave {
	var slaves []Slave
	collection := client.Database("Anton").Collection("Slaves")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var slave Slave
		err = cur.Decode(&slave)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		slaves = append(slaves, slave)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSlaves] Failed to Close Connection", err)
	}

	// Return Results
	return slaves
}

func GatherProfiles(client *mongo.Client, filter bson.M) []Profile {
	var profiles []Profile
	collection := client.Database("Anton").Collection("Profiles")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var profile Profile
		err = cur.Decode(&profile)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		profiles = append(profiles, profile)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherProfiles] Failed to Close Connection", err)
	}

	// Return Results
	return profiles

}

func GatherProxies(client *mongo.Client, filter bson.M) []Proxy {
	var proxies []Proxy
	collection := client.Database("Anton").Collection("Proxies")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var proxy Proxy
		err = cur.Decode(&proxy)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		proxies = append(proxies, proxy)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherProfiles] Failed to Close Connection", err)
	}

	// Return Results
	return proxies
}

func GatherSubnets(client *mongo.Client, filter bson.M) []Subnet {
	var subnets []Subnet
	collection := client.Database("Anton").Collection("Subnets")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var subnet Subnet
		err = cur.Decode(&subnet)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		subnets = append(subnets, subnet)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherProfiles] Failed to Close Connection", err)
	}

	// Return Results
	return subnets
}

func GatherSiteDictionaries(client *mongo.Client, filter bson.M) []SiteDictionary {
	var siteDictionaries []SiteDictionary
	collection := client.Database("Anton").Collection("SitesDictionary")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var siteDictionary SiteDictionary
		err = cur.Decode(&siteDictionary)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		siteDictionaries = append(siteDictionaries, siteDictionary)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSiteDictionaries] Failed to Close Connection", err)
	}

	// Return Results
	return siteDictionaries

}

func GatherAntonUsers(client *mongo.Client, filter bson.M) []AntonUser {
	var antonUsers []AntonUser
	collection := client.Database("Anton").Collection("AntonUsers")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var antonUser AntonUser
		err = cur.Decode(&antonUser)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		antonUsers = append(antonUsers, antonUser)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherAntonUsers] Failed to Close Connection", err)
	}

	// Return Results
	return antonUsers

}

func GatherMasterLines(client *mongo.Client, filter bson.M) []Lines {
	var lines []Lines
	collection := client.Database("Anton").Collection("MastersLines")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var line Lines
		err = cur.Decode(&line)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		lines = append(lines, line)
	}

	// Close the Cursor
	cur.Close(context.TODO())

	// Return Results
	return lines

}

func GatherSlaveLines(client *mongo.Client, filter bson.M) []Lines {
	var lines []Lines
	collection := client.Database("Anton").Collection("SlavesLines")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var line Lines
		err = cur.Decode(&line)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		lines = append(lines, line)
	}

	// Close the Cursor
	cur.Close(context.TODO())

	// Return Results
	return lines

}

func GatherSiteStatus(client *mongo.Client, filter bson.M) []SiteStatus {
	var siteStatuses []SiteStatus
	collection := client.Database("Anton").Collection("SiteStatus")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var siteStatus SiteStatus
		err = cur.Decode(&siteStatus)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		siteStatuses = append(siteStatuses, siteStatus)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSiteStatus] Failed to Close Connection", err)
	}

	// Return Results
	return siteStatuses
}

func GatherLinesMapping(client *mongo.Client, filter bson.M) []LinesMappingStruct {
	var returnLinesMapping []LinesMappingStruct
	collection := client.Database("Anton").Collection("LinesMapping")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var linesMapping LinesMappingStruct
		err = cur.Decode(&linesMapping)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		returnLinesMapping = append(returnLinesMapping, linesMapping)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherLinesMapping] Failed to Close Connection", err)
	}

	// Return Results
	return returnLinesMapping
}

func GatherScrapingProcess(client *mongo.Client, errorUserTelegram, antonBotTelegramTokenID string) Process {

	// Create a slice, although there should only be one Process running, which we'll check
	var currentProcessSlice []Process

	// Create the return process
	var returnProcess Process

	collection := client.Database("Anton").Collection("Connections")
	cur, err := collection.Find(context.TODO(), bson.M{"purpose": "masterscrapingprocessid"})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var process Process
		err = cur.Decode(&process)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		currentProcessSlice = append(currentProcessSlice, process)
	}

	// Check length of the results
	if len(currentProcessSlice) == 1 {
		returnProcess = currentProcessSlice[0]
	} else {
		SendTelegram("[#GatherScrapingProcessSalt] More/Less than one Current Master Scraping Process ID in Database", errorUserTelegram, antonBotTelegramTokenID)
	}

	return returnProcess
}

func GatherProxyIPGeoProcess(client *mongo.Client, errorUserTelegram, antonBotTelegramTokenID string) Process {

	// Create a slice, although there should only be one Process running, which we'll check
	var currentProcessSlice []Process

	// Create the return process
	var returnProcess Process

	collection := client.Database("Anton").Collection("Connections")
	cur, err := collection.Find(context.TODO(), bson.M{"purpose": "proxyipgeosyncingprocessid"})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var process Process
		err = cur.Decode(&process)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		currentProcessSlice = append(currentProcessSlice, process)
	}

	// Check length of the results
	if len(currentProcessSlice) == 1 {
		returnProcess = currentProcessSlice[0]
	} else {
		SendTelegram("[#GatherProxyProcessSalt] More/Less than one Current Master Proxy Sync Process ID in Database", errorUserTelegram, antonBotTelegramTokenID)
	}

	return returnProcess
}

func GatherProxyDatabaseProcess(client *mongo.Client, errorUserTelegram, antonBotTelegramTokenID string) Process {

	// Create a slice, although there should only be one Process running, which we'll check
	var currentProcessSlice []Process

	// Create the return process
	var returnProcess Process

	collection := client.Database("Anton").Collection("Connections")
	cur, err := collection.Find(context.TODO(), bson.M{"purpose": "proxydatabasesyncingprocessid"})
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}
	for cur.Next(context.TODO()) {
		var process Process
		err = cur.Decode(&process)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		currentProcessSlice = append(currentProcessSlice, process)
	}

	// Check length of the results
	if len(currentProcessSlice) == 1 {
		returnProcess = currentProcessSlice[0]
	} else {
		SendTelegram("[#GatherProxyProcessSalt] More/Less than one Current Master Proxy Sync Process ID in Database", errorUserTelegram, antonBotTelegramTokenID)
	}

	return returnProcess
}

// Master Method to push all MasterLines to MongoDB
func (master Master) PushMasterLines(MongoURI, AntonTelegramBot string) {

	var helper Helper

	var sentHeadersToTelegram []string

	masterHeader := "----------- Master -----------"

	// First format the Master and Lines, will need the telegram message to look like:

	// {Master} ({MasterPass}) #{TicketID}
	// LineType: {LineType}

	// If it is a Total, The bottom lines need to look like this
	// {Team} {OverUnder} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
	// {Sport} - {League}

	// If it is a Spread, The bottom lines need to look like this
	// {Team} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
	// {Sport} - {League}

	// We need to format each line for the Telegram Msg
	for i := range master.MasterLines {

		// If the Master is an Agent account, we want to send the User we are following, and not the Agent Login
		tempUserName := ""
		tempUserPass := ""

		if master.AccountType == "Agent" {
			tempUserName = master.MasterLines[i].MasterName
			tempUserPass = master.MasterLines[i].MasterPass
		} else {
			tempUserName = master.MasterName
			tempUserPass = master.MasterPass
		}

		lines := master.MasterLines[i]

		// First we need to send headers to all the Telegram Groups and save the Groups we did send it to
		if len(sentHeadersToTelegram) == 0 {
			SendTelegram(masterHeader, lines.AntonOwnerTelegram, AntonTelegramBot)
			sentHeadersToTelegram = append(sentHeadersToTelegram, lines.AntonOwnerTelegram)
		}

		// Now, if the string is not in the slice of saved Telegram Groups, we will resend the headers
		if helper.FindIfStringInSlice(lines.AntonOwnerTelegram, sentHeadersToTelegram) == "False" {
			SendTelegram(masterHeader, lines.AntonOwnerTelegram, AntonTelegramBot)
			sentHeadersToTelegram = append(sentHeadersToTelegram, lines.AntonOwnerTelegram)
		}

		// Start Formatting the Telegram Message

		telegramMsg := helper.ReplaceParameters("{Master} ({MasterPass}) #{TicketID}\n", "{Master}", tempUserName,
			"{MasterPass}", tempUserPass, "{TicketID}", lines.TicketID)
		telegramMsg += helper.ReplaceParameters("{LineType} - {Period}\n", "{Period}", lines.Period, "{LineType}",
			lines.LineType)

		// Based on the LineType, the next telegramMsg line will be formatted differently
		if lines.LineType == "Total" || lines.LineType == "TeamTotal" {
			telegramMsg += helper.ReplaceParameters("{Team} {OverUnder} {Spread} ({Juice}) [{RiskAmount}/{ToWin}]\n",
				"{Team}", lines.Team, "{OverUnder}", lines.OverUnder, "{Spread}", lines.LineSpread, "{Juice}",
				lines.LineJuice, "{RiskAmount}", lines.RiskAmount, "{ToWin}", lines.ToWinAmount)
		}

		// Based on the LineType, the next telegramMsg line will be formatted differently
		if lines.LineType == "MoneyLine" {
			telegramMsg += helper.ReplaceParameters("{Team} ({Juice}) [{RiskAmount}/{ToWin}]\n", "{Team}",
				lines.Team, "{Juice}", lines.LineJuice, "{RiskAmount}", lines.RiskAmount, "{ToWin}", lines.ToWinAmount)
		}

		// Based on the LineType, the next telegramMsg line will be formatted differently
		if lines.LineType == "Spread" {
			telegramMsg += helper.ReplaceParameters("{Team} {Spread} ({Juice}) [{RiskAmount}/{ToWin}]\n",
				"{Team}", lines.Team, "{Spread}", lines.LineSpread, "{Juice}", lines.LineJuice, "{RiskAmount}",
				lines.RiskAmount, "{ToWin}", lines.ToWinAmount)
		}

		// The last line on the TelegramMsg will be the same for the different leagues
		telegramMsg += helper.ReplaceParameters("{Sport} - {League}\n", "{Sport}", lines.Sport, "{League}",
			lines.League)

		// Now we can send the Telegram Msg within this loop, we will send each Master line to whoever the Master belongs to
		SendTelegram(telegramMsg, lines.AntonOwnerTelegram, AntonTelegramBot)
	}

	// Start database connections
	client := GetClient(MongoURI)
	results := client.Database("Anton").Collection("MastersLines")

	// This is how MongoDB needs to take the lines
	var linesToInsert []interface{}
	for i := range master.MasterLines {
		linesToInsert = append(linesToInsert, master.MasterLines[i])
	}

	results.InsertMany(context.Background(), linesToInsert)
	DisconnectClient(client)

}

// Master Method to push all MasterLines to MongoDB
func (returnSlaveResults SlaveResults) PushSlaveLines(MongoURI, AntonTelegramBot string) {

	var helper Helper

	slaveHeader := "------------ Slave -----------"
	slaveEnd := "------------- End ------------"

	var sentHeadersToTelegram []string

	// Send all the Placed Lines
	for i := range returnSlaveResults.PlacedLines {

		// Easier to read
		line := returnSlaveResults.PlacedLines[i]

		// First we need to send headers to all the Telegram Groups and save the Groups we did send it to
		if len(sentHeadersToTelegram) == 0 {

			// First, Send Slave Header
			SendTelegram(slaveHeader, line.AntonOwnerTelegram, AntonTelegramBot)

			// Next, Send Number of Placed Line for this Anton User

			countOfPlacedLines := 0

			for l := range returnSlaveResults.PlacedLines {
				if returnSlaveResults.PlacedLines[l].AntonOwnerTelegram == line.AntonOwnerTelegram {
					countOfPlacedLines = countOfPlacedLines + 1
				}
			}

			placedLinesString := helper.ReplaceParameters("Placed Lines: {Amount}", "{Amount}", strconv.Itoa(countOfPlacedLines))

			SendTelegram(placedLinesString, line.AntonOwnerTelegram, AntonTelegramBot)
			sentHeadersToTelegram = append(sentHeadersToTelegram, line.AntonOwnerTelegram)
		}

		// Now, if the string is not in the slice of saved Telegram Groups, we will resend the headers
		if helper.FindIfStringInSlice(line.AntonOwnerTelegram, sentHeadersToTelegram) == "False" {

			// First, Send Slave Header
			SendTelegram(slaveHeader, line.AntonOwnerTelegram, AntonTelegramBot)

			// Next, Send Number of Placed Line for this Anton User

			countOfPlacedLines := 0

			for l := range returnSlaveResults.PlacedLines {
				if returnSlaveResults.PlacedLines[l].AntonOwnerTelegram == line.AntonOwnerTelegram {
					countOfPlacedLines = countOfPlacedLines + 1
				}
			}

			placedLinesString := helper.ReplaceParameters("Placed Lines: {Amount}", "{Amount}", strconv.Itoa(countOfPlacedLines))

			SendTelegram(placedLinesString, line.AntonOwnerTelegram, AntonTelegramBot)
			sentHeadersToTelegram = append(sentHeadersToTelegram, line.AntonOwnerTelegram)
		}

		// Next, format the Master and Lines, will need the telegram message to look like:

		// {Master} ({MasterPass}) - Following #{TicketID}
		// LineType: {LineType}

		// If it is a Total, The bottom lines need to look like this
		// {Team} {OverUnder} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
		// {Sport} - {League}

		// If it is a Spread, The bottom lines need to look like this
		// {Team} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
		// {Sport} - {League}

		telegramMsg := helper.ReplaceParameters("Following #{TicketID}\n", "{TicketID}", line.MasterTicketID)

		telegramMsg += helper.ReplaceParameters("{SlaveName} ({SlavePass})\n", "{SlaveName}", line.SlaveName,
			"{SlavePass}", line.SlavePass)
		telegramMsg += helper.ReplaceParameters("{LineType} - {Period}\n", "{Period}", line.Period, "{LineType}",
			line.LineType)

		// Based on the LineType, the next telegramMsg line will be formatted differently
		if line.LineType == "Total" || line.LineType == "TeamTotal" {
			telegramMsg += helper.ReplaceParameters("{Team} {OverUnder} {Spread} ({Juice}) [{RiskAmount}/{ToWin}]\n",
				"{Team}", line.Team, "{OverUnder}", line.OverUnder, "{Spread}", line.LineSpread, "{Juice}",
				line.LineJuice, "{RiskAmount}", line.RiskAmount, "{ToWin}", line.ToWinAmount)
		}

		// Based on the LineType, the next telegramMsg line will be formatted differently
		if line.LineType == "MoneyLine" {
			telegramMsg += helper.ReplaceParameters("{Team} ({Juice}) [{RiskAmount}/{ToWin}]\n", "{Team}",
				line.Team, "{Juice}", line.LineJuice, "{RiskAmount}", line.RiskAmount, "{ToWin}", line.ToWinAmount)
		}

		// Based on the LineType, the next telegramMsg line will be formatted differently
		if line.LineType == "Spread" {
			telegramMsg += helper.ReplaceParameters("{Team} {Spread} ({Juice}) [{RiskAmount}/{ToWin}]\n",
				"{Team}", line.Team, "{Spread}", line.LineSpread, "{Juice}", line.LineJuice, "{RiskAmount}",
				line.RiskAmount, "{ToWin}", line.ToWinAmount)
		}

		// The last line on the TelegramMsg will be the same for the different leagues
		telegramMsg += helper.ReplaceParameters("{Sport} - {League}\n", "{Sport}", line.Sport, "{League}",
			line.League)

		// Now we can send the Telegram Msg within this loop, each line has an assigned telegram group
		SendTelegram(telegramMsg, line.AntonOwnerTelegram, AntonTelegramBot)

	}

	// Now we need to send the enders to each person we sent this to
	for i := range sentHeadersToTelegram {
		countOfSkippedLines := 0
		countOfErrorLines := 0

		for j := range returnSlaveResults.SkippedLines {
			if returnSlaveResults.SkippedLines[j].AntonOwnerTelegram == sentHeadersToTelegram[i] {
				countOfSkippedLines += 1
			}
		}

		skippedLinesString := helper.ReplaceParameters("Skipped Lines: {Amount}", "{Amount}", strconv.Itoa(countOfSkippedLines))
		SendTelegram(skippedLinesString, sentHeadersToTelegram[i], AntonTelegramBot)

		for k := range returnSlaveResults.ErrorLines {
			if returnSlaveResults.ErrorLines[k].AntonOwnerTelegram == sentHeadersToTelegram[i] {
				countOfErrorLines += 1
			}
		}

		errorLinesString := helper.ReplaceParameters("Error Lines: {Amount}", "{Amount}", strconv.Itoa(countOfErrorLines))
		SendTelegram(errorLinesString, sentHeadersToTelegram[i], AntonTelegramBot)

		// Send End
		SendTelegram(slaveEnd, sentHeadersToTelegram[i], AntonTelegramBot)

	}

	// Start database connections
	client := GetClient(MongoURI)
	results := client.Database("Anton").Collection("SlavesLines")

	// This is how MongoDB needs to take the Placed Lines
	var placedLinesToInsert []interface{}
	for i := range returnSlaveResults.PlacedLines {
		placedLinesToInsert = append(placedLinesToInsert, returnSlaveResults.PlacedLines[i])
	}
	results.InsertMany(context.Background(), placedLinesToInsert)

	// I also want to keep the Error Lines, sending it to the Error Log
	results = client.Database("Anton").Collection("Errors")
	var errorLinesToInsert []interface{}
	for i := range returnSlaveResults.ErrorLines {
		errorLinesToInsert = append(errorLinesToInsert, returnSlaveResults.ErrorLines[i])
	}
	results.InsertMany(context.Background(), errorLinesToInsert)

	// Disconnect MongoDB Client
	DisconnectClient(client)

}
