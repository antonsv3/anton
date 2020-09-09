package anton

// MASTER DONE

// These are the functions that help us ingest values, will be used to format when Creating Lines in the Methods
// Used in both methodsMaster.go and methodsSlave.go

func formatRotationNumber(rotationNumber string) string {
	// Declare the helper struct to access the helper functions
	var helper Helper
	returnRotationNumber := helper.ReplaceParameters(rotationNumber, "\u00a0", "", " ", "")
	return returnRotationNumber
}

func formatLineSpread(lineSpread string) string {
	// Declare the helper struct to access the helper functions
	var helper Helper
	returnLineSpread := helper.ReplaceParameters(lineSpread, "½", ".5", "\u00a0", "", " ", "")
	return returnLineSpread
}

func formatLineJuice(lineJuice string) string {
	// Declare the helper struct to access the helper functions
	var helper Helper
	returnLineJuice := helper.ReplaceParameters(lineJuice, "½", ".5", "\u00a0", "", " ", "")
	return returnLineJuice
}

func formatPeriod(period string) string {
	// Declare the helper struct to access the helper functions
	var helper Helper
	returnPeriod := helper.ReplaceParameters(period, "\u00a0", "", " ", "")

	return returnPeriod
}

func formatTeam(team string) string {
	returnTeam := team
	return returnTeam
}

/*








 */
