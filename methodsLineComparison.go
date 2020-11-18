package anton

// SLAVE CONSOLIDATION CHECKED

import (
	"fmt"
	"strconv"
	"strings"
)

/*

Compare Lines Function Breakdown:

- Line Method

	- CompareSlaveLineToMasterLine

		- ValidateSingleLine
			- Master
			- Slave

		- ValidateAgainstProfile
		- ValidateAgainstLine
		- CompareJuiceValues
		- CompareSpreadLine
		- CompareTotalLine
		- CompareTeamTotalLine

*/

func (slaveLine *Lines) CompareSlaveLineToMasterLine(masterLine Lines, slave Slave, profile Profile) {

	// This flag is to ensure that we have met all other criteria prior to comparing, default "False"
	preChecksValidFlag := "False"

	var helper Helper

	// I want to add the Team from MasterLine to SlaveLine
	slaveLine.Team = masterLine.Team

	// Now lets start these checks, Let's validate both Lines to see if values are populated correctly
	slaveLine.ValidateSingleLine()
	masterLine.ValidateSingleLine()

	// Let's start by getting our parameters from the Profile
	var juiceParameter float64
	var spreadParameter float64

	// If they both are Validated, that means they are populated correctly, so we can compare them now
	if slaveLine.LineStatus == "Validated" && masterLine.LineStatus == "Validated" {

		// Next, let's compare it to the profiles to see whether the Slave is following the master on these lines
		if len(slave.Profiles) >= 1 {

			// This variable is our parameter for what is the maximum difference of Juice between Slave and Master
			// Because GoLang uses 0 if there is an error when converting to float/integers, we need to see if error or not
			juiceParameter, _ = strconv.ParseFloat(slave.Profiles[0].JuiceParameter, 32)
			if juiceParameter == 0 && profile.JuiceParameter != "0" && profile.JuiceParameter != "0.0" {
				slaveLine.LineStatus = "Error"
				slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Could Not Parse Juice Parameter to Float Value")
			}

			// This variable is our parameter for what is the maximum difference of Spread values between Slave and Master
			// Because GoLang uses 0 if there is an error when converting to float/integers, we need to see if error or not
			spreadParameter, _ = strconv.ParseFloat(slave.Profiles[0].SpreadParameter, 32)
			if spreadParameter == 0 && profile.SpreadParameter != "0" && profile.SpreadParameter != "0.0" {
				slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Could Not Parse Spread Parameter to Float Value")
			}

			slaveLine.ValidateAgainstProfile(slave.Profiles[0])

		} else {
			slaveLine.LineStatus = "Error"
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave does not have a Profile Attached")
		}
	}

	// If slaveLine is still Validated, then we can pre-check versus the Master Line
	if slaveLine.LineStatus == "Validated" && masterLine.LineStatus == "Validated" {
		slaveLine.ValidateAgainst(masterLine)
	}

	// If slaveLine is still Validated after single line validation, profile validation, pre-check validation, flip flag
	if slaveLine.LineStatus == "Validated" && masterLine.LineStatus == "Validated" {
		preChecksValidFlag = "True"
	}

	// This will print if it failed the Single Validation or the Pre Check Comparisons
	if slaveLine.LineStatus == "Error" || masterLine.LineStatus == "Error" {
		if len(masterLine.ErrorLog) >= 1 {
			for i := range masterLine.ErrorLog {
				slaveLine.ErrorLog = append(slaveLine.ErrorLog, "MasterLine Error: "+masterLine.ErrorLog[i])
			}
		}
		slaveLine.LineStatus = "Error"
		slaveLine.FunctionLog = "[#CompareSlaveLineToMasterLine] Comparing Slave Line to an Error'd Master Line"
		slaveLine.PrintComparedLines(masterLine)
	}

	// Now that we've checked, we can compare the two Lines now by calling the functions below
	if preChecksValidFlag == "True" {

		// First, lets compare the juice to see if it is within the parameter
		slaveLine.compareJuiceValues(masterLine, juiceParameter)

		// Only continue if Juice Values comparisons are passed
		if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareJuiceValues Passed]") {
			// Since MoneyLine doesn't have anything else to compare since Juice is done, let's do Total & Spread
			if slaveLine.LineType == "Total" && masterLine.LineType == "Total" {
				slaveLine.compareTotalLine(masterLine, spreadParameter)
			} else if slaveLine.LineType == "Spread" && masterLine.LineType == "Spread" {
				slaveLine.compareSpreadLine(masterLine, spreadParameter)
			} else if slaveLine.LineType == "TeamTotal" && masterLine.LineType == "TeamTotal" {
				slaveLine.compareTeamTotalLine(masterLine, spreadParameter)
			} else if slaveLine.LineType != "MoneyLine" && masterLine.LineType != "MoneyLine" {
				slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Unable to triage Line Type to Compare")
			}
		} else if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareJuiceValues Failed]") {
			slaveLine.LineStatus = "Skipped"
			slaveLine.FunctionLog = helper.ReplaceParameters(slaveLine.FunctionLog, "[#CompareJuiceValues Failed]", "[#CompareJuiceValues Skipped]")
		} else if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareJuiceValues Skipped]") {
			slaveLine.LineStatus = "Skipped"
		}

		// Let's validate slave line one more time, let's create a new variable so we don't mess up any function logs
		finalValidation := *slaveLine
		finalValidation.ValidateSingleLine()
		if finalValidation.LineStatus == "Error" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Line did not pass Final "+
				"Validation after Line Comparisons")
		}

		// If there are anything attached to the Error Log, we know there is Errors, so Flip Status to Error
		if len(slaveLine.ErrorLog) > 0 {
			slaveLine.LineStatus = "Error"
			slaveLine.FunctionLog = "[#ComparedLines] Error Occurred during Lines Comparison"
		}

	}
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
	if line.BetType != "Master" {
		if helper.FindIfStringInSlice(line.CreatedViaFunction, createdViaFunctionValues) == "False" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: CreatedViaFunction -> Line "+
				"should be created leveraging Create Function")
		}
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

	if line.BetType == "Master" {
		if line.MasterSite == "" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingSite -> Please assign the"+
				"Site Name where Line is found")
		} else if !strings.HasPrefix(line.MasterSite, "http") || !strings.HasSuffix(line.MasterSite, "/") {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingSite -> Incorrectly "+
				"structured, ex. http://247sports.bet/")
		}
	} else if line.BetType == "Slave" {
		if line.SlaveSite == "" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingSite -> Please assign the"+
				"Site Name where Line is found")
		} else if !strings.HasPrefix(line.SlaveSite, "http") || !strings.HasSuffix(line.SlaveSite, "/") {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: BettingSite -> Incorrectly "+
				"structured, ex. http://247sports.bet/")
		}
	}

	// ------------------------------------------------------------------------------- MasterName & MasterPass

	if line.BetType == "Master" {
		if line.MasterName == "" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: MasterName -> Please assign"+
				" the Master Name where Line is found")
		}
		if line.MasterPass == "" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: MasterPass -> Please assign"+
				" the Master Pass where Line is found")
		}
	}

	// ------------------------------------------------------------------------------- MasterName & MasterPass

	if line.BetType == "Slave" {
		if line.SlaveName == "" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: SlaveName -> Please assign"+
				" the Slave Name where Line is found")
		}
		if line.SlavePass == "" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: SlavePass -> Please assign"+
				" the Slave Pass where Line is found")
		}
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
	/*
		if line.TicketID == "" {

			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: TicketID -> Please "+
				"assign the TicketID")

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

		if line.LineSpreadFloat == 0 && line.FavoredUnderdog != "Pick" {
			line.ErrorLog = append(line.ErrorLog, "{betType} {lineType}: LineSpreadFloat -> Need Consistency,"+
				" Value is Not Zero, LineType = Spread & Pick, Float should be Zero")
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

// Check Line against a profile
func (line *Lines) ValidateAgainstProfile(profile Profile) {

	// Things we need to check for are Period and LineType after checking League first
	if line.League == "MLB" {

		// These are the four different LineTypes we will check against the profile
		if line.LineType == "MoneyLine" && profile.SportsSettings.MLB.MoneyLine != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following MLB MoneyLine from this Master"
		} else if line.LineType == "Spread" && profile.SportsSettings.MLB.Spread != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following MLB Spread from this Master"
		} else if line.LineType == "Total" && profile.SportsSettings.MLB.Total != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following MLB Total from this Master"
		} else if line.LineType == "TeamTotal" && profile.SportsSettings.MLB.TeamTotal != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following MLB TeamTotal from this Master"
		}

		// These are the different Periods, for MLB, there is only two
		if line.Period == "" && profile.SportsSettings.MLB.OneFiveInnings != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following MLB 1st Five Innings from this Master"
		} else if line.Period == "" && profile.SportsSettings.MLB.Game != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following MLB Game from this Master"
		}

	} else if line.League == "NBA" {

		// These are the four different LineTypes we will check against the profile
		if line.LineType == "MoneyLine" && profile.SportsSettings.NBA.MoneyLine != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA MoneyLine from this Master"
		} else if line.LineType == "Spread" && profile.SportsSettings.NBA.Spread != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Spread from this Master"
		} else if line.LineType == "Total" && profile.SportsSettings.NBA.Total != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Total from this Master"
		} else if line.LineType == "TeamTotal" && profile.SportsSettings.NBA.TeamTotal != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA TeamTotal from this Master"
		}

		// These are the different Periods, for NBA, there is six
		if line.Period == "" && profile.SportsSettings.NBA.Game != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Game from this Master"
		} else if line.Period == "" && profile.SportsSettings.NBA.OneHalf != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA First Half from this Master"
		} else if line.Period == "" && profile.SportsSettings.NBA.TwoHalf != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Second Half from this Master"
		} else if line.Period == "" && profile.SportsSettings.NBA.OneQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA First Quarter from this Master"
		} else if line.Period == "" && profile.SportsSettings.NBA.TwoQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Second Quarter from this Master"
		} else if line.Period == "" && profile.SportsSettings.NBA.ThreeQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Third Quarter from this Master"
		} else if line.Period == "" && profile.SportsSettings.NBA.FourQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NBA Fourth Quarter from this Master"
		}

	} else if line.League == "NFL" {

		// These are the four different LineTypes we will check against the profile
		if line.LineType == "MoneyLine" && profile.SportsSettings.NFL.MoneyLine != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL MoneyLine from this Master"
		} else if line.LineType == "Spread" && profile.SportsSettings.NFL.Spread != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Spread from this Master"
		} else if line.LineType == "Total" && profile.SportsSettings.NFL.Total != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Total from this Master"
		} else if line.LineType == "TeamTotal" && profile.SportsSettings.NFL.TeamTotal != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL TeamTotal from this Master"
		}

		// These are the different Periods, for NFL, there is six
		if line.Period == "" && profile.SportsSettings.NFL.Game != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Game from this Master"
		} else if line.Period == "" && profile.SportsSettings.NFL.OneHalf != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL First Half from this Master"
		} else if line.Period == "" && profile.SportsSettings.NFL.TwoHalf != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Second Half from this Master"
		} else if line.Period == "" && profile.SportsSettings.NFL.OneQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL First Quarter from this Master"
		} else if line.Period == "" && profile.SportsSettings.NFL.TwoQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Second Quarter from this Master"
		} else if line.Period == "" && profile.SportsSettings.NFL.ThreeQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Third Quarter from this Master"
		} else if line.Period == "" && profile.SportsSettings.NFL.FourQuarter != "Yes" {
			line.FunctionLog = "[#ValidateAgainstProfile] Slave is not following NFL Fourth Quarter from this Master"
		}
	}

	// If there's a function log from above criteria, we know the Slave is not following the Master on Profile
	if strings.HasPrefix(line.FunctionLog, "[#ValidateAgainstProfile]") {
		line.LineStatus = "Skipped"
	} else {
		line.FunctionLog = "[#ValidateAgainstProfile] Slave is following Master on this Line"
	}
}

// This function is to help do Pre-Checks prior to Comparing Lines
func (slaveLine *Lines) ValidateAgainst(masterLine Lines) {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Apply inheritance values from MasterLine
	slaveLine.MasterName = masterLine.MasterName
	slaveLine.MasterPass = masterLine.MasterPass
	slaveLine.MasterSite = masterLine.MasterSite
	slaveLine.MasterTicketID = masterLine.TicketID

	// Let's start by making sure both Slave Line has passed the Single Line Validation
	// We do not yet know how Master Lines will be formatted when they come in so I won't check for their LineStatus yet
	// But I do know that it shouldn't masterLines shouldn't be "New" OR "Error" OR "Ignore"

	if slaveLine.LineStatus != "Validated" && strings.HasPrefix(slaveLine.FunctionLog,
		"[#validateLine]") && len(slaveLine.ErrorLog) > 0 {

		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave {lineType}: has not been Validated"+
			" prior to Pre-Checks")
	}

	// If Rotation Number / LineTypes don't match, I want to set the LineStatus to "Ignored", because there are multiple
	// inputs for each Rotation Number, think of it mapping like X and Y on a graph, both must match for it to compare
	// I don't want to call it out as an "Error" so I'm putting it as "Ignored" so it won't get printed because of this

	// Rotation Numbers - Same between Slave and Master Lines
	if slaveLine.RotationNumber != masterLine.RotationNumber {
		slaveLine.LineStatus = "Ignored"
		slaveLine.FunctionLog = "[#PreCheck Failed] Slave Line Ignored, Rotation Numbers are not matching"
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave {lineType}: Rotation Numbers do "+
			"not match with Master")
	}

	// LineType - Same between Slave, Master
	if slaveLine.LineType != masterLine.LineType {
		slaveLine.LineStatus = "Ignored"
		slaveLine.FunctionLog = "[#PreCheck Failed] Slave Line Ignored, LineTypes are not matching"
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave {lineType}: Line Types do not "+
			"match between Slave and Master")
	}

	// Now, we don't know what LineStatus Master will come in as, but we do know what it should not be: Error, etc,
	if masterLine.LineStatus == "Error" {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave {lineType}: Comparing to an "+
			"Error'ed Master Line")
	}

	// The rest of the Pre-Checks will turn the Slave Line to an Error

	// OverUnder - If both are Total or TeamTotal, make sure they are the same between Slave, Master
	if (slaveLine.LineType == "Total" && masterLine.LineType == "Total") || (slaveLine.LineType == "TeamTotal" &&
		masterLine.LineType == "TeamTotal") {
		if slaveLine.OverUnder != masterLine.OverUnder {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Total: OverUnder Values are not matching")
		}
	}

	/*

		Commented out, spread FavoredUnderdogs should be consistent due to rotation numbers matching, otherwise if it flips
		between +1 and -1 spread, it won't place

		// FavoredUnderdog - If both are Spread, Values should be the same between Slave, Master
		if slaveLine.FavoredUnderdog != masterLine.FavoredUnderdog {
			if slaveLine.FavoredUnderdog != "Pick" && masterLine.FavoredUnderdog != "Pick" {
				if slaveLine.LineType == "Spread" && masterLine.LineType == "Spread" {
					slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Spread: FavoredUnderdog Values"+
						" are not matching")
				}
			}
		}

	*/

	// If LineType == "MoneyLine" and either OverUnder or FavoredUnderdog is not blank
	if slaveLine.LineType == "MoneyLine" && masterLine.LineType == "MoneyLine" {

		if slaveLine.OverUnder != "" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave MoneyLine: OverUnder Values"+
				" shouldn't be populated")
		}

		if masterLine.OverUnder != "" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Master MoneyLine: OverUnder Values"+
				" shouldn't be populated")
		}
	}

	// BetType - Slave is "Slave", Master is "Master"
	if slaveLine.BetType != "Slave" {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave {lineType}: Line for Pre-Check is"+
			" not BetType = 'Slave'")
	}

	if masterLine.BetType != "Master" {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Master {lineType}: Line for Pre-Check is not"+
			" BetType = 'Master'")
	}

	// ErrorLog = Is Empty for Master, Slave will be checked at the end of this function

	if len(masterLine.ErrorLog) != 0 {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave {lineType}: Compared Master Line"+
			" has Error Logs attached")
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, masterLine.ErrorLog...)
	}

	// Now we can start towards ending this function

	// If there are any ErrorLogs, that means it failed whether a check from Validation or this Pre-Check
	if len(slaveLine.ErrorLog) > 0 {

		for i := range slaveLine.ErrorLog {
			slaveLine.ErrorLog[i] = helper.ReplaceParameters(slaveLine.ErrorLog[i], "{lineType}",
				slaveLine.LineType)
		}

		if slaveLine.LineStatus == "Ignored" {
			slaveLine.FunctionLog = "[#Pre-Check Failed] Ignored due to either Rotation Number or LineType not" +
				" matching"
		} else if slaveLine.LineStatus == "Validated" && slaveLine.FunctionLog == "[#validateLine] New "+
			"Slave Line passed Validation" {

			// If it's in here, we know that the Slave Line passed Single Line Validation but failed Pre-Checks
			slaveLine.LineStatus = "Error"
			slaveLine.FunctionLog = "[#Pre-Check Failed] Slave Line Passed Singular Validation Checks, But " +
				"Failed Pre-Checks prior to Comparisons"

		} else if slaveLine.LineStatus != "Ignored" {
			// If in here, we know that the Slave Line failed Single Line Validation, which should fail Pre-Checks
			// The only exception is if Rotation and LineType's don't match, which should be ignored
			slaveLine.LineStatus = "Error"
			slaveLine.FunctionLog = "[#Pre-Check Failed] Slave Line Failed Singular Validation Checks, And " +
				"therefore Pre-Checks, prior to Comparisons"
		}

	} else if slaveLine.LineStatus != "Ignored" {

		// This means that the Slave Line has passed Validation and Pre-Check Validation
		slaveLine.LineStatus = "Validated"
		slaveLine.FunctionLog = "[#Pre-Check Passed] Slave Line has Passed Singular Line & Pre-Check Validation"
	}

}

// ----------------------------------- Helper Function #1 - Compare Juice values ------------------------------------ //
// -------------- Goal of this function is to take Slave and Master Line Struct & Compare Juice values -------------- //

func (slaveLine *Lines) compareJuiceValues(masterLine Lines, juiceParam float64) {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Assigning these to new variables for easier to read
	slaveJuiceFloat := slaveLine.LineJuiceFloat
	masterJuiceFloat := masterLine.LineJuiceFloat

	// --------------------------------------------------------

	// First, lets see if any of them are EVEN values
	if slaveJuiceFloat == 100 || masterJuiceFloat == 100 {

		// Remember I set these Juice values to positive 100 if they are EVEN, so I will need to catch it to see whether
		// to use +100 or -100 to compare against the juiceParam, can't use Absolute Value because they can be inverted

		// If Slave Juice is EVEN (100) and Master Juice is negatives (-120), change Slave Juice to -100, if
		// Master Juice is positive and over, or equal to +100, then leaving Slave as +100 as it is will be fine
		if slaveJuiceFloat == 100 && masterJuiceFloat < 0 {
			slaveJuiceFloat = slaveJuiceFloat * (-1)
			slaveLine.LineJuiceFloat = slaveJuiceFloat
		}

		// If Master Juice is EVEN (100) and Slave Juice is negatives (-120), change Master Juice to -100, if
		// Slave Juice is positive and over, or equal to +100, then leaving Master as +100 as it is will be fine
		if slaveJuiceFloat < 100 && masterJuiceFloat == 100 {
			masterJuiceFloat = masterJuiceFloat * (-1)
		}

	}

	// Lets first see if they are inverted juice values (one negative, one positive), inverted needs to be within params
	// We know if we multiply them, and it is negative, then it is inverted
	if slaveJuiceFloat*masterJuiceFloat < 0 {

		// Let's create two variables to track the positive diff (from +100) and negative diff (from -100)
		var positiveDiff float64
		var negativeDiff float64

		// Now we need to find out which one is the positive and which one is the negative
		if slaveJuiceFloat >= 100 && masterJuiceFloat <= -100 {
			positiveDiff = slaveJuiceFloat - 100         // Since Slave Juice is positive, 105 - 100 = 5
			negativeDiff = (masterJuiceFloat + 100) * -1 // Since Master Juice is negative, -105 + 100 = -5 * -1 = 5
		}

		if slaveJuiceFloat <= -100 && masterJuiceFloat >= 100 {
			positiveDiff = masterJuiceFloat - 100       // Since Master Juice is positive, 105 - 100 = 5
			negativeDiff = (slaveJuiceFloat + 100) * -1 // Since Slave Juice is neg, -105 + 100 = -5 * -1 = 5
		}

		// If subtracted up and it is less than or equal to the juice parameter, we know it is within the parameters
		if positiveDiff+negativeDiff <= juiceParam {
			slaveLine.FunctionLog = fmt.Sprintf("[#CompareJuiceValues Passed] Master Juice (%v) w/ Juice "+
				"Parameter (%v) vs Slave Juice (%v)", masterJuiceFloat, juiceParam, slaveJuiceFloat)
		} else {
			slaveLine.FunctionLog = fmt.Sprintf("[#CompareJuiceValues Failed] Master Juice (%v) w/ Juice "+
				"Parameter (%v) vs Slave Juice (%v)", masterJuiceFloat, juiceParam, slaveJuiceFloat)
		}
	}

	// Now if they are not inverted, multiplying them will be positive
	if slaveJuiceFloat*masterJuiceFloat > 0 {
		// If we know that Slave Juice Float is greater than Master, then we know we should place if within parameter
		if masterJuiceFloat <= slaveJuiceFloat+juiceParam {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareJuiceValues Passed] Master Juice (%v) w/ Juice "+
				"Parameter (%v) vs Slave Juice (%v)", masterJuiceFloat, juiceParam, slaveJuiceFloat)

			// If we know that Slave Juice Float is less than Master, then we know we should skip the line
		} else {
			slaveLine.FunctionLog = fmt.Sprintf("[#CompareJuiceValues Skipped] Master Juice (%v) w/ Juice "+
				"Parameter (%v) vs Slave Juice (%v)", masterJuiceFloat, juiceParam, slaveJuiceFloat)
		}
	}

	// Since MoneyLine doesn't need spread comparison, we can "Authorized" or "Skipped" it now
	if slaveLine.LineType == "MoneyLine" && masterLine.LineType == "MoneyLine" {

		// This prefix means that it passed Juice Comparisons, Authorized if it's a MoneyLine
		if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareJuiceValues Passed]") {
			slaveLine.FunctionLog = helper.ReplaceParameters(slaveLine.FunctionLog,
				"[#CompareJuiceValues Passed]", "[#CompareJuiceValues Authorized]")
			slaveLine.LineStatus = "Authorized"
		}

		// This prefix means that it did not pass Juice Comparisons
		if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareJuiceValues Failed]") {
			slaveLine.FunctionLog = helper.ReplaceParameters(slaveLine.FunctionLog,
				"[#CompareJuiceValues Failed]", "[#CompareJuiceValues Skipped]")
			slaveLine.LineStatus = "Skipped"
		}
	}

}

func (slaveLine *Lines) compareSpreadLine(approvedLine Lines, spreadParam float64) {

	// Regardless if it is Favored or Underdog, it'll use the same function
	/*
		eitherValueIsPick := "False"
		if slaveLine.FavoredUnderdog == "Pick" || approvedLine.FavoredUnderdog == "Pick" {
			eitherValueIsPick = "True"
		}

	*/

	// Since Rotation Numbers are already checked, we do not need to have the same FavoredUnderdog
	//if slaveLine.FavoredUnderdog == approvedLine.FavoredUnderdog || eitherValueIsPick == "True" {
	if approvedLine.LineSpreadFloat <= slaveLine.LineSpreadFloat+spreadParam {

		slaveLine.FunctionLog = fmt.Sprintf("[#CompareSpreadLine Authorized] Master Spread (%v) w/ "+
			"Spread Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
			slaveLine.LineSpreadFloat)
	} else {

		slaveLine.FunctionLog = fmt.Sprintf("[#CompareSpreadLine Skipped] Master Spread (%v) w/ Spread "+
			"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
			slaveLine.LineSpreadFloat)
	}
	//}

	// This prefix means that it passed Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareSpreadLine Authorized]") {
		slaveLine.LineStatus = "Authorized"
	}

	// This prefix means that it did not pass Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareSpreadLine Skipped]") {
		slaveLine.LineStatus = "Skipped"
	}

}

func (slaveLine *Lines) compareTotalLine(approvedLine Lines, spreadParam float64) {

	if slaveLine.OverUnder == "Over" && approvedLine.OverUnder == "Over" {
		if approvedLine.LineSpreadFloat >= slaveLine.LineSpreadFloat-spreadParam {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Authorized] Master Over (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)

		} else {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Skipped] Master Over (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)

		}

	} else if slaveLine.OverUnder == "Under" && approvedLine.OverUnder == "Under" {
		if approvedLine.LineSpreadFloat <= slaveLine.LineSpreadFloat+spreadParam {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Authorized] Master Under (%v) w/ Spread"+
				" Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)

		} else {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Skipped] Master Under (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)
		}

	} else {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Total: Inverted OverUnder with "+
			"Authorized bet")
	}

	// This prefix means that it passed Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareTotalLine Authorized]") {
		slaveLine.LineStatus = "Authorized"
	}

	// This prefix means that it did not pass Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareTotalLine Skipped]") {
		slaveLine.LineStatus = "Skipped"
	}

}

func (slaveLine *Lines) compareTeamTotalLine(approvedLine Lines, spreadParam float64) {

	if slaveLine.OverUnder == "Over" && approvedLine.OverUnder == "Over" {
		if approvedLine.LineSpreadFloat >= slaveLine.LineSpreadFloat-spreadParam {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Authorized] Master Over (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)

		} else {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Skipped] Master Over (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)

		}

	} else if slaveLine.OverUnder == "Under" && approvedLine.OverUnder == "Under" {
		if approvedLine.LineSpreadFloat <= slaveLine.LineSpreadFloat+spreadParam {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Authorized] Master Under (%v) w/ Spread"+
				" Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)

		} else {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareTotalLine Skipped] Master Under (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)
		}

	} else {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Total: Inverted OverUnder with "+
			"Authorized bet")
	}

	// This prefix means that it passed Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareTotalLine Authorized]") {
		slaveLine.LineStatus = "Authorized"
	}

	// This prefix means that it did not pass Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareTotalLine Skipped]") {
		slaveLine.LineStatus = "Skipped"
	}

}
