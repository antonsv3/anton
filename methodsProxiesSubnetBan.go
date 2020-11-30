package anton

// Proxy Method to take in slice of Subnets, and either return "False", "Banned", "Error" string
func (proxy *Proxy) SeeIfProxiesBanned(allSubnets []Subnet) string {

	returnString := ""

	var helper Helper

	var matchedSubnet Subnet

	var foundMatchedSubnetFlag = "False"

	// Loop through all Subnets and find the matching one
	for i := range allSubnets {
		if proxy.Subnet == allSubnets[i].Subnet {
			foundMatchedSubnetFlag = "True"
			matchedSubnet = allSubnets[i]
			break
		}
	}

	// If we found the Subnet, loop through the Subnet banned list to see if the SiteName is in there
	if foundMatchedSubnetFlag == "True" {

		// Only iterate the slice if there is something in there, to prevent errors
		if len(matchedSubnet.BannedSites) > 0 {
			if helper.FindIfStringInSlice(proxy.SiteName, matchedSubnet.BannedSites) != "False" {
				returnString = "Banned"
			} else {
				returnString = "False"
			}
		} else {
			returnString = "False"
		}

	} else {
		returnString = "Error"
	}

	return returnString

}
