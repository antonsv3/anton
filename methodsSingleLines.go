package anton

// I NEED TO MAKE SURE THE VALUES FOR ADDMASTERLINEVALUES() ARE GOOD THROUGH DICTIONARIES

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// Properties I need to validate:
// Team
// League
// Sport

// This method lets me know what additional values I need for a Master Line to ensure consistency
func (masterLine *Lines) AddMasterLineValues(riskAmount, toWinAmount, period, league, sport, team, masterName, masterPass string) {

	// Format Line with additional values, will need a method
	masterLine.RiskAmount = riskAmount
	masterLine.ToWinAmount = toWinAmount
	masterLine.Team = team
	masterLine.MasterName = masterName
	masterLine.MasterPass = masterPass

	// These three properties I need to run through and translate, will need their own functions
	masterLine.League = league
	masterLine.Sport = sport
	masterLine.Period = period
}

// This function validates the shared properties between Slave and Master, and is used by the two functions above
func (line *Lines) ValidateSingleLine() {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Let's start by validating the values that we know are fixed in slices in the configurableParameters file

	// ------------------------------------------------------------------------------- BetType
	// All Lines need to be differentiated between being an Master or Slave Line

	// Values are: Master, Slave
	betTypeValues := GetBetTypeValues()
	if helper.FindIfStringInSlice(line.BetType, betTypeValues) == "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BetType -> Should be either: "+
			strings.Join(betTypeValues, ", "))
	}

	// ------------------------------------------------------------------------------- LineType
	// All Lines need to be differentiated between having some sort of LineType

	// Values are: MoneyLine, Total, Spread, TeamTotal
	lineTypeValues := GetLineTypeValues()

	if helper.FindIfStringInSlice(line.LineType, lineTypeValues) == "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineType -> Should be either: "+
			strings.Join(lineTypeValues, ", "))
	}

	// ------------------------------------------------------------------------------- HomeAway
	// Removed HomeAway as required field, only throw Error if it's not empty and is not either one of the Values

	// Values are: Home, Away
	homeAwayValues := GetHomeAwayValues()

	if line.HomeAway != "" {
		if helper.FindIfStringInSlice(line.HomeAway, homeAwayValues) == "False" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: HomeAway -> Should be either: "+
				strings.Join(homeAwayValues, ", "))
		}
	}

	// ------------------------------------------------------------------------------- CreatedViaFunction
	// All Lines should be created through Functions

	// Values are: True
	createdViaFunctionValues := GetCreatedViaFunctionValues()
	if helper.FindIfStringInSlice(line.CreatedViaFunction, createdViaFunctionValues) == "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: CreatedViaFunction -> Line "+
			"should be created leveraging Create Function")
	}

	// ------------------------------------------------------------------------------- FavoredUnderdog
	// This should only be applicable to LineType = Spread

	// Values are: Favored, Underdog
	favoredUnderdogValues := GetFavoredUnderdogValues()

	if line.LineType == "Spread" {
		if helper.FindIfStringInSlice(line.FavoredUnderdog, favoredUnderdogValues) == "False" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: FavoredUnderdog -> Should be "+
				"either: "+strings.Join(favoredUnderdogValues, ", "))
		}
	}

	// ------------------------------------------------------------------------------- OverUnder
	// This should only be applicable to LineType = Total

	// Values are: Over, Under
	overUnderValues := GetOverUnderValues()

	if line.LineType == "Total" {
		if helper.FindIfStringInSlice(line.OverUnder, overUnderValues) == "False" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: OverUnder -> Should be either: "+
				strings.Join(overUnderValues, ", "))
		}
	}

	// Let's now check the fields that should be required, but do not have set values, and check if they are blank

	// ------------------------------------------------------------------------------- BettingSite
	// Need to make sure that BettingSites are correctly configured and not empty

	if line.BettingSite == "" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingSite -> Please assign the Site "+
			"Name where Line is found")

		// BettingSite should also have prefix of at least "http" (will work for https) , as well as suffix of "/"
	} else if !strings.HasPrefix(line.BettingSite, "http") || !strings.HasSuffix(line.BettingSite, "/") {

		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingSite -> Incorrectly "+
			"structured, ex. http://247sports.bet/")
	}

	// ------------------------------------------------------------------------------- BettingUser
	if line.BettingUser == "" {

		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingUser -> Please assign"+
			" the BettingUser where Line is found")
	}

	// ------------------------------------------------------------------------------- RotationNumber
	if line.RotationNumber == "" {

		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: RotationNumber -> Please "+
			"assign the RotationNumber")

	}

	// Convert to Float and if it doesn't match the current Float value, there is an Error
	tempRotationFloatValue, _ := strconv.ParseFloat(line.RotationNumber, 32)

	// Rotation Numbers should always be a Positive Number, change it to a Float and see if it is greater than 0
	if tempRotationFloatValue <= 0 {

		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: RotationNumber -> Please "+
			"assign a string value of a Number")

	}

	// ------------------------------------------------------------------------------- TicketID
	if line.TicketID == "" {

		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: TicketID -> Please "+
			"assign the TicketID")

	}

	// I'm going to comment these three out until I finish getting the mapping and consistency between sites figured out

	/*
		// ------------------------------------------------------------------------------- TeamName
		if returnLine.TeamName == "" {
			returnLine.ErrorLog = append(returnLine.ErrorLog, "TeamName -> Please assign the Team Name")
		}

		// ------------------------------------------------------------------------------- LeagueName
		if returnLine.LeagueName == "" {
			returnLine.ErrorLog = append(returnLine.ErrorLog, "LeagueName -> Please assign the LeagueName")
		}

		// ------------------------------------------------------------------------------- LeagueID
		if returnLine.LeagueID == "" {
			returnLine.ErrorLog = append(returnLine.ErrorLog, "LeagueID -> Please assign the LeagueID")
		}
	*/

	// Let's now do specific requirement values for specific LineTypes: MoneyLine vs Spread vs Total

	// ------------------------------------------------------------------------------- LineSpread

	// Let's check if LineSpread is even populated, if not, add failed check
	if line.LineSpread == "" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpread -> Please assign the Spread "+
			"Value")
	}

	// Let's first make sure MoneyLine has a LineSpread of some EVEN values
	lineSpreadEvenValues := GetSpreadEvenValues()

	// Let's see if the LineSpread is an EVEN value, this will be "False" if the Spread is not an EVEN Value
	lineSpreadEvenFlag := helper.FindIfStringInSlice(strings.ToUpper(line.LineSpread), lineSpreadEvenValues)

	if line.LineType == "MoneyLine" && lineSpreadEvenFlag == "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpread -> LineType is MoneyLine but "+
			"LineSpread is not an EVEN value")
	}

	// ------------------------------------------------------------------------------- LineSpreadFloat

	// Let's check if LineSpread is even populated, if not, add failed check
	if line.LineSpreadFloat == 0 && lineSpreadEvenFlag == "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> LineSpread is not an "+
			"EVEN value but LineSpreadFloat is 0")
	}

	// We know it errors if not populated, we can start Validating the values and see if LineSpreadFloat is okay
	// Lets start first by checking "0" or EVEN values, this checks if it is a value in the slice lineSpreadEvenValues
	if line.LineSpreadFloat != 0 && lineSpreadEvenFlag != "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> LineSpread is an EVEN"+
			" value but LineSpreadFloat is not 0")
	}

	// Convert to Float and if it doesn't match the current Float value, there is an Error
	tempSpreadFloatValue, _ := strconv.ParseFloat(line.LineSpread, 32)

	// The exception is that Total Lines will have their spread always positive, we'll check this next
	if line.LineType != "Total" {
		if tempSpreadFloatValue != line.LineSpreadFloat && tempSpreadFloatValue != (line.LineSpreadFloat*-1) {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> LineSpread did "+
				"not get properly parsed into a Float Value")
		}
	}

	// Let's check for consistency on LineSpreadFloats, All Total Spread Floats should be positive
	if line.LineType == "Total" && line.LineSpreadFloat < 0 {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> Need Consistency, "+
			"Value is Negative, LineType = Total, Float should always be positive")
	}

	// -------------------- Spread - Spread Floats

	if line.LineType == "Spread" {

		// All Spread Favored should be negatives, if it is positive, multiply it by (-1)
		if line.LineSpreadFloat > 0 && line.FavoredUnderdog == "Favored" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> Need Consistency,"+
				" Value is Positive, LineType = Spread & Favored, Float should be Negative")
		}

		// All Spread Underdog should be positive, if it is negative, multiply it by (-1)
		if line.LineSpreadFloat < 0 && line.FavoredUnderdog == "Underdog" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> Need Consistency,"+
				" Value is Negative, LineType = Spread & Underdog, Float should be Positive")
		}
	}

	// ------------------------------------------------------------------------------- LineJuice

	// Let's check if LineJuice is even populated, if not, add failed check
	if line.LineJuice == "" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineJuice -> Please assign Juice Value")
	}

	// Let's first make sure MoneyLine has a LineJuice of some EVEN values
	lineJuiceEvenValues := GetJuiceEvenValues()

	// Let's see if the LineJuice is an EVEN value, this will be "False" if the Spread is not an EVEN Value
	lineJuiceEvenFlag := helper.FindIfStringInSlice(strings.ToUpper(line.LineJuice), lineJuiceEvenValues)

	// ------------------------------------------------------------------------------- LineJuiceFloat

	// We know if the LineJuice is not in the EVEN slice, the Flag will not be "False", check that vs if Float is 100
	if lineJuiceEvenFlag != "False" && line.LineJuiceFloat != 100 {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineJuiceFloat -> LineJuice is an EVEN "+
			"value but LineJuiceFloat is not 100")
	}

	// Let's check if LineJuiceFloat is even populated, default value is 0 and if it is not an EVEN Value, Error
	if lineJuiceEvenFlag == "False" && line.LineJuiceFloat == 100 {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineJuiceFloat -> LineJuice is not an "+
			"EVEN value but LineJuiceFloat is 100")
	}

	// Convert to Float and if it doesn't match the current Float value, there is an Error
	tempJuiceFloatValue, _ := strconv.ParseFloat(line.LineJuice, 32)

	if tempJuiceFloatValue != line.LineJuiceFloat && tempJuiceFloatValue != (line.LineJuiceFloat*-1) {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineJuiceFloat -> LineJuice did not get "+
			"properly parsed into a Float Value")
	}

	// I also know that FloatValues shouldn't be between -99 and 99 so I will check for that
	if line.LineJuiceFloat <= 99 && line.LineJuiceFloat >= -99 {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: Line Juice Float Value is between +99 "+
			"and -99")
	}

	// Now I can start ending this function

	// ------------------------------------------------------------------------------- LineStatus
	// Values are: New, Validated, Authorized, Placed, Error
	lineStatusValues := GetLineStatusValues()
	if line.LineStatus == "" {
		line.LineStatus = "Error"
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineStatus -> Currently Blank, Please assign one"+
			" of the values: "+strings.Join(lineStatusValues, ", "))
	} else if helper.FindIfStringInSlice(line.LineStatus, lineStatusValues) == "False" {
		line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineStatus -> Should be either: "+
			strings.Join(lineStatusValues, ", "))
	}

	// Let's now replace all Error log's {betType} and {lineType}
	for i := range line.ErrorLog {
		line.ErrorLog[i] = helper.ReplaceParameters(line.ErrorLog[i], "{betType}", line.BetType, "{lineType}",
			line.LineType)
	}

	// ------------------------------------------------------------------------------- LineLog
	// Slave Line Logs
	// If there are failed checks from this function and Log is empty, log it
	if line.BetType == "Slave" && len(line.ErrorLog) != 0 {
		line.LineStatus = "Error"
		line.FunctionLog = "[#validateSingleLine] New Slave Line has not yet passed Validation"
		// If there are no failed checks and Log is empty, Log it and Change Status to Validated
	} else if line.BetType == "Slave" && len(line.ErrorLog) == 0 {
		line.LineStatus = "Validated"
		line.FunctionLog = "[#validateLine] New Slave Line passed Validation"
	}

	// Master Line Logs
	// If there are failed checks from this function and Log is empty, log it
	if line.BetType == "Master" && len(line.ErrorLog) != 0 {
		line.LineStatus = "Error"
		line.FunctionLog = "[#validateLine] New Master Line has not yet passed Validation"
		// If there are no failed checks and Log is empty, Log it and Change Status to Validated
	} else if line.BetType == "Master" && len(line.ErrorLog) == 0 {
		line.LineStatus = "Validated"
		line.FunctionLog = "[#validateLine] New Master Line passed Validation"
	}

}

// This function will print lines properties based from the configurable parameters
func (line Lines) PrintSingleLine() {

	// Only print the Line if the Line Status is not "Ignored"
	if line.LineStatus != "Ignored" {

		// Formatting the Print Header
		tempHeader := centerString("Print Single Line", 20)
		tempHeader = "< ---------------------------------------" + tempHeader +
			"                                         "

		fmt.Println("")
		fmt.Println(centerString(tempHeader, 127))
		fmt.Println("")

		propertiesToPrint := GetPrintProperties()

		// Hold the field and values of the Slave Line Struct and Master Line Struct
		slaveProperties := reflect.TypeOf(line)
		slaveValues := reflect.ValueOf(line)

		// Hold the amount of fields of the Line Struct to loop over
		num := slaveProperties.NumField()

		// Loop over all the properties in the slice to print
		for i := 0; i < len(propertiesToPrint); i++ {

			// We know the first property in the slice is propertiesToPrint[i], now lets iterate through all properties
			for j := 0; j < num; j++ {

				// The current, single Slave Line Property and Value in this loop
				slaveProperty := slaveProperties.Field(j)
				slaveValue := slaveValues.Field(j)

				// Now we know we have the right property and in the correct order, lets format our print inside this If
				if slaveProperty.Name == propertiesToPrint[i] && slaveProperty.Name != "FunctionLog" &&
					slaveProperty.Name != "ErrorLog" {

					centerPrint := "     <------------" + centerString(slaveProperty.Name, 20)

					// Only Print FavoredUnderdog if LineType == "Spread", and OverUnder if LineType == "Total"
					if slaveProperty.Name != "FavoredUnderdog" && slaveProperty.Name != "OverUnder" {
						fmt.Printf("     %-30v%v\n", slaveValue, centerPrint)
					} else if slaveProperty.Name == "FavoredUnderdog" && line.LineType == "Spread" {
						fmt.Printf("     %-30v%v\n", slaveValue, centerPrint)
					} else if slaveProperty.Name == "OverUnder" && line.LineType == "Total" {
						fmt.Printf("     %-30v%v\n", slaveValue, centerPrint)
					}
				}

			}
		}

		fmt.Printf("\n     Function Log: %v\n\n", slaveValues.FieldByName("FunctionLog"))
		for i := 0; i < slaveValues.FieldByName("ErrorLog").Len(); i++ {
			fmt.Printf("     Error Log #%v: %v\n", i+1, slaveValues.FieldByName("ErrorLog").Index(i))
		}
		fmt.Println("")
	}

}

// Helper function for above to help center the print
func centerString(s string, w int) string {
	return fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
}
