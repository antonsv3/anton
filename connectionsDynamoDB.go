package anton

// fmt.Println("Hello")

// The different functions we have on this file are:

// - GatherMasters()
// - GatherOpenBets()
// - GatherSportsDict()
// - PushOpenBet()
// - RemoveOpenBet()
/*
import (
	"fmt"
	"os"
	"strings"

	"github.com/antonsv3/helper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// Sent API call to DynamoDB to gather Users, add them into overallWorkflow.Users slice of Users struct
func GatherMasters(status string) []Master {

	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Start creating Queries with these arguments
	tableName := "Masters"
	Status := status

	// Create the Expression to fill the input struct with.
	filt := expression.Name("Status").Equal(expression.Value(Status))

	// What values do we want back from our DynamoDB Query?
	proj := expression.NamesList(
		expression.Name("Status"),
		expression.Name("UserName"),
		expression.Name("UserPass"),
		expression.Name("SiteName"),
		expression.Name("UserAgent"),
		expression.Name("ProxyAddress"),
		expression.Name("AccountType"),
		expression.Name("AgentPlayers"),
	)

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	// This is the slice of Masters we will be appending to and returning from this function
	returnAllMasters := []Master{}

	// For each user from the Query API call
	for _, i := range result.Items {
		master := Master{}

		err = dynamodbattribute.UnmarshalMap(i, &master)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create Temporary User struct that we will append back outside to save this data from DynamoDB
		var tempMaster = Master{
			Status:       master.Status,
			UserName:     master.UserName,
			UserPass:     master.UserPass,
			SiteName:     master.SiteName,
			UserAgent:    master.UserAgent,
			ProxyAddress: "http://" + master.ProxyAddress,
			AccountType:  master.AccountType,
			AgentPlayers: master.AgentPlayers,
		}

		// If the Account Type is an Agent Account, we know there are users they follow
		if tempMaster.AccountType == "Agent" {
			// Split the Agent Players String after removing all spaces, and splitting by commas into a Slice
			tempMaster.AgentPlayersSlice = strings.Split(helper.ReplaceParameters(tempMaster.AgentPlayers, " ", ""), ",")
		}

		// Use tempUserStruct and append to the slice of Users Struct held outside in overallWorkflow
		returnAllMasters = append(returnAllMasters, tempMaster)
	}

	return returnAllMasters
}

// Sent API call to DynamoDB to gather All Open Bets
func GatherOpenBets() []Lines {

	// Initialize a session that the SDK will use to load
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Start creating Queries with these arguments
	tableName := "OpenBets"

	// What values do we want back from our DynamoDB Query?
	proj := expression.NamesList(
		expression.Name("TicketID"),
		expression.Name("UserName"),
		expression.Name("RotationNumber"),
		expression.Name("LineJuice"),
		expression.Name("LineType"),
		expression.Name("LineSpread"),
		expression.Name("OverUnder"),
		expression.Name("FavoredUnderdog"),
	)

	expr, err := expression.NewBuilder().WithProjection(proj).Build()
	if err != nil {
		fmt.Println("Got error building expression:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	// This is the slice of Lines we will be appending to and returning from this function
	returnPendingBets := []Lines{}

	// For each user from the Query API call
	for _, i := range result.Items {
		lines := Lines{}

		err = dynamodbattribute.UnmarshalMap(i, &lines)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		// Create Temporary User struct that we will append back outside to save this data from DynamoDB
		var tempLine = Lines{
			TicketID:       lines.TicketID,
			RotationNumber: lines.RotationNumber,
			LineJuice:      lines.LineJuice,
			LineType:       lines.LineType,
			UserName:       lines.UserName,

			// This is required by Total:
			LineSpread: lines.LineSpread,
			OverUnder:  lines.OverUnder,

			// This is required by Spread:
			FavoredUnderdog: lines.FavoredUnderdog,
		}

		// Use tempUserStruct and append to the slice of Users Struct held outside in overallWorkflow
		returnPendingBets = append(returnPendingBets, tempLine)
	}

	return returnPendingBets
}

// Sent API call to DynamoDB to gather Site Dict, will append it Slaves based on their SiteName
func GatherSportsDict() []SiteDictionary {

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Start creating Queries with these arguments
	tableName := "SitesDict"

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)
	if err != nil {
		fmt.Println("Query API call failed:")
		fmt.Println((err.Error()))
		os.Exit(1)
	}

	// This is the slice of Site Dictionaries we will be appending to and returning from this function
	var returnAllSiteDict []SiteDictionary

	// For each user from the Query API call
	for _, i := range result.Items {
		// Create singular SiteDict
		siteDict := SiteDictionary{}
		// Unmarshal into a singular SiteDict
		err = dynamodbattribute.UnmarshalMap(i, &siteDict)
		// Append to the slice of SiteDict
		returnAllSiteDict = append(returnAllSiteDict, siteDict)
	}

	// Return the slice of SiteDict
	return returnAllSiteDict
}

// This function will be used to send one new OpenBets to DynamoDB
func (master *Master) PushOpenBet(lines Lines) {

	// First format the Master and Lines, will need the telegram message to look like:

	// {Master} ({MasterPass}) #{TicketID}
	// LineType: {LineType}

	// If it is a Total, The bottom lines need to look like this
	// {Team} {OverUnder} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
	// {Sport} - {League}

	// If it is a Spread, The bottom lines need to look like this
	// {Team} {Spread} ({Juice}) [{RiskAmount}/{ToWinAmount}]
	// {Sport} - {League}

	// If the Master is an Agent account, we want to send the account we are following, and not the Agent account
	tempUserName := ""
	tempUserPass := ""

	if master.AccountType == "Agent" {
		tempUserName = lines.UserName
		tempUserPass = lines.UserPass
	} else {
		tempUserName = master.MasterName
		tempUserPass = master.MasterPass
	}

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

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	a, err := dynamodbattribute.MarshalMap(lines)
	if err != nil {
		fmt.Println("Got error marshalling new movie item:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Create item in table Movies
	tableName := "OpenBets"

	input := &dynamodb.PutItemInput{
		Item:      a,
		TableName: aws.String(tableName),
	}

	_, err = svc.PutItem(input)
	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Now that we've added it to the Database, we need to add it to the list of Open Bets on the Master so we don't
	// scrape it again
	master.MasterLines = append(master.MasterLines, lines)

	SendTelegram(telegramMsg, master.Status)

}

// This function will be used to send one new OpenBets to DynamoDB
func (master *Master) RemoveOpenBet(lines Lines) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create item in table Movies
	tableName := "OpenBets"

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"TicketID": {
				S: aws.String(lines.TicketID),
			},
			"UserName": {
				S: aws.String(lines.UserName),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		SendTelegram("Error deleting OpenBet "+lines.TicketID+", Master - "+master.MasterName, master.Status)
		fmt.Println(err.Error())
		return
	}

}

*/
