package anton

// The reason I want to centralize this is because when maintaining, I only want to edit this file

// ------------------------------------------ Configurable Parameters ----------------------------------------------- //

// Slice of EVEN values for Juice that is ToUpperCased, if it's a value in here, then set float to +100
var lineJuiceEvenValues = []string{"-100", "+100", "100", "EV", "EVEN", "PK", "PICK"}

// Slice of EVEN values for Spread that is ToUpperCased, if it's a value in here, then set float to 0
var lineSpreadEvenValues = []string{"0", "EV", "EVEN", "PK", "PICK"}

// Slice of Possible Period Values
var OneHalf = []string{}

// ------------------------------- Configurable Parameters That Shouldn't Change Much ------------------------------- //

// These Line Properties have set values that I will be validating against, listing the values here
var lineTypeValues = []string{"MoneyLine", "Total", "Spread", "TeamTotal"}
var homeAwayValues = []string{"Away", "Home"}
var favoredUnderdogValues = []string{"Favored", "Underdog"}
var overUnderValues = []string{"Over", "Under"}
var lineStatusValues = []string{"New", "Validated", "Authorized", "Placed", "Error", "Skipped", "Ignored"}
var createdViaFunctionValues = []string{"True"}
var betTypeValues = []string{"Master", "Slave"}

// Slice of Properties my printLineValues function will print out
var printTheseProperties = []string{
	"BetType",
	"LineStatus",
	"BettingUser",
	"BettingSite",
	"RotationNumber",
	"Team",
	"Sport",
	"League",
	"Period",
	"LineType",
	"LineSpread",
	"LineSpreadFloat",
	"LineJuice",
	"LineJuiceFloat",
	"HomeAway",
	"FavoredUnderdog",
	"OverUnder",
	"FunctionLog",
	"ErrorLog",
}

// ----------------------------------- Functions to return the slice values above ----------------------------------- //

// Return slice of even Juice values
func GetJuiceEvenValues() []string {
	return lineJuiceEvenValues
}

// Return slice of even Spread vales
func GetSpreadEvenValues() []string {
	return lineSpreadEvenValues
}

func GetLineTypeValues() []string {
	return lineTypeValues
}

func GetHomeAwayValues() []string {
	return homeAwayValues
}

func GetFavoredUnderdogValues() []string {
	return favoredUnderdogValues
}

func GetOverUnderValues() []string {
	return overUnderValues
}

func GetLineStatusValues() []string {
	return lineStatusValues
}

func GetCreatedViaFunctionValues() []string {
	return createdViaFunctionValues
}

func GetBetTypeValues() []string {
	return betTypeValues
}

func GetPrintProperties() []string {
	return printTheseProperties
}
