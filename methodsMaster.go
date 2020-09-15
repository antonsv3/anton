package anton

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// The Different Methods in the Anton Package for Masters are listed in the order below:
// - CreateNewMaster()				(<-- Although not a Method, it returns an empty Master)
// - SendToAnton()
// -

func (master Master) CreateMasterMoneyLine(ticketID, rotationNumber, lineJuice string) Lines {

	// Create the New Line Struct that will be returned by this function
	newMasterMoneyLine := Lines{
		LineType: "MoneyLine",
		TicketID: ticketID,
	}

	// Now that we've got MoneyLine Specific values, we can call the formatMasterValues for the shared values
	newMasterMoneyLine = formatMasterValues(master, newMasterMoneyLine, rotationNumber, "0", lineJuice)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(newMasterMoneyLine.ErrorLog) > 0 {
		newMasterMoneyLine.LineStatus = "Error"
		newMasterMoneyLine.FunctionLog = "[#createMasterMoneyLine] Unable to Create New Master MoneyLine"
	}

	return newMasterMoneyLine
}

func (master Master) CreateMasterSpread(ticketID, rotationNumber, lineSpread, lineJuice,
	favoredUnderdog string) Lines {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Create New Line Struct that will be returned by this function
	newMasterSpread := Lines{
		LineType: "Spread",
		TicketID: ticketID,
	}

	// Make sure the value of the FavoredUnderdog are in within the Parameters
	favoredUnderdogValues := GetFavoredUnderdogValues()
	if helper.FindIfStringInSlice(favoredUnderdog, favoredUnderdogValues) != "False" {
		newMasterSpread.FavoredUnderdog = favoredUnderdog
	} else {
		newMasterSpread.ErrorLog = append(newMasterSpread.ErrorLog, "Master Spread: FavoredUnderdog -> Should be either: "+
			strings.Join(favoredUnderdogValues, ", "))
	}

	// Now that we've got Spread Specific values, we can call the formatMasterValues for the shared values
	newMasterSpread = formatMasterValues(master, newMasterSpread, rotationNumber, lineSpread, lineJuice)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(newMasterSpread.ErrorLog) > 0 {
		newMasterSpread.LineStatus = "Error"
		newMasterSpread.FunctionLog = "[#createMasterSpread] Unable to Create New Master Line"
	}

	return newMasterSpread
}

func (master Master) CreateMasterTotal(ticketID, rotationNumber, lineSpread, lineJuice,
	overUnder string) Lines {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Create New Line Struct that will be returned by this function
	newMasterTotal := Lines{
		LineType: "Total",
		TicketID: ticketID,
	}

	// Make sure the value of the OverUnder are in within the Parameters
	overUnderValues := GetOverUnderValues()
	if helper.FindIfStringInSlice(overUnder, overUnderValues) != "False" {
		newMasterTotal.OverUnder = overUnder
	} else {
		newMasterTotal.ErrorLog = append(newMasterTotal.ErrorLog, "Master Total: FavoredUnderdog -> Should be either: "+
			strings.Join(overUnderValues, ", "))
	}

	// Now that we've got Total Specific values, we can call the formatMasterValues for the shared values
	newMasterTotal = formatMasterValues(master, newMasterTotal, rotationNumber, lineSpread, lineJuice)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(newMasterTotal.ErrorLog) > 0 {
		newMasterTotal.LineStatus = "Error"
		newMasterTotal.FunctionLog = "[#createMasterTotal] Unable to Create New Master Total Line"
	}

	return newMasterTotal
}

func (master Master) CreateMasterTeamTotal(ticketID, rotationNumber, lineSpread, lineJuice,
	overUnder string) Lines {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Create New Line Struct that will be returned by this function
	newMasterTotal := Lines{
		LineType: "TeamTotal",
		TicketID: ticketID,
	}

	// Make sure the value of the OverUnder are in within the Parameters
	overUnderValues := GetOverUnderValues()
	if helper.FindIfStringInSlice(overUnder, overUnderValues) != "False" {
		newMasterTotal.OverUnder = overUnder
	} else {
		newMasterTotal.ErrorLog = append(newMasterTotal.ErrorLog, "Master Total: FavoredUnderdog -> Should be either: "+
			strings.Join(overUnderValues, ", "))
	}

	// Now that we've got Total Specific values, we can call the formatMasterValues for the shared values
	newMasterTotal = formatMasterValues(master, newMasterTotal, rotationNumber, lineSpread, lineJuice)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(newMasterTotal.ErrorLog) > 0 {
		newMasterTotal.LineStatus = "Error"
		newMasterTotal.FunctionLog = "[#createMasterTotal] Unable to create New Master Total Line"
	}

	return newMasterTotal
}

func formatMasterValues(master Master, tempLine Lines, rotationNumber, lineSpread, lineJuice string) Lines {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Create a Line that takes the tempLine values, this function will return this line when it is done
	returnMasterLine := tempLine

	// Append Static Values that'll be changed if there are any errors with this function
	returnMasterLine.CreatedViaFunction = "True"
	returnMasterLine.LineStatus = "New"
	returnMasterLine.FunctionLog = "[#createMasterLine] New Master Line created through function"

	// Append the parameters from the calling of one of those three functions, which we passed into here
	returnMasterLine.RotationNumber = helper.ReplaceParameters(rotationNumber, "\u00a0", "", " ", "")
	returnMasterLine.LineJuice = helper.ReplaceParameters(lineJuice, "½", ".5", "\u00a0", "")
	returnMasterLine.LineSpread = helper.ReplaceParameters(lineSpread, "½", ".5", "\u00a0", "")

	// I want to add "+" in front of the LineSpread, if it is Positive and only if it's not Total or TeamTotal
	if returnMasterLine.LineType != "Total" && returnMasterLine.LineType != "TeamTotal" {
		if helper.StringNegativePositiveZero(returnMasterLine.LineSpread) == "Positive" {
			if !strings.HasPrefix(returnMasterLine.LineSpread, "+") {
				returnMasterLine.LineSpread = "+" + returnMasterLine.LineSpread
			}
		}
	} else {
		if strings.HasPrefix(returnMasterLine.LineSpread, "+") || strings.HasPrefix(returnMasterLine.LineSpread, "-") {
			returnMasterLine.LineSpread = helper.ReplaceParameters(returnMasterLine.LineSpread, "+", "", "-", "")
		}
	}

	// I want to add "+" in front of the LineJuice, if it is Positive and only if it's not Total or TeamTotal
	if helper.StringNegativePositiveZero(returnMasterLine.LineJuice) == "Positive" {
		if !strings.HasPrefix(returnMasterLine.LineJuice, "+") {
			returnMasterLine.LineJuice = "+" + returnMasterLine.LineJuice
		}
	}

	// Append Inherited Values from the Current User
	returnMasterLine.BetType = "Master"
	returnMasterLine.MasterSite = master.SiteName

	// ----------------------------------- Converting LineSpread to LineSpreadFloat --------------------------------- //

	// Assign the Slice of Even Spread Values to new variable to help compare
	spreadEvenValues := GetSpreadEvenValues()

	// We know that the spread for MoneyLine is going to be Zero, so lets just start with populating that
	if returnMasterLine.LineType == "MoneyLine" {
		returnMasterLine.LineSpread = "0"
	}

	// Because GoLang uses 0 if there is an error when converting to float/integers, we need to see if error or not
	var spreadErrorFlag string

	// Use our helper function that sees if the Spread in the EVEN slice, otherwise will be a string value of "False"
	if helper.FindIfStringInSlice(strings.ToUpper(returnMasterLine.LineSpread), spreadEvenValues) != "False" {
		returnMasterLine.LineSpreadFloat = 0
		spreadErrorFlag = "False"
	} else {
		returnMasterLine.LineSpreadFloat, _ = strconv.ParseFloat(returnMasterLine.LineSpread, 32)
		spreadErrorFlag = "True"
	}

	// If the LineSpreadFloat is 0 and the Flag is "True", then an error happened when converting, log error if true
	if returnMasterLine.LineSpreadFloat == 0 && spreadErrorFlag == "True" {
		returnMasterLine.ErrorLog = append(returnMasterLine.ErrorLog, "Could Not Parse LineSpread to Float Value")
		spreadErrorFlag = "True"
	} else {
		spreadErrorFlag = "False"
	}

	// ----------------------------------------- LineSpreadFloat Consistencies -------------------------------------- //

	// -------------------- MoneyLine - Spread Floats

	// Remember, MoneyLine does not have any Spread, it will be always be 0
	if returnMasterLine.LineType == "MoneyLine" {
		returnMasterLine.LineSpread = "0"
		returnMasterLine.LineSpreadFloat = 0
	}

	// -------------------- Total & Team Total - Spread Floats

	// All Total Spread Floats will be positive, Over 50 points, Under 50 points, we will use Over/Under for comparisons
	if returnMasterLine.LineType == "Total" && returnMasterLine.LineSpreadFloat < 0 && spreadErrorFlag == "False" {
		returnMasterLine.LineSpreadFloat = returnMasterLine.LineSpreadFloat * (-1)
	}

	// -------------------- Spread - Spread Floats

	if returnMasterLine.LineType == "Spread" && spreadErrorFlag == "False" {

		// All Spread Favored should be negatives, if it is positive, multiply it by (-1)
		if returnMasterLine.LineSpreadFloat > 0 && returnMasterLine.FavoredUnderdog == "Favored" {
			returnMasterLine.LineSpreadFloat = returnMasterLine.LineSpreadFloat * (-1)
		}

		// All Spread Underdog should be positive, if it is negative, multiply it by (-1)
		if returnMasterLine.LineSpreadFloat < 0 && returnMasterLine.FavoredUnderdog == "Underdog" {
			returnMasterLine.LineSpreadFloat = returnMasterLine.LineSpreadFloat * (-1)
		}

	}

	// ------------------------------------ Converting LineJuice to LineJuiceFloat ---------------------------------- //

	// Assign the Slice of Even Spread Values to new variable to help compare
	juiceEvenValues := GetJuiceEvenValues()

	// Use our helper function that sees if the Juice in the EVEN slice, otherwise will be a string value of "False"
	// Remember we are setting the Float Value to Positive 100, we will need to test using -100 as well later on
	if helper.FindIfStringInSlice(strings.ToUpper(returnMasterLine.LineJuice), juiceEvenValues) != "False" {
		returnMasterLine.LineJuiceFloat = 100
	} else {
		returnMasterLine.LineJuiceFloat, _ = strconv.ParseFloat(returnMasterLine.LineJuice, 32)
	}

	// If the LineJuiceFloat is equal to 0, or that means that an error happened when converting, log error if true
	if returnMasterLine.LineJuiceFloat == 0 {
		returnMasterLine.ErrorLog = append(returnMasterLine.ErrorLog, "Could Not Parse LineJuice to Float Value")

		// I also know Juice Values shouldn't be between 99 and -99, log error if true
	} else if returnMasterLine.LineJuiceFloat > -99 && returnMasterLine.LineJuiceFloat < 99 {
		returnMasterLine.ErrorLog = append(returnMasterLine.ErrorLog, "LineJuiceFloat is between -99 and 99")
	}

	if returnMasterLine.LineType == "MoneyLine" {
		if returnMasterLine.LineJuiceFloat < -99 {
			returnMasterLine.FavoredUnderdog = "Underdog"
		} else if returnMasterLine.LineJuiceFloat > 99 {
			returnMasterLine.FavoredUnderdog = "Favored"
		}
	}

	// ---------------------------------------------- LineCharacteristic -------------------------------------------- //

	// Grab the Characteristic, which is either FavoredUnderdog or OverUnder
	if tempLine.FavoredUnderdog != "" {
		returnMasterLine.LineCharacteristic = tempLine.FavoredUnderdog
	} else if tempLine.OverUnder != "" {
		returnMasterLine.LineCharacteristic = tempLine.OverUnder
	} else {
		returnMasterLine.ErrorLog = append(returnMasterLine.ErrorLog, "Could not append Line Characteristic")
	}

	// ------------------------------------------ LineJuiceFloat Consistencies -------------------------------------- //

	// There are no LineJuiceFloat Consistencies, Juice isn't affected whether FavoredUnderdog or OverUnder

	// -------------------- Capture Error if they don't fit into any of these functions

	if returnMasterLine.LineType != "MoneyLine" && returnMasterLine.LineType != "Spread" {
		if returnMasterLine.LineType != "Total" && returnMasterLine.LineType != "TeamTotal" {
			returnMasterLine.ErrorLog = append(returnMasterLine.ErrorLog, "LineType is not MoneyLine, Spread, "+
				"Total, or TeamTotal")
		}
	}

	return returnMasterLine
}

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
