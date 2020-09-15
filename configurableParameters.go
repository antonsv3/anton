package anton

// The reason I want to centralize this is because when maintaining, I only want to edit this file

// ------------------------------------------ Configurable Parameters ----------------------------------------------- //

// Use ToUpperCased and RemoveSpaces, then FindIfStringInSlice to see if it is in list of possible values

// Slice of EVEN values for Juice that is ToUpperCased, if it's a value in here, then set float to +100
var lineJuiceEvenValues = []string{"-100", "+100", "100", "EV", "EVEN", "PK", "PICK"} // Set Float to +100

// Slice of EVEN values for Spread that is ToUpperCased, if it's a value in here, then set float to 0
var lineSpreadEvenValues = []string{"0", "EV", "EVEN", "PK", "PICK"}

// ------------------------------------------------------------------------------- Period - Possible Values
// "Game"
var periodValuesGame = []string{"GAME", "FULLGAME"}

// "First 5 Innings"
var periodValuesOneFiveInnings = []string{"1STFIVEINNINGS", "ONEFIVEINNINGS", "FIRSTFIVEINNINGS"}

// "First Half"
var periodValuesOneHalf = []string{"1STHALF", "1STH", "1H", "FIRSTHALF", "ONEHALF", "1HALF"}

// "Second Half"
var periodValuesTwoHalf = []string{"2NDHALF", "2NDH", "2H", "SECONDHALF", "TWOHALF", "2HALF"}

// "First Quarter"
var periodValuesOneQuarter = []string{"1STQUARTER", "1STQ", "1Q", "ONEQUARTER", "FIRSTQUARTER", "1QUARTER"}

// "Second Quarter"
var periodValuesTwoQuarter = []string{"2NDQUARTER", "2NDQ", "2Q", "TWOQUARTER", "SECONDQUARTER", "2QUARTER"}

// "Third Quarter"
var periodValuesThreeQuarter = []string{"3RDQUARTER", "3RDQ", "3Q", "THREEQUARTER", "THRIDQUARTER", "3QUARTER"}

// "Fourth Quarter"
var periodValuesFourQuarter = []string{"4THQUARTER", "4THQ", "4Q", "FOURQUARTER", "FOURTHQUARTER", "4QUARTER"}

// ------------------------------------------------------------------------------- Sport - Possible Values
// "Baseball"
var sportValuesBaseball = []string{"BASEBALL"}

// "Basketball"
var sportValuesBasketball = []string{"BASKETBALL"}

// "Football"
var sportValuesFootball = []string{"FOOTBALL", "AMERICANFOOTBALL"}

// ------------------------------------------------------------------------------- League - Possible Values
// "MLB"
var leagueValuesMLB = []string{"MLB", "MAJORLEAGUEBASEBALL"}

// "NBA"
var leagueValuesNBA = []string{"NBA", "NATIONALBASKETBALLASSOCIATION"}

// "NFL"
var leagueValuesNFL = []string{"NFL", "NATIONALFOOTBALLLEAGUE"}

// "C Football"
var leagueValuesCFootball = []string{"COLLEGEFOOTBALL"}

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
