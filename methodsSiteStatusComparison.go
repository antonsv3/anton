package anton

// Function to compare two SiteStatus and return one
func CompareSiteStatuses(MasterSiteStatus, SlaveSiteStatus SiteStatus) SiteStatus {

	var returnSiteStatus SiteStatus

	// I only care about the Master Site Status for the Master Agent Section

	// --------------- Master Agent Values - MLB

	if MasterSiteStatus.Master.Agent.MLB.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.MLB.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.Agent.MLB.Spread != "Yes" {
		returnSiteStatus.Master.Agent.MLB.Spread = "No"
	}

	if MasterSiteStatus.Master.Agent.MLB.Total != "Yes" {
		returnSiteStatus.Master.Agent.MLB.Total = "No"
	}

	if MasterSiteStatus.Master.Agent.MLB.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.MLB.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.Agent.MLB.Game != "Yes" {
		returnSiteStatus.Master.Agent.MLB.Game = "No"
	}

	if MasterSiteStatus.Master.Agent.MLB.OneFiveInnings != "Yes" {
		returnSiteStatus.Master.Agent.MLB.OneFiveInnings = "No"
	}

	// --------------- Master Agent Values - NBA

	if MasterSiteStatus.Master.Agent.NBA.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.NBA.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.Spread != "Yes" {
		returnSiteStatus.Master.Agent.NBA.Spread = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.Total != "Yes" {
		returnSiteStatus.Master.Agent.NBA.Total = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.NBA.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.Game != "Yes" {
		returnSiteStatus.Master.Agent.NBA.Game = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.NBA.OneHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.NBA.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.NBA.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.FourQuarter = "No"
	}

	// --------------- Master Agent Values - NFL

	if MasterSiteStatus.Master.Agent.NFL.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.NFL.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.Spread != "Yes" {
		returnSiteStatus.Master.Agent.NFL.Spread = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.Total != "Yes" {
		returnSiteStatus.Master.Agent.NFL.Total = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.NFL.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.Game != "Yes" {
		returnSiteStatus.Master.Agent.NFL.Game = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.NFL.OneHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.NFL.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.NFL.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.FourQuarter = "No"
	}

	// --------------- Master Agent Values - College Basketball

	if MasterSiteStatus.Master.Agent.CollegeBasketball.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.Spread != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.Spread = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.Total != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.Total = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.Game != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.Game = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.OneHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeBasketball.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.FourQuarter = "No"
	}

	// --------------- Master Agent Values - College Football

	if MasterSiteStatus.Master.Agent.CollegeFootball.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.Spread != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.Spread = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.Total != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.Total = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.Game != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.Game = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.OneHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.Agent.CollegeFootball.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.FourQuarter = "No"
	}

	// I only care about the Master Site Status for the Master User Section

	// --------------- Master User Values - MLB

	if MasterSiteStatus.Master.User.MLB.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.MLB.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.User.MLB.Spread != "Yes" {
		returnSiteStatus.Master.User.MLB.Spread = "No"
	}

	if MasterSiteStatus.Master.User.MLB.Total != "Yes" {
		returnSiteStatus.Master.User.MLB.Total = "No"
	}

	if MasterSiteStatus.Master.User.MLB.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.MLB.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.User.MLB.Game != "Yes" {
		returnSiteStatus.Master.User.MLB.Game = "No"
	}

	if MasterSiteStatus.Master.User.MLB.OneFiveInnings != "Yes" {
		returnSiteStatus.Master.User.MLB.OneFiveInnings = "No"
	}

	// --------------- Master User Values - NBA

	if MasterSiteStatus.Master.User.NBA.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.NBA.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.User.NBA.Spread != "Yes" {
		returnSiteStatus.Master.User.NBA.Spread = "No"
	}

	if MasterSiteStatus.Master.User.NBA.Total != "Yes" {
		returnSiteStatus.Master.User.NBA.Total = "No"
	}

	if MasterSiteStatus.Master.User.NBA.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.NBA.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.User.NBA.Game != "Yes" {
		returnSiteStatus.Master.User.NBA.Game = "No"
	}

	if MasterSiteStatus.Master.User.NBA.OneHalf != "Yes" {
		returnSiteStatus.Master.User.NBA.OneHalf = "No"
	}

	if MasterSiteStatus.Master.User.NBA.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.NBA.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.User.NBA.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.User.NBA.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.User.NBA.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.User.NBA.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.FourQuarter = "No"
	}

	// --------------- Master User Values - NFL

	if MasterSiteStatus.Master.User.NFL.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.NFL.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.User.NFL.Spread != "Yes" {
		returnSiteStatus.Master.User.NFL.Spread = "No"
	}

	if MasterSiteStatus.Master.User.NFL.Total != "Yes" {
		returnSiteStatus.Master.User.NFL.Total = "No"
	}

	if MasterSiteStatus.Master.User.NFL.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.NFL.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.User.NFL.Game != "Yes" {
		returnSiteStatus.Master.User.NFL.Game = "No"
	}

	if MasterSiteStatus.Master.User.NFL.OneHalf != "Yes" {
		returnSiteStatus.Master.User.NFL.OneHalf = "No"
	}

	if MasterSiteStatus.Master.User.NFL.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.NFL.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.User.NFL.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.User.NFL.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.User.NFL.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.User.NFL.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.FourQuarter = "No"
	}

	// --------------- Master User Values - College Basketball

	if MasterSiteStatus.Master.User.CollegeBasketball.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.Spread != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.Spread = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.Total != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.Total = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.Game != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.Game = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.OneHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.OneHalf = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.User.CollegeBasketball.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.FourQuarter = "No"
	}

	// --------------- Master User Values - College Football

	if MasterSiteStatus.Master.User.CollegeFootball.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.MoneyLine = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.Spread != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.Spread = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.Total != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.Total = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.TeamTotal = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.Game != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.Game = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.OneHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.OneHalf = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.TwoHalf = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.OneQuarter = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.TwoQuarter = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.ThreeQuarter = "No"
	}

	if MasterSiteStatus.Master.User.CollegeFootball.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.FourQuarter = "No"
	}

	// I only care about the Slave Site Status for the Slave Section

	// --------------- Slave Values - MLB

	if SlaveSiteStatus.Slave.MLB.MoneyLine != "Yes" {
		returnSiteStatus.Slave.MLB.MoneyLine = "No"
	}

	if SlaveSiteStatus.Slave.MLB.Spread != "Yes" {
		returnSiteStatus.Slave.MLB.Spread = "No"
	}

	if SlaveSiteStatus.Slave.MLB.Total != "Yes" {
		returnSiteStatus.Slave.MLB.Total = "No"
	}

	if SlaveSiteStatus.Slave.MLB.TeamTotal != "Yes" {
		returnSiteStatus.Slave.MLB.TeamTotal = "No"
	}

	if SlaveSiteStatus.Slave.MLB.Game != "Yes" {
		returnSiteStatus.Slave.MLB.Game = "No"
	}

	if SlaveSiteStatus.Slave.MLB.OneFiveInnings != "Yes" {
		returnSiteStatus.Slave.MLB.OneFiveInnings = "No"
	}

	// --------------- Slave Values - NBA

	if SlaveSiteStatus.Slave.NBA.MoneyLine != "Yes" {
		returnSiteStatus.Slave.NBA.MoneyLine = "No"
	}

	if SlaveSiteStatus.Slave.NBA.Spread != "Yes" {
		returnSiteStatus.Slave.NBA.Spread = "No"
	}

	if SlaveSiteStatus.Slave.NBA.Total != "Yes" {
		returnSiteStatus.Slave.NBA.Total = "No"
	}

	if SlaveSiteStatus.Slave.NBA.TeamTotal != "Yes" {
		returnSiteStatus.Slave.NBA.TeamTotal = "No"
	}

	if SlaveSiteStatus.Slave.NBA.Game != "Yes" {
		returnSiteStatus.Slave.NBA.Game = "No"
	}

	if SlaveSiteStatus.Slave.NBA.OneHalf != "Yes" {
		returnSiteStatus.Slave.NBA.OneHalf = "No"
	}

	if SlaveSiteStatus.Slave.NBA.TwoHalf != "Yes" {
		returnSiteStatus.Slave.NBA.TwoHalf = "No"
	}

	if SlaveSiteStatus.Slave.NBA.OneQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.OneQuarter = "No"
	}

	if SlaveSiteStatus.Slave.NBA.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.TwoQuarter = "No"
	}

	if SlaveSiteStatus.Slave.NBA.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.ThreeQuarter = "No"
	}

	if SlaveSiteStatus.Slave.NBA.FourQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.FourQuarter = "No"
	}

	// --------------- Slave Values - NFL

	if SlaveSiteStatus.Slave.NFL.MoneyLine != "Yes" {
		returnSiteStatus.Slave.NFL.MoneyLine = "No"
	}

	if SlaveSiteStatus.Slave.NFL.Spread != "Yes" {
		returnSiteStatus.Slave.NFL.Spread = "No"
	}

	if SlaveSiteStatus.Slave.NFL.Total != "Yes" {
		returnSiteStatus.Slave.NFL.Total = "No"
	}

	if SlaveSiteStatus.Slave.NFL.TeamTotal != "Yes" {
		returnSiteStatus.Slave.NFL.TeamTotal = "No"
	}

	if SlaveSiteStatus.Slave.NFL.Game != "Yes" {
		returnSiteStatus.Slave.NFL.Game = "No"
	}

	if SlaveSiteStatus.Slave.NFL.OneHalf != "Yes" {
		returnSiteStatus.Slave.NFL.OneHalf = "No"
	}

	if SlaveSiteStatus.Slave.NFL.TwoHalf != "Yes" {
		returnSiteStatus.Slave.NFL.TwoHalf = "No"
	}

	if SlaveSiteStatus.Slave.NFL.OneQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.OneQuarter = "No"
	}

	if SlaveSiteStatus.Slave.NFL.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.TwoQuarter = "No"
	}

	if SlaveSiteStatus.Slave.NFL.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.ThreeQuarter = "No"
	}

	if SlaveSiteStatus.Slave.NFL.FourQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.FourQuarter = "No"
	}

	// --------------- Slave Values - College Basketball

	if SlaveSiteStatus.Slave.CollegeBasketball.MoneyLine != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.MoneyLine = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.Spread != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.Spread = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.Total != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.Total = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.TeamTotal != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.TeamTotal = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.Game != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.Game = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.OneHalf != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.OneHalf = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.TwoHalf != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.TwoHalf = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.OneQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.OneQuarter = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.TwoQuarter = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.ThreeQuarter = "No"
	}

	if SlaveSiteStatus.Slave.CollegeBasketball.FourQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.FourQuarter = "No"
	}

	// --------------- Slave Values - College Football

	if SlaveSiteStatus.Slave.CollegeFootball.MoneyLine != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.MoneyLine = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.Spread != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.Spread = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.Total != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.Total = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.TeamTotal != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.TeamTotal = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.Game != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.Game = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.OneHalf != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.OneHalf = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.TwoHalf != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.TwoHalf = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.OneQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.OneQuarter = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.TwoQuarter = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.ThreeQuarter = "No"
	}

	if SlaveSiteStatus.Slave.CollegeFootball.FourQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.FourQuarter = "No"
	}

	return returnSiteStatus

}
