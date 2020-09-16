package anton

// I NEED TO ADD JUICE COMPARISONS FOR ALL LINES, NOT ONLY MONEYLINE
// I NEED TO ADD PERIOD COMPARISON AROUND LINE 390

// ----------------------------------------- The Methods in this file are: ------------------------------------------ //
// CompareSlaveLineToMasterLine()      	<< This is Overarching Method that will be used outside of this function
// - compareJuiceValues()		<< Used by the Overarching Method, not able to be referenced outside of this package
// - compareTotalLine()			<< Used by the Overarching Method, not able to be referenced outside of this package
// - compareSpreadLine()		<< Used by the Overarching Method, not able to be referenced outside of this package
// - compareTeamTotalLine()		<< Used by the Overarching Method, not able to be referenced outside of this package

// ValidateAgainst()
// PrintComparedLines()

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ----------------- Overarching helper function to call the three functions below conditionally -------------------- //
// -------------------------- (compareMoneyLine, compareTotalsLine, compareSpreadsLine) ----------------------------- //
// ------------- Goal of these function are to take a Slave line struct, Master line struct and Compare ------------- //
// -------------- If Slave Line is better than Master Line, add it to currentUser's Authorized Lines ---------------- //

func (slaveLine *Lines) CompareSlaveLineToMasterLine(masterLine Lines, slave Slave, profile Profile) {

	// Let's start by getting our parameters from the profile

	// This variable is our parameter for what is the maximum difference of Juice between Slave and Master
	juiceParameter, _ := strconv.ParseFloat(profile.JuiceParameter, 32)

	// Because GoLang uses 0 if there is an error when converting to float/integers, we need to see if error or not
	if juiceParameter == 0 && profile.JuiceParameter != "0" {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Could Not Parse Juice Parameter to Float Value")
	}

	// This variable is our parameter for what is the maximum difference of Spread values between Slave and Master
	spreadParameter, _ := strconv.ParseFloat(profile.SpreadParameter, 32)

	// Because GoLang uses 0 if there is an error when converting to float/integers, we need to see if error or not
	if spreadParameter == 0 && profile.SpreadParameter != "0" {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Could Not Parse Spread Parameter to Float Value")
	}

	// This flag is to ensure that we have met all other criteria's prior to comparing, default "False"
	preChecksValidFlag := "False"

	// Now lets start these checks, Let's validate both Lines to see if values are populated correctly

	slaveLine.ValidateSingleLine()
	masterLine.ValidateSingleLine()

	// Next, let's compare it to the profiles to see whether the Slave is following the master on these lines
	if len(slave.Profiles) > 1 {
		slaveLine.ValidateAgainstProfile(slave.Profiles[0])
	} else {
		slaveLine.LineStatus = "Error"
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave does not have a Profile Attached")
	}

	// If slaveLine passes Validation, then they're populated correctly, then we can pre-check versus the Master Line
	if slaveLine.LineStatus == "Validated" {
		slaveLine.ValidateAgainst(masterLine)
	}

	// This will print if it failed the Single Validation or the Pre Check Comparisons
	if slaveLine.LineStatus == "Error" || masterLine.LineStatus == "Error" {
		slaveLine.PrintComparedLines(masterLine)
	}

	// Change Flag to True if both lines passed the single line Validation and is still validated after Pre-Checks
	if slaveLine.LineStatus == "Validated" && masterLine.LineStatus == "Validated" {
		preChecksValidFlag = "True"
	}

	// Now that we've checked, we can compare the two Lines now by calling the functions below

	if preChecksValidFlag == "True" && slaveLine.LineStatus != "Ignored" {

		// The 4 Helper Functions return the Slave Line, they're not Methods because they are part of this Method

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
				slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Unable to triage Lines to Compare")
			}
		}

		// Let's validate slave line one more time, let's create a new variable so we don't mess up any function logs
		finalValidation := slaveLine
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

	// FIX THIS

	// Now let's attach the Slave Line back to the Slave, Slice depends on the the status of the comparison
	if slaveLine.LineStatus == "Authorized" {
		slave.AuthorizedLines = append(slave.AuthorizedLines, *slaveLine)
	} else if slaveLine.LineStatus == "Skipped" {
		slave.AuthorizedLines = append(slave.AuthorizedLines, *slaveLine)
	} else if slaveLine.LineStatus == "Error" {
		slave.AuthorizedLines = append(slave.AuthorizedLines, *slaveLine)
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

// -------------------------------- Helper Function #2 - Compare Total Line values ---------------------------------- //
// --------------- Goal of this function is to take Slave and Approved Struct & compare Total values ---------------- //
// ------------------- If approved, attach the Slave Line Struct to currentUser's Betting Struct -------------------- //

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

// -------------------------------- Helper Function #2 - Compare Total Line values ---------------------------------- //
// --------------- Goal of this function is to take Slave and Approved Struct & compare Total values ---------------- //
// ------------------- If approved, attach the Slave Line Struct to currentUser's Betting Struct -------------------- //

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

// ------------------------------- Helper function #3 - Compare Spread Line values ---------------------------------- //
// -------------- Goal of this function is to take Slave and Approved Struct & compare Spread values ---------------- //
// ------------------- If approved, attach the Slave Line Struct to currentUser's Betting Struct -------------------- //

func (slaveLine *Lines) compareSpreadLine(approvedLine Lines, spreadParam float64) {

	// Regardless if it is Favored or Underdog, it'll use the same function

	if slaveLine.FavoredUnderdog == approvedLine.FavoredUnderdog {
		if approvedLine.LineSpreadFloat <= slaveLine.LineSpreadFloat+spreadParam {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareSpreadLine Authorized] Master Spread (%v) w/ "+
				"Spread Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)
		} else {

			slaveLine.FunctionLog = fmt.Sprintf("[#CompareSpreadLine Skipped] Master Spread (%v) w/ Spread "+
				"Parameter (%v) vs. Slave Spread: (%v)", approvedLine.LineSpreadFloat, spreadParam,
				slaveLine.LineSpreadFloat)
		}
	} else {
		slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Total: Inverted OverUnder with "+
			"Authorized bet")
	}

	// This prefix means that it passed Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareSpreadLine Authorized]") {
		slaveLine.LineStatus = "Authorized"
	}

	// This prefix means that it did not pass Juice Comparisons
	if strings.HasPrefix(slaveLine.FunctionLog, "[#CompareSpreadLine Skipped]") {
		slaveLine.LineStatus = "Skipped"
	}

}

// This function is to help do Pre-Checks prior to Comparing Lines

func (slaveLine *Lines) ValidateAgainst(masterLine Lines) {

	// Declare the helper struct to access the helper functions
	var helper Helper

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

	// FavoredUnderdog - If both are Spread, Values should be the same between Slave, Master
	if slaveLine.FavoredUnderdog != masterLine.FavoredUnderdog {
		if slaveLine.LineType == "Spread" && masterLine.LineType == "Spread" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave Spread: FavoredUnderdog Values"+
				" are not matching")
		}
	}

	// If LineType == "MoneyLine" and either OverUnder or FavoredUnderdog is not blank
	if slaveLine.LineType == "MoneyLine" && masterLine.LineType == "MoneyLine" {
		if slaveLine.FavoredUnderdog != "" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave MoneyLine: FavoredUnderdog"+
				" Values should not be populated")
		}

		if slaveLine.OverUnder != "" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Slave MoneyLine: OverUnder Values"+
				" shouldn't be populated")
		}

		if masterLine.FavoredUnderdog != "" {
			slaveLine.ErrorLog = append(slaveLine.ErrorLog, "Master MoneyLine: FavoredUnderdog Values"+
				" should not be populated")
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

// Helper Function that will help format and print two Line's values side by side for comparison and testing
func (slaveLine Lines) PrintComparedLines(masterLine Lines) {

	if slaveLine.LineStatus != "Ignored" {

		// Formatting the Print Header
		tempHeader := centerString("Begin Comparison", 20)
		tempHeader = "< ---------------------------------------" + tempHeader +
			"--------------------------------------- >"

		fmt.Println("")
		fmt.Println(centerString(tempHeader, 127))
		fmt.Println("")

		propertiesToPrint := GetPrintProperties()

		// Hold the field and values of the Slave Line Struct and Master Line Struct
		slaveProperties := reflect.TypeOf(slaveLine)
		slaveValues := reflect.ValueOf(slaveLine)
		masterValues := reflect.ValueOf(masterLine)

		// Hold the amount of fields of the Line Struct to loop over
		num := slaveProperties.NumField()

		// Loop over all the properties in the slice to print
		for i := 0; i < len(propertiesToPrint); i++ {

			// We know the first property in the slice is propertiesToPrint[i], now lets iterate through all properties
			for j := 0; j < num; j++ {

				// The current Slave Line Property and Value in this loop
				slaveProperty := slaveProperties.Field(j)
				slaveValue := slaveValues.Field(j)

				// The current Master Line Property and Value in this loop
				//masterProperty := masterProperties.Field(j)
				masterValue := masterValues.Field(j)

				// Now we know we have the right property and in the correct order, lets format our print inside this If
				if slaveProperty.Name == propertiesToPrint[i] && slaveProperty.Name != "FunctionLog" &&
					slaveProperty.Name != "ErrorLog" {

					centerPrint := "     <------------" + centerString(slaveProperty.Name, 20) +
						"------------>     "

					// Only Print FavoredUnderdog if LineType == "Spread", and OverUnder if LineType == "Total"
					if slaveProperty.Name != "FavoredUnderdog" && slaveProperty.Name != "OverUnder" {
						fmt.Printf("     %-30v%v%+30v\n", masterValue, centerPrint, slaveValue)
					} else if slaveProperty.Name == "FavoredUnderdog" && slaveLine.LineType == "Spread" {
						fmt.Printf("     %-30v%v%+30v\n", masterValue, centerPrint, slaveValue)
					} else if slaveProperty.Name == "OverUnder" && (slaveLine.LineType == "Total" ||
						slaveLine.LineType == "TeamTotal") {

						fmt.Printf("     %-30v%v%+30v\n", masterValue, centerPrint, slaveValue)
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
