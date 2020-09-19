package anton

import (
	"fmt"
	"reflect"
)

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
