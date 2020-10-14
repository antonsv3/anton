package anton

import (
	"strconv"
	"strings"
	"time"
)

/*

Create Slave Lines Function Breakdown:

- Slave Method

	- CreateSlaveMoneyLine
		- formatSlaveLineValues
			- formatPeriod
			- formatSport
			- formatLeague

	- CreateSlaveSpread
		- formatSlaveLineValues
			- formatPeriod
			- formatSport
			- formatLeague

	- CreateSlaveTotal
		- formatSlaveLineValues
			- formatPeriod
			- formatSport
			- formatLeague

	- CreateSlaveTeamTotal
		- formatSlaveLineValues
			- formatPeriod
			- formatSport
			- formatLeague

*/

func (slave Slave) CreateSlaveMoneyLine(rotationNumber, period, lineJuice, sport, league string) Lines {

	// Create the New Line Struct that will be returned by this function
	returnSlaveMoneyLine := Lines{
		LineType: "MoneyLine",
	}

	// Now that we've got MoneyLine Specific values, we can call the formatSlaveLineValues for the shared values
	returnSlaveMoneyLine = formatSlaveLineValues(slave, returnSlaveMoneyLine, rotationNumber, period, "0",
		lineJuice, sport, league)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(returnSlaveMoneyLine.ErrorLog) > 0 {
		returnSlaveMoneyLine.LineStatus = "Error"
		returnSlaveMoneyLine.FunctionLog = "[#createSlaveMoneyLine] Unable to Create New Slave MoneyLine"
	}

	return returnSlaveMoneyLine
}

func (slave Slave) CreateSlaveSpread(rotationNumber, period, lineSpread, lineJuice, favoredUnderdog,
	sport, league string) Lines {

	// Create New Line Struct that will be returned by this function
	returnSlaveSpread := Lines{
		LineType: "Spread",
	}

	// Lets populate OverUnder as it is specific to Total
	returnSlaveSpread.FavoredUnderdog = favoredUnderdog

	// Now that we've got Spread Specific values, we can call the formatSlaveLineValues for the shared values
	returnSlaveSpread = formatSlaveLineValues(slave, returnSlaveSpread, rotationNumber, period, lineSpread, lineJuice,
		sport, league)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(returnSlaveSpread.ErrorLog) > 0 {
		returnSlaveSpread.LineStatus = "Error"
		returnSlaveSpread.FunctionLog = "[#createSlaveSpread] Unable to Create New Slave Spread"
	}

	return returnSlaveSpread
}

func (slave Slave) CreateSlaveTotal(rotationNumber, period, lineSpread, lineJuice, overUnder, sport, league string) Lines {

	// Create New Line Struct that will be returned by this function
	returnSlaveTotal := Lines{
		LineType: "Total",
	}

	// Lets populate OverUnder as it is specific to Total
	returnSlaveTotal.OverUnder = overUnder

	// Now that we've got Total Specific values, we can call the formatSlaveLineValues for the shared values
	returnSlaveTotal = formatSlaveLineValues(slave, returnSlaveTotal, rotationNumber, period, lineSpread, lineJuice,
		sport, league)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(returnSlaveTotal.ErrorLog) > 0 {
		returnSlaveTotal.LineStatus = "Error"
		returnSlaveTotal.FunctionLog = "[#createSlaveTotal] Unable to Create New Slave Total Line"
	}

	return returnSlaveTotal
}

func (slave Slave) CreateSlaveTeamTotal(rotationNumber, period, lineSpread, lineJuice, overUnder,
	sport, league string) Lines {

	// Create New Line Struct that will be returned by this function
	returnSlaveTeamTotal := Lines{
		LineType: "TeamTotal",
	}

	// Lets populate OverUnder as it is specific to Total
	returnSlaveTeamTotal.OverUnder = overUnder

	// Now that we've got Total Specific values, we can call the formatSlaveLineValues for the shared values
	returnSlaveTeamTotal = formatSlaveLineValues(slave, returnSlaveTeamTotal, rotationNumber, period,
		lineSpread, lineJuice, sport, league)

	// If there is any Errors, set the Status to "Error" and Log the Function
	if len(returnSlaveTeamTotal.ErrorLog) > 0 {
		returnSlaveTeamTotal.LineStatus = "Error"
		returnSlaveTeamTotal.FunctionLog = "[#createSlaveTotal] Unable to Create New Slave Total Line"
	}

	return returnSlaveTeamTotal
}

func formatSlaveLineValues(slave Slave, slaveLine Lines, rotationNumber, period, lineSpread, lineJuice,
	sport, league string) Lines {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Create a Line that takes the tempLine values, this function will return this line when it is done
	returnSlaveLine := slaveLine

	returnSlaveLine.CreatedTimestamp = time.Now()

	// Append Static Values that'll be changed if there are any errors with this function
	returnSlaveLine.CreatedViaFunction = "True"
	returnSlaveLine.LineStatus = "New"
	returnSlaveLine.FunctionLog = "[#createSlave{LineType}] New Slave {LineType} created through function"
	returnSlaveLine.FunctionLog = helper.ReplaceParameters(returnSlaveLine.FunctionLog, "{LineType}",
		returnSlaveLine.LineType)

	// Append the parameters from the calling of one of those three functions, which we passed into here
	returnSlaveLine.RotationNumber = rotationNumber
	returnSlaveLine.LineJuice = helper.ReplaceParameters(lineJuice, "½", ".5", " ", "")
	returnSlaveLine.LineSpread = helper.ReplaceParameters(lineSpread, "½", ".5", " ", "")

	// I want to add "+" in front of the LineSpread, if it is Positive and only if it's not Total or TeamTotal
	// This will remove "+" and "-" if it is a Total or TeamTotal, which we'll be using OverUnder to compare
	if returnSlaveLine.LineType != "Total" && returnSlaveLine.LineType != "TeamTotal" {
		if helper.StringNegativePositiveZero(returnSlaveLine.LineSpread) == "Positive" ||
			helper.StringNegativePositiveZero(returnSlaveLine.LineSpread) == "Even" {

			if !strings.HasPrefix(returnSlaveLine.LineSpread, "+") {
				returnSlaveLine.LineSpread = "+" + returnSlaveLine.LineSpread
			}
		}
	} else {
		if strings.HasPrefix(returnSlaveLine.LineSpread, "+") || strings.HasPrefix(returnSlaveLine.LineSpread, "-") {
			returnSlaveLine.LineSpread = helper.ReplaceParameters(returnSlaveLine.LineSpread, "+", "", "-", "")
		}
	}

	// I want to add "+" in front of the LineJuice, if it is Positive
	if helper.StringNegativePositiveZero(returnSlaveLine.LineJuice) == "Positive" ||
		helper.StringNegativePositiveZero(returnSlaveLine.LineSpread) == "Even" {

		if !strings.HasPrefix(returnSlaveLine.LineJuice, "+") {
			returnSlaveLine.LineJuice = "+" + returnSlaveLine.LineJuice
		}
	}

	// Append Inherited Values from the Current Slave
	returnSlaveLine.BetType = "Slave"
	returnSlaveLine.SlaveName = slave.SlaveName
	returnSlaveLine.SlavePass = slave.SlavePass
	returnSlaveLine.SlaveSite = slave.SiteName

	// Format the Sport, League and Period
	returnSlaveLine.Sport = helper.FormatSport(sport)
	returnSlaveLine.League = helper.FormatLeague(league)
	returnSlaveLine.Period = helper.FormatPeriod(period)

	// ----------------------------------- Converting LineSpread to LineSpreadFloat --------------------------------- //

	// Assign the Slice of Even Spread Values to new variable to help compare
	spreadEvenValues := GetSpreadEvenValues()

	// We know that the spread for MoneyLine is going to be Zero, so lets just start with populating that
	if returnSlaveLine.LineType == "MoneyLine" {
		returnSlaveLine.LineSpread = "0"
	}

	// Because GoLang uses 0 if there is an error when converting to float/integers, we need to see if error or not
	var spreadErrorFlag string

	// Use our helper function that sees if the Spread in the EVEN slice, otherwise will be a string value of "False"
	if helper.FindIfStringInSlice(strings.ToUpper(returnSlaveLine.LineSpread), spreadEvenValues) != "False" {
		returnSlaveLine.LineSpreadFloat = 0
		spreadErrorFlag = "False"
	} else {
		returnSlaveLine.LineSpreadFloat, _ = strconv.ParseFloat(returnSlaveLine.LineSpread, 32)
		spreadErrorFlag = "True"
	}

	// If the LineSpreadFloat is 0 and the Flag is "True", then an error happened when converting, log error if true
	if returnSlaveLine.LineSpreadFloat == 0 && spreadErrorFlag == "True" {
		returnSlaveLine.ErrorLog = append(returnSlaveLine.ErrorLog, "Could Not Parse LineSpread to Float Value")
		spreadErrorFlag = "True"
	} else {
		spreadErrorFlag = "False"
	}

	// ----------------------------------------- LineSpreadFloat Consistencies -------------------------------------- //

	// -------------------- MoneyLine - Spread Floats

	// Remember, MoneyLine does not have any Spread, it will be always be 0
	if returnSlaveLine.LineType == "MoneyLine" {
		returnSlaveLine.LineSpread = "0"
		returnSlaveLine.LineSpreadFloat = 0
	}

	// -------------------- Total & TeamTotal - Spread Floats

	// All Total Spread Floats will be positive, Over 50 points, Under 50 points, we will use Over/Under for comparisons
	if returnSlaveLine.LineType == "Total" || returnSlaveLine.LineType == "TeamTotal" {
		if returnSlaveLine.LineSpreadFloat < 0 && spreadErrorFlag == "False" {
			returnSlaveLine.LineSpreadFloat = returnSlaveLine.LineSpreadFloat * (-1)
		}

	}

	// -------------------- Spread - Spread Floats

	if returnSlaveLine.LineType == "Spread" && spreadErrorFlag == "False" {

		// All Spread Favored should be negatives, if it is positive, multiply it by (-1)
		if returnSlaveLine.LineSpreadFloat > 0 && returnSlaveLine.FavoredUnderdog == "Favored" {
			returnSlaveLine.LineSpreadFloat = returnSlaveLine.LineSpreadFloat * (-1)
		}

		// All Spread Underdog should be positive, if it is negative, multiply it by (-1)
		if returnSlaveLine.LineSpreadFloat < 0 && returnSlaveLine.FavoredUnderdog == "Underdog" {
			returnSlaveLine.LineSpreadFloat = returnSlaveLine.LineSpreadFloat * (-1)
		}

	}

	// ------------------------------------ Converting LineJuice to LineJuiceFloat ---------------------------------- //

	// Assign the Slice of Even Spread Values to new variable to help compare
	juiceEvenValues := GetJuiceEvenValues()

	// Use our helper function that sees if the Juice in the EVEN slice, otherwise will be a string value of "False"
	// Remember we are setting the Float Value to Positive 100, we will need to test using -100 as well later on
	if helper.FindIfStringInSlice(strings.ToUpper(returnSlaveLine.LineJuice), juiceEvenValues) != "False" {
		returnSlaveLine.LineJuiceFloat = 100
	} else {
		returnSlaveLine.LineJuiceFloat, _ = strconv.ParseFloat(returnSlaveLine.LineJuice, 32)
	}

	// If the LineJuiceFloat is equal to 0, or that means that an error happened when converting, log error if true
	if returnSlaveLine.LineJuiceFloat == 0 {
		returnSlaveLine.ErrorLog = append(returnSlaveLine.ErrorLog, "Could Not Parse LineJuice to Float Value")

		// I also know Juice Values shouldn't be between 99 and -99, log error if true
	} else if returnSlaveLine.LineJuiceFloat > -99 && returnSlaveLine.LineJuiceFloat < 99 {
		returnSlaveLine.ErrorLog = append(returnSlaveLine.ErrorLog, "LineJuiceFloat is between -99 and 99")
	}

	// ------------------------------------------ LineJuiceFloat Consistencies -------------------------------------- //

	// There are no LineJuiceFloat Consistencies, Juice isn't affected whether FavoredUnderdog or OverUnder

	// -------------------- Capture Error if they don't fit into any of these functions, which is sorted by LineType

	if returnSlaveLine.LineType != "Total" && returnSlaveLine.LineType != "TeamTotal" {
		if returnSlaveLine.LineType != "MoneyLine" && returnSlaveLine.LineType != "Spread" {
			returnSlaveLine.ErrorLog = append(returnSlaveLine.ErrorLog, "LineType is not MoneyLine, Spread, Total, or TeamTotal")
		}
	}

	return returnSlaveLine
}
