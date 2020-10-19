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

func GatherFrontEndUsers(client *mongo.Client, filter bson.M) []FrontEndUser {
	var antonUsers []FrontEndUser
	collection := client.Database("Anton").Collection("FrontEndUsers")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Fatal("Error on Finding all the documents", err)
	}

	for cur.Next(context.TODO()) {
		var antonUser FrontEndUser
		err = cur.Decode(&antonUser)
		if err != nil {
			log.Fatal("Error on Decoding the document", err)
		}
		antonUsers = append(antonUsers, antonUser)
	}

	// Close the Cursor
	err = cur.Close(context.TODO())
	if err != nil {
		log.Fatal("[#GatherSiteDictionaries] Failed to Close Connection", err)
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

// Master Method to push all MasterLines to MongoDB
func (master Master) PushMasterLines(MongoURI, TelegramGroupID, TelegramToken string) {

	var helper Helper

	masterHeader := "----------- Master -----------"
	SendTelegram(masterHeader, TelegramGroupID, TelegramToken)

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

		// Now we can send the Telegram Msg within this loop
		SendTelegram(telegramMsg, TelegramGroupID, TelegramToken)
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
func (returnSlaveResults SlaveResults) PushSlaveLines(MongoURI, TelegramGroupID, TelegramToken string) {

	var helper Helper

	/*
		----------- Master -----------
		------------ Slave -----------
		------------- End ------------
	*/

	// We are first going to send all the 'Placed Lines' first
	placedLinesString := helper.ReplaceParameters("Placed Lines: {Amount}", "{Amount}", strconv.Itoa(len(returnSlaveResults.PlacedLines)))
	skippedLinesString := helper.ReplaceParameters("Skipped Lines: {Amount}", "{Amount}", strconv.Itoa(len(returnSlaveResults.SkippedLines)))
	errorLinesString := helper.ReplaceParameters("Error Lines: {Amount}", "{Amount}", strconv.Itoa(len(returnSlaveResults.ErrorLines)))

	slaveHeader := "------------ Slave -----------"
	slaveEnd := "------------- End ------------"

	// First, Send the Header for Placed Lines
	SendTelegram(slaveHeader, TelegramGroupID, TelegramToken)
	SendTelegram(placedLinesString, TelegramGroupID, TelegramToken)

	// Send all the Placed Lines
	for i := range returnSlaveResults.PlacedLines {

		// First format the Master and Lines, will need the telegram message to look like:

		// {Master} ({MasterPass}) - Following #{TicketID}
		// LineType: {LineType}

		// If it is a Total, The bottom lines need to look like this
		// {Team} {OverUnder} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
		// {Sport} - {League}

		// If it is a Spread, The bottom lines need to look like this
		// {Team} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
		// {Sport} - {League}

		line := returnSlaveResults.PlacedLines[i]

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

		// Now we can send the Telegram Msg within this loop
		SendTelegram(telegramMsg, TelegramGroupID, TelegramToken)

	}

	// Next, Send the Headers for Skipped and Error Lines
	SendTelegram(skippedLinesString, TelegramGroupID, TelegramToken)
	SendTelegram(errorLinesString, TelegramGroupID, TelegramToken)
	SendTelegram(slaveEnd, TelegramGroupID, TelegramToken)

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
