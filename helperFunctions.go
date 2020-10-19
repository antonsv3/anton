package anton

// SLAVE CONSOLIDATION CHECKED

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// We'll create a new struct and have these as methods to that
type Helper struct{}

// Helper function to remove all spaces and new lines
func (helper Helper) StringRemoveSpaces(stringIntake string) string {
	temp := regexp.MustCompile(`\s+`)
	s := temp.ReplaceAllString(stringIntake, "")
	return s
}

// Helper function to replace parameters in a string
func (helper Helper) ReplaceParameters(format string, args ...string) string {
	r := strings.NewReplacer(args...)
	return r.Replace(format)
}

// Helper function that compares strings to see if they're in a slice, and returns the value if it is, otherwise "False"
func (helper Helper) FindIfStringInSlice(inputString string, inputSlice []string) string {
	returnString := ""
	for k, v := range inputSlice {
		if inputString == v {
			returnString = inputSlice[k]
		}
	}
	if returnString == "" {
		// Notice this value is a string and not a boolean, since I declared this function will only return a string
		returnString = "False"
	}
	return returnString
}

// GetStringInBetween returns empty string if no start or end string found
func (helper Helper) GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}

// Helper function that checks if a float value is negative, positive, or zero, returns "Negative", "Positive", or "Zero"
func (helper Helper) FloatNegativePositiveZero(floatValue float64) string {
	var returnString = ""
	if floatValue > 0 {
		returnString = "Positive"
	} else if floatValue < 0 {
		returnString = "Negative"
	} else if floatValue == 0 {
		returnString = "Zero"
	}
	return returnString
}

func (helper Helper) IsEven(number int) bool {
	return number%2 == 0
}

func (helper Helper) IsOdd(number int) bool {
	// Odd should return not even, We cannot check for 1 remainder because it fails for negative numbers
	return !helper.IsEven(number)
}

func (helper Helper) BytesToString(data []byte) string {
	return string(data[:])
}

func (helper Helper) PrintCollectorTraffic(currentCollector *colly.Collector, section string, startTime time.Time) {

	if section == "a" {

		currentCollector.OnRequest(func(r *colly.Request) {
			fmt.Println("(a) Headers = ", r.Headers)
			fmt.Println("(a) Visiting = ", r.URL)
		})

		currentCollector.OnResponse(func(r *colly.Response) {
			fmt.Println("(a) Response received", r.StatusCode)
			//	fmt.Println("(a) Authentication Finished =", time.Since(startTime))
		})

		currentCollector.OnHTML("body", func(temp *colly.HTMLElement) {
			fmt.Println("(a) Authentication Finished =", time.Since(startTime))
		})
	}

	if section == "b" {

		currentCollector.OnRequest(func(r *colly.Request) {
			fmt.Println("(b) Headers = ", r.Headers)
			fmt.Println("(b) Visiting = ", r.URL)
		})

		currentCollector.OnResponse(func(r *colly.Response) {
			fmt.Println("(b) Response received", r.StatusCode)
		})

		currentCollector.OnHTML("body", func(temp *colly.HTMLElement) {
			fmt.Println("(b) Landed on the tokens page =", time.Since(startTime))
		})
	}

	if section == "c" {

		currentCollector.OnRequest(func(r *colly.Request) {
			fmt.Println("(c) Headers = ", r.Headers)
			fmt.Println("(c) Visiting = ", r.URL)
		})

		currentCollector.OnResponse(func(r *colly.Response) {
			fmt.Println("(c) Response received", r.StatusCode)
		})

		currentCollector.OnHTML("body", func(temp *colly.HTMLElement) {
			fmt.Println("(c) Landed on the Lines page =", time.Since(startTime))
		})
	}
}

// Helper function that checks if a string value is negative, positive, or zero, returns "Negative", "Positive", "Error"
// I'm going to return Zeros as Positive, it will get reformatted and error checked by other functions
func (helper Helper) StringNegativePositiveZero(stringValue string) string {
	stringValue = helper.ReplaceParameters(stringValue, "½", ".5", "\u00a0", "", "", "")

	var returnString string
	var isZeroAnError string
	var tempFloat float64

	juiceEvenValues := GetJuiceEvenValues()
	spreadEvenValues := GetSpreadEvenValues()

	// First lets see if it is Zero or any of the Even Values, if it is, then assign it "Zero"
	if helper.FindIfStringInSlice(strings.ToUpper(stringValue), juiceEvenValues) != "False" ||
		helper.FindIfStringInSlice(strings.ToUpper(stringValue), spreadEvenValues) != "False" {

		returnString = "Zero"
		isZeroAnError = "False"
		tempFloat = 0
	} else {
		tempFloat, _ = strconv.ParseFloat(stringValue, 32)
		isZeroAnError = "True"
	}

	// If the tempFloat is 0 and the Flag is "True", then an error happened when converting, log error if true
	if tempFloat == 0 && isZeroAnError == "True" {
		returnString = "Error"
		isZeroAnError = "True"
	} else {
		isZeroAnError = "False"
	}

	if isZeroAnError == "False" {
		if tempFloat > 0 {
			returnString = "Positive"
		} else if tempFloat < 0 {
			returnString = "Negative"
		} else if tempFloat == 0 {
			returnString = "Zero"
		}
	}
	return returnString
}

// This function intakes the Master and user, and a string to format, and returns a slice of string that is formatted
// The parameters to replace are {SportID}, {LeagueID}, {PeriodID}
// ex. Intake string can look like this: `"{IdSport":{LeagueID},"Period":{PeriodID}` from Smash66
func (helper Helper) FormatStringSportLeaguePeriod(master Master, slave Slave, stringToFormat string) []string {

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
			} else if master.MasterLines[i].League == "College Football" {
				// Add in the LeagueID, which we know is "CollegeFootball" on this conditional
				tempString = helper.ReplaceParameters(tempString, "{LeagueID}", sportsDict.Football.CollegeFootball.ID)

				// Now we have to add in the Period of the bets (Game, 1st Half, etc)
				if master.MasterLines[i].Period == "Game" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.Game)
				} else if master.MasterLines[i].Period == "1st Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.OneHalf)
				} else if master.MasterLines[i].Period == "2nd Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.TwoHalf)
				} else if master.MasterLines[i].Period == "1st Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.OneQuarter)
				} else if master.MasterLines[i].Period == "2nd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.TwoQuarter)
				} else if master.MasterLines[i].Period == "3rd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.ThreeQuarter)
				} else if master.MasterLines[i].Period == "4th Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Football.CollegeFootball.Period.FourQuarter)
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
			} else if master.MasterLines[i].League == "College Football" {
				// Add in the LeagueID, which we know is "College Basketball" on this conditional
				tempString = helper.ReplaceParameters(tempString, "{LeagueID}", sportsDict.Basketball.CollegeBasketball.ID)

				// Now we have to add in the Period of the bets (Game, 1st Half, etc)
				if master.MasterLines[i].Period == "Game" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.Game)
				} else if master.MasterLines[i].Period == "1st Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.OneHalf)
				} else if master.MasterLines[i].Period == "2nd Half" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.TwoHalf)
				} else if master.MasterLines[i].Period == "1st Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.OneQuarter)
				} else if master.MasterLines[i].Period == "2nd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.TwoQuarter)
				} else if master.MasterLines[i].Period == "3rd Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.ThreeQuarter)
				} else if master.MasterLines[i].Period == "4th Quarter" {
					tempString = helper.ReplaceParameters(tempString, "{PeriodID}", sportsDict.Basketball.CollegeBasketball.Period.FourQuarter)
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

// This function helps get a value from a Period
func (helper Helper) GetPeriodTextValue(dictionary SiteDictionary, periodValue string) string {

	returnPeriodText := ""

	// These are all the "Game" Period values
	if periodValue == dictionary.SportsDict.Baseball.MLB.Period.Game ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.Game ||
		periodValue == dictionary.SportsDict.Basketball.NBA.Period.Game ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.Game ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.Game {
		returnPeriodText = "Game"

	} else if periodValue == dictionary.SportsDict.Baseball.MLB.Period.OneFiveInnings {
		returnPeriodText = "1st 5 Innings"

	} else if periodValue == dictionary.SportsDict.Basketball.NBA.Period.OneHalf ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.OneHalf ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.OneHalf ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.OneHalf {
		returnPeriodText = "1st Half"

	} else if periodValue == dictionary.SportsDict.Basketball.NBA.Period.TwoHalf ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.TwoHalf ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.TwoHalf ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.TwoHalf {
		returnPeriodText = "2nd Half"

	} else if periodValue == dictionary.SportsDict.Basketball.NBA.Period.OneQuarter ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.OneQuarter ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.OneQuarter ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.OneQuarter {
		returnPeriodText = "1st Quarter"

	} else if periodValue == dictionary.SportsDict.Basketball.NBA.Period.TwoQuarter ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.TwoQuarter ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.TwoQuarter ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.TwoQuarter {
		returnPeriodText = "2nd Quarter"

	} else if periodValue == dictionary.SportsDict.Basketball.NBA.Period.ThreeQuarter ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.ThreeQuarter ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.ThreeQuarter ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.ThreeQuarter {
		returnPeriodText = "3rd Quarter"

	} else if periodValue == dictionary.SportsDict.Basketball.NBA.Period.FourQuarter ||
		periodValue == dictionary.SportsDict.Basketball.CollegeBasketball.Period.FourQuarter ||
		periodValue == dictionary.SportsDict.Football.NFL.Period.FourQuarter ||
		periodValue == dictionary.SportsDict.Football.CollegeFootball.Period.FourQuarter {
		returnPeriodText = "4th Quarter"

	} else {
		returnPeriodText = "Undefined"
	}

	return returnPeriodText
}

// Helper Function to help print JSON formatted of a struct
func (helper Helper) PrintStructInJSON(v interface{}) {
	printString, _ := json.MarshalIndent(v, "", "    ")
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, printString, "", "    ")
	if err != nil {
		fmt.Println("JSON parse error: ", err)
	}

	fmt.Println("---------------------------------------")
	fmt.Println()
	fmt.Println(string(prettyJSON.Bytes()))
	fmt.Println()
	fmt.Println("---------------------------------------")
}
