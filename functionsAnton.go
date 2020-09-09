package anton

import (
	"fmt"
	"strings"
)

// This function intakes the Master and user, and a string to format, and returns a slice of string that is formatted
// The parameters to replace are {SportID}, {LeagueID}, {PeriodID}
// ex. Intake string can look like this: `"{IdSport":{LeagueID},"Period":{PeriodID}` from Smash66
func FormatStringSportLeaguePeriod(master Master, slave Slave, stringToFormat string) []string {

	// Declare the helper struct to access the helper functions
	var helper Helper

	// Create the Slice of Strings that we will return
	var returnStringSlice []string

	// Let's create a variable so it's easier to read
	sportsDict := slave.SiteDictionary.SportsDict

	// Grab the Scope of the Lines by looping through all the attached Master Lines so I know where to go on the site
	for i := range master.MasterLines {

		// Create a copy of the string we need to format, append to the returnSlice and reset every loop run,
		tempString := stringToFormat

		// First, we need to sort it out based on the Sport, ex. Football
		if master.MasterLines[i].Sport == "Football" {
			// Add in the SportID, which we know is "Football" on this conditional
			tempString = helper.ReplaceParameters(tempString, "{SportID}", sportsDict.Football.ID)

			// Sort it out based on the different leagues, this is NFL
			if master.MasterLines[i].League == "NFL" {
				// Add in the LeagueID, which we know is "NFL" on this conditional
				tempString = helper.ReplaceParameters(tempString, "{LeagueID}", sportsDict.Football.NFL.ID)

				// Now we have to add in the Period of the bets (Game, 1st Half, etc)
				if master.MasterLines[i].Period == "Game" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.Game)
				} else if master.MasterLines[i].Period == "1st Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.OneHalf)
				} else if master.MasterLines[i].Period == "2nd Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.TwoHalf)
				} else if master.MasterLines[i].Period == "1st Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.OneQuarter)
				} else if master.MasterLines[i].Period == "2nd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.TwoQuarter)
				} else if master.MasterLines[i].Period == "3rd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.ThreeQuarter)
				} else if master.MasterLines[i].Period == "4th Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.NFL.Period.FourQuarter)
				}
			}
		} else if master.MasterLines[i].Sport == "Basketball" {
			// Add in the SportID, which we know is "Basketball" on this conditional
			tempString = helper.ReplaceParameters(tempString, "{SportID}", sportsDict.Basketball.ID)

			// Sort it out based on the different leagues, this is NBA
			if master.MasterLines[i].League == "NBA" {
				// Add in the LeagueID, which we know is "NFL" on this conditional
				tempString = helper.ReplaceParameters(tempString, "{LeagueID}", sportsDict.Basketball.NBA.ID)

				// Now we have to add in the Period of the bets (Game, 1st Half, etc)
				if master.MasterLines[i].Period == "Game" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.Game)
				} else if master.MasterLines[i].Period == "1st Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.OneHalf)
				} else if master.MasterLines[i].Period == "2nd Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.TwoHalf)
				} else if master.MasterLines[i].Period == "1st Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.OneQuarter)
				} else if master.MasterLines[i].Period == "2nd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.TwoQuarter)
				} else if master.MasterLines[i].Period == "3rd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.ThreeQuarter)
				} else if master.MasterLines[i].Period == "4th Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.NBA.Period.FourQuarter)
				}
			}
		} else if master.MasterLines[i].Sport == "Baseball" {
			// Add in the SportID, which we know is "Baseball" on this conditional
			tempString = helper.ReplaceParameters(tempString, "{SportID}", sportsDict.Baseball.ID)

			// Sort it out based on the different leagues, this is MLB
			if master.MasterLines[i].League == "MLB" {
				// Add in the LeagueID, which we know is "MLB" on this conditional
				tempString = helper.ReplaceParameters(tempString, "{LeagueID}", sportsDict.Baseball.MLB.ID)

				// Now we have to add in the Period of the bets (Game, 1st 5 Innings, etc)
				if master.MasterLines[i].Period == "Game" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Baseball.MLB.Period.Game)
				} else if master.MasterLines[i].Period == "1st 5 Innings" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Baseball.MLB.Period.OneFiveInnings)
				}
			}
		}

		// Now if the tempString still has {SportID}, {LeagueID}, or {PeriodID}, we know Master Line is either
		// incorrectly filled out, or we have not yet gotten the Sport/League/Period coded in
		if strings.Contains(tempString, "{SportID}") {
			fmt.Println("Could not finish filling in " + tempString + ", still contains the Parameter {SportID}")
		} else if strings.Contains(tempString, "{LeagueID}") {
			fmt.Println("Could not finish filling in " + tempString + ", still contains the Parameter {LeagueID}")
		} else if strings.Contains(tempString, "{PeriodID}") {
			fmt.Println("Could not finish filling in " + tempString + ", still contains the Parameter {PeriodID}")
		} else {
			returnStringSlice = append(returnStringSlice, tempString)
		}
	}

	return returnStringSlice
}
