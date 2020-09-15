package anton

import (
	"fmt"
	"strings"
)

// MASTER DONE

// These are the functions that help us ingest values, will be used to format when Creating Lines in the Methods
// Used in both methodsMaster.go and methodsSlave.go

// Possible Sport Values can be found in configurableParameters.go
func formatSport(sport string) string {

	// Declare the helper struct to access the helper functions
	var helper Helper

	returnSport := strings.ToUpper(helper.ReplaceParameters(sport, "\u00a0", "", " ", ""))

	if helper.FindIfStringInSlice(returnSport, sportValuesBaseball) != "False" {
		returnSport = "Baseball"
	} else if helper.FindIfStringInSlice(returnSport, sportValuesBasketball) != "False" {
		returnSport = "Basketball"
	} else if helper.FindIfStringInSlice(returnSport, sportValuesFootball) != "False" {
		returnSport = "Football"
	} else {
		fmt.Println("Error, can not format Period with a value of " + returnSport)
		returnSport = "Undefined"
	}

	return returnSport
}

// Possible Period Values can be found in configurableParameters.go
func formatLeague(league string) string {
	// Declare the helper struct to access the helper functions
	var helper Helper
	returnLeague := strings.ToUpper(helper.ReplaceParameters(league, "\u00a0", "", " ", ""))

	// The different Leagues are "MLB", "NBA", "NFL"

	if helper.FindIfStringInSlice(returnLeague, leagueValuesMLB) != "False" {
		returnLeague = "MLB"
	} else if helper.FindIfStringInSlice(returnLeague, leagueValuesNBA) != "False" {
		returnLeague = "NBA"
	} else if helper.FindIfStringInSlice(returnLeague, leagueValuesNFL) != "False" {
		returnLeague = "NFL"
	} else {
		fmt.Println("Error, can not format Period with a value of " + returnLeague)
		returnLeague = "Undefined"
	}

	return returnLeague
}

// Possible Period Values can be found in configurableParameters.go
func formatPeriod(period string) string {
	// Declare the helper struct to access the helper functions
	var helper Helper
	returnPeriod := strings.ToUpper(helper.ReplaceParameters(period, "\u00a0", "", " ", ""))

	if helper.FindIfStringInSlice(returnPeriod, periodValuesGame) != "False" {
		returnPeriod = "Game"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesOneFiveInnings) != "False" {
		returnPeriod = "1st Five Innings"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesOneHalf) != "False" {
		returnPeriod = "1st Half"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesTwoHalf) != "False" {
		returnPeriod = "2nd Half"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesOneQuarter) != "False" {
		returnPeriod = "1st Quarter"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesTwoQuarter) != "False" {
		returnPeriod = "2nd Quarter"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesThreeQuarter) != "False" {
		returnPeriod = "3rd Quarter"
	} else if helper.FindIfStringInSlice(returnPeriod, periodValuesFourQuarter) != "False" {
		returnPeriod = "4th Quarter"
	} else {
		fmt.Println("Error, can not format Period with a value of " + returnPeriod)
		returnPeriod = "Undefined"
	}

	return returnPeriod
}
