package anton

import "fmt"

// Function to compare two SiteStatus and return one
func CompareSiteStatuses(SiteStatusOne, SiteStatusTwo SiteStatus) SiteStatus {

	var returnSiteStatus SiteStatus

	// --------------- Master Agent Values - MLB

	if SiteStatusOne.Master.Agent.MLB.MoneyLine != "Yes" || SiteStatusTwo.Master.Agent.MLB.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.MLB.MoneyLine = "No"
	}

	if SiteStatusOne.Master.Agent.MLB.Spread != "Yes" || SiteStatusTwo.Master.Agent.MLB.Spread != "Yes" {
		returnSiteStatus.Master.Agent.MLB.Spread = "No"
	}

	if SiteStatusOne.Master.Agent.MLB.Total != "Yes" || SiteStatusTwo.Master.Agent.MLB.Total != "Yes" {
		returnSiteStatus.Master.Agent.MLB.Total = "No"
	}

	if SiteStatusOne.Master.Agent.MLB.TeamTotal != "Yes" || SiteStatusTwo.Master.Agent.MLB.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.MLB.TeamTotal = "No"
	}

	if SiteStatusOne.Master.Agent.MLB.Game != "Yes" || SiteStatusTwo.Master.Agent.MLB.Game != "Yes" {
		returnSiteStatus.Master.Agent.MLB.Game = "No"
	}

	if SiteStatusOne.Master.Agent.MLB.OneFiveInnings != "Yes" || SiteStatusTwo.Master.Agent.MLB.OneFiveInnings != "Yes" {
		returnSiteStatus.Master.Agent.MLB.OneFiveInnings = "No"
	}

	// --------------- Master Agent Values - NBA

	if SiteStatusOne.Master.Agent.NBA.MoneyLine != "Yes" || SiteStatusTwo.Master.Agent.NBA.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.NBA.MoneyLine = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.Spread != "Yes" || SiteStatusTwo.Master.Agent.NBA.Spread != "Yes" {
		returnSiteStatus.Master.Agent.NBA.Spread = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.Total != "Yes" || SiteStatusTwo.Master.Agent.NBA.Total != "Yes" {
		returnSiteStatus.Master.Agent.NBA.Total = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.TeamTotal != "Yes" || SiteStatusTwo.Master.Agent.NBA.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.NBA.TeamTotal = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.Game != "Yes" || SiteStatusTwo.Master.Agent.NBA.Game != "Yes" {
		returnSiteStatus.Master.Agent.NBA.Game = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.OneHalf != "Yes" || SiteStatusTwo.Master.Agent.NBA.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.NBA.OneHalf = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.TwoHalf != "Yes" || SiteStatusTwo.Master.Agent.NBA.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.NBA.TwoHalf = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.OneQuarter != "Yes" || SiteStatusTwo.Master.Agent.NBA.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.OneQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.TwoQuarter != "Yes" || SiteStatusTwo.Master.Agent.NBA.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.ThreeQuarter != "Yes" || SiteStatusTwo.Master.Agent.NBA.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.NBA.FourQuarter != "Yes" || SiteStatusTwo.Master.Agent.NBA.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NBA.FourQuarter = "No"
	}

	// --------------- Master Agent Values - NFL

	if SiteStatusOne.Master.Agent.NFL.MoneyLine != "Yes" || SiteStatusTwo.Master.Agent.NFL.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.NFL.MoneyLine = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.Spread != "Yes" || SiteStatusTwo.Master.Agent.NFL.Spread != "Yes" {
		returnSiteStatus.Master.Agent.NFL.Spread = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.Total != "Yes" || SiteStatusTwo.Master.Agent.NFL.Total != "Yes" {
		returnSiteStatus.Master.Agent.NFL.Total = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.TeamTotal != "Yes" || SiteStatusTwo.Master.Agent.NFL.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.NFL.TeamTotal = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.Game != "Yes" || SiteStatusTwo.Master.Agent.NFL.Game != "Yes" {
		returnSiteStatus.Master.Agent.NFL.Game = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.OneHalf != "Yes" || SiteStatusTwo.Master.Agent.NFL.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.NFL.OneHalf = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.TwoHalf != "Yes" || SiteStatusTwo.Master.Agent.NFL.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.NFL.TwoHalf = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.OneQuarter != "Yes" || SiteStatusTwo.Master.Agent.NFL.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.OneQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.TwoQuarter != "Yes" || SiteStatusTwo.Master.Agent.NFL.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.ThreeQuarter != "Yes" || SiteStatusTwo.Master.Agent.NFL.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.NFL.FourQuarter != "Yes" || SiteStatusTwo.Master.Agent.NFL.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.NFL.FourQuarter = "No"
	}

	// --------------- Master Agent Values - College Basketball

	if SiteStatusOne.Master.Agent.CollegeBasketball.MoneyLine != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.MoneyLine = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.Spread != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.Spread != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.Spread = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.Total != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.Total != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.Total = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.TeamTotal != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.TeamTotal = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.Game != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.Game != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.Game = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.OneHalf != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.OneHalf = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.TwoHalf != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.TwoHalf = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.OneQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.OneQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.TwoQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.ThreeQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeBasketball.FourQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeBasketball.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeBasketball.FourQuarter = "No"
	}

	// --------------- Master Agent Values - College Football

	if SiteStatusOne.Master.Agent.CollegeFootball.MoneyLine != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.MoneyLine != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.MoneyLine = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.Spread != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.Spread != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.Spread = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.Total != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.Total != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.Total = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.TeamTotal != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.TeamTotal != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.TeamTotal = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.Game != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.Game != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.Game = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.OneHalf != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.OneHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.OneHalf = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.TwoHalf != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.TwoHalf != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.TwoHalf = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.OneQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.OneQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.OneQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.TwoQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.ThreeQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.Agent.CollegeFootball.FourQuarter != "Yes" || SiteStatusTwo.Master.Agent.CollegeFootball.FourQuarter != "Yes" {
		returnSiteStatus.Master.Agent.CollegeFootball.FourQuarter = "No"
	}

	// --------------- Master User Values - MLB

	if SiteStatusOne.Master.User.MLB.MoneyLine != "Yes" || SiteStatusTwo.Master.User.MLB.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.MLB.MoneyLine = "No"
	}

	if SiteStatusOne.Master.User.MLB.Spread != "Yes" || SiteStatusTwo.Master.User.MLB.Spread != "Yes" {
		returnSiteStatus.Master.User.MLB.Spread = "No"
	}

	if SiteStatusOne.Master.User.MLB.Total != "Yes" || SiteStatusTwo.Master.User.MLB.Total != "Yes" {
		returnSiteStatus.Master.User.MLB.Total = "No"
	}

	if SiteStatusOne.Master.User.MLB.TeamTotal != "Yes" || SiteStatusTwo.Master.User.MLB.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.MLB.TeamTotal = "No"
	}

	if SiteStatusOne.Master.User.MLB.Game != "Yes" || SiteStatusTwo.Master.User.MLB.Game != "Yes" {
		returnSiteStatus.Master.User.MLB.Game = "No"
	}

	if SiteStatusOne.Master.User.MLB.OneFiveInnings != "Yes" || SiteStatusTwo.Master.User.MLB.OneFiveInnings != "Yes" {
		returnSiteStatus.Master.User.MLB.OneFiveInnings = "No"
	}

	// --------------- Master User Values - NBA

	if SiteStatusOne.Master.User.NBA.MoneyLine != "Yes" || SiteStatusTwo.Master.User.NBA.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.NBA.MoneyLine = "No"
	}

	if SiteStatusOne.Master.User.NBA.Spread != "Yes" || SiteStatusTwo.Master.User.NBA.Spread != "Yes" {
		returnSiteStatus.Master.User.NBA.Spread = "No"
	}

	if SiteStatusOne.Master.User.NBA.Total != "Yes" || SiteStatusTwo.Master.User.NBA.Total != "Yes" {
		returnSiteStatus.Master.User.NBA.Total = "No"
	}

	if SiteStatusOne.Master.User.NBA.TeamTotal != "Yes" || SiteStatusTwo.Master.User.NBA.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.NBA.TeamTotal = "No"
	}

	if SiteStatusOne.Master.User.NBA.Game != "Yes" || SiteStatusTwo.Master.User.NBA.Game != "Yes" {
		returnSiteStatus.Master.User.NBA.Game = "No"
	}

	if SiteStatusOne.Master.User.NBA.OneHalf != "Yes" || SiteStatusTwo.Master.User.NBA.OneHalf != "Yes" {
		returnSiteStatus.Master.User.NBA.OneHalf = "No"
	}

	if SiteStatusOne.Master.User.NBA.TwoHalf != "Yes" || SiteStatusTwo.Master.User.NBA.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.NBA.TwoHalf = "No"
	}

	if SiteStatusOne.Master.User.NBA.OneQuarter != "Yes" || SiteStatusTwo.Master.User.NBA.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.OneQuarter = "No"
	}

	if SiteStatusOne.Master.User.NBA.TwoQuarter != "Yes" || SiteStatusTwo.Master.User.NBA.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.User.NBA.ThreeQuarter != "Yes" || SiteStatusTwo.Master.User.NBA.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.User.NBA.FourQuarter != "Yes" || SiteStatusTwo.Master.User.NBA.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.NBA.FourQuarter = "No"
	}

	// --------------- Master User Values - NFL

	if SiteStatusOne.Master.User.NFL.MoneyLine != "Yes" || SiteStatusTwo.Master.User.NFL.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.NFL.MoneyLine = "No"
	}

	if SiteStatusOne.Master.User.NFL.Spread != "Yes" || SiteStatusTwo.Master.User.NFL.Spread != "Yes" {
		returnSiteStatus.Master.User.NFL.Spread = "No"
	}

	if SiteStatusOne.Master.User.NFL.Total != "Yes" || SiteStatusTwo.Master.User.NFL.Total != "Yes" {
		returnSiteStatus.Master.User.NFL.Total = "No"
	}

	if SiteStatusOne.Master.User.NFL.TeamTotal != "Yes" || SiteStatusTwo.Master.User.NFL.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.NFL.TeamTotal = "No"
	}

	if SiteStatusOne.Master.User.NFL.Game != "Yes" || SiteStatusTwo.Master.User.NFL.Game != "Yes" {
		returnSiteStatus.Master.User.NFL.Game = "No"
	}

	if SiteStatusOne.Master.User.NFL.OneHalf != "Yes" || SiteStatusTwo.Master.User.NFL.OneHalf != "Yes" {
		returnSiteStatus.Master.User.NFL.OneHalf = "No"
	}

	if SiteStatusOne.Master.User.NFL.TwoHalf != "Yes" || SiteStatusTwo.Master.User.NFL.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.NFL.TwoHalf = "No"
	}

	if SiteStatusOne.Master.User.NFL.OneQuarter != "Yes" || SiteStatusTwo.Master.User.NFL.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.OneQuarter = "No"
	}

	if SiteStatusOne.Master.User.NFL.TwoQuarter != "Yes" || SiteStatusTwo.Master.User.NFL.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.User.NFL.ThreeQuarter != "Yes" || SiteStatusTwo.Master.User.NFL.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.User.NFL.FourQuarter != "Yes" || SiteStatusTwo.Master.User.NFL.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.NFL.FourQuarter = "No"
	}

	// --------------- Master User Values - College Basketball

	if SiteStatusOne.Master.User.CollegeBasketball.MoneyLine != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.MoneyLine = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.Spread != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.Spread != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.Spread = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.Total != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.Total != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.Total = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.TeamTotal != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.TeamTotal = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.Game != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.Game != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.Game = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.OneHalf != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.OneHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.OneHalf = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.TwoHalf != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.TwoHalf = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.OneQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.OneQuarter = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.TwoQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.ThreeQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.User.CollegeBasketball.FourQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeBasketball.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeBasketball.FourQuarter = "No"
	}

	// --------------- Master User Values - College Football

	if SiteStatusOne.Master.User.CollegeFootball.MoneyLine != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.MoneyLine != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.MoneyLine = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.Spread != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.Spread != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.Spread = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.Total != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.Total != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.Total = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.TeamTotal != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.TeamTotal != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.TeamTotal = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.Game != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.Game != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.Game = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.OneHalf != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.OneHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.OneHalf = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.TwoHalf != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.TwoHalf != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.TwoHalf = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.OneQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.OneQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.OneQuarter = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.TwoQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.TwoQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.TwoQuarter = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.ThreeQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.ThreeQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.ThreeQuarter = "No"
	}

	if SiteStatusOne.Master.User.CollegeFootball.FourQuarter != "Yes" || SiteStatusTwo.Master.User.CollegeFootball.FourQuarter != "Yes" {
		returnSiteStatus.Master.User.CollegeFootball.FourQuarter = "No"
	}

	// --------------- Slave Values - MLB

	if SiteStatusOne.Slave.MLB.MoneyLine != "Yes" || SiteStatusTwo.Slave.MLB.MoneyLine != "Yes" {
		returnSiteStatus.Slave.MLB.MoneyLine = "No"
	}

	if SiteStatusOne.Slave.MLB.Spread != "Yes" || SiteStatusTwo.Slave.MLB.Spread != "Yes" {
		returnSiteStatus.Slave.MLB.Spread = "No"
	}

	if SiteStatusOne.Slave.MLB.Total != "Yes" || SiteStatusTwo.Slave.MLB.Total != "Yes" {
		returnSiteStatus.Slave.MLB.Total = "No"
	}

	if SiteStatusOne.Slave.MLB.TeamTotal != "Yes" || SiteStatusTwo.Slave.MLB.TeamTotal != "Yes" {
		returnSiteStatus.Slave.MLB.TeamTotal = "No"
	}

	if SiteStatusOne.Slave.MLB.Game != "Yes" || SiteStatusTwo.Slave.MLB.Game != "Yes" {
		returnSiteStatus.Slave.MLB.Game = "No"
	}

	if SiteStatusOne.Slave.MLB.OneFiveInnings != "Yes" || SiteStatusTwo.Slave.MLB.OneFiveInnings != "Yes" {
		returnSiteStatus.Slave.MLB.OneFiveInnings = "No"
	}

	// --------------- Slave Values - NBA

	if SiteStatusOne.Slave.NBA.MoneyLine != "Yes" || SiteStatusTwo.Slave.NBA.MoneyLine != "Yes" {
		returnSiteStatus.Slave.NBA.MoneyLine = "No"
	}

	if SiteStatusOne.Slave.NBA.Spread != "Yes" || SiteStatusTwo.Slave.NBA.Spread != "Yes" {
		returnSiteStatus.Slave.NBA.Spread = "No"
	}

	if SiteStatusOne.Slave.NBA.Total != "Yes" || SiteStatusTwo.Slave.NBA.Total != "Yes" {
		returnSiteStatus.Slave.NBA.Total = "No"
	}

	if SiteStatusOne.Slave.NBA.TeamTotal != "Yes" || SiteStatusTwo.Slave.NBA.TeamTotal != "Yes" {
		returnSiteStatus.Slave.NBA.TeamTotal = "No"
	}

	if SiteStatusOne.Slave.NBA.Game != "Yes" || SiteStatusTwo.Slave.NBA.Game != "Yes" {
		returnSiteStatus.Slave.NBA.Game = "No"
	}

	if SiteStatusOne.Slave.NBA.OneHalf != "Yes" || SiteStatusTwo.Slave.NBA.OneHalf != "Yes" {
		returnSiteStatus.Slave.NBA.OneHalf = "No"
	}

	if SiteStatusOne.Slave.NBA.TwoHalf != "Yes" || SiteStatusTwo.Slave.NBA.TwoHalf != "Yes" {
		returnSiteStatus.Slave.NBA.TwoHalf = "No"
	}

	if SiteStatusOne.Slave.NBA.OneQuarter != "Yes" || SiteStatusTwo.Slave.NBA.OneQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.OneQuarter = "No"
	}

	if SiteStatusOne.Slave.NBA.TwoQuarter != "Yes" || SiteStatusTwo.Slave.NBA.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.TwoQuarter = "No"
	}

	if SiteStatusOne.Slave.NBA.ThreeQuarter != "Yes" || SiteStatusTwo.Slave.NBA.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.ThreeQuarter = "No"
	}

	if SiteStatusOne.Slave.NBA.FourQuarter != "Yes" || SiteStatusTwo.Slave.NBA.FourQuarter != "Yes" {
		returnSiteStatus.Slave.NBA.FourQuarter = "No"
	}

	// --------------- Slave Values - NFL

	if SiteStatusOne.Slave.NFL.MoneyLine != "Yes" || SiteStatusTwo.Slave.NFL.MoneyLine != "Yes" {
		returnSiteStatus.Slave.NFL.MoneyLine = "No"
	}

	if SiteStatusOne.Slave.NFL.Spread != "Yes" || SiteStatusTwo.Slave.NFL.Spread != "Yes" {
		returnSiteStatus.Slave.NFL.Spread = "No"
	}

	if SiteStatusOne.Slave.NFL.Total != "Yes" || SiteStatusTwo.Slave.NFL.Total != "Yes" {
		returnSiteStatus.Slave.NFL.Total = "No"
	}

	if SiteStatusOne.Slave.NFL.TeamTotal != "Yes" || SiteStatusTwo.Slave.NFL.TeamTotal != "Yes" {
		returnSiteStatus.Slave.NFL.TeamTotal = "No"
	}

	if SiteStatusOne.Slave.NFL.Game != "Yes" || SiteStatusTwo.Slave.NFL.Game != "Yes" {
		returnSiteStatus.Slave.NFL.Game = "No"
	}

	if SiteStatusOne.Slave.NFL.OneHalf != "Yes" || SiteStatusTwo.Slave.NFL.OneHalf != "Yes" {
		returnSiteStatus.Slave.NFL.OneHalf = "No"
	}

	if SiteStatusOne.Slave.NFL.TwoHalf != "Yes" || SiteStatusTwo.Slave.NFL.TwoHalf != "Yes" {
		returnSiteStatus.Slave.NFL.TwoHalf = "No"
	}

	if SiteStatusOne.Slave.NFL.OneQuarter != "Yes" || SiteStatusTwo.Slave.NFL.OneQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.OneQuarter = "No"
	}

	if SiteStatusOne.Slave.NFL.TwoQuarter != "Yes" || SiteStatusTwo.Slave.NFL.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.TwoQuarter = "No"
	}

	if SiteStatusOne.Slave.NFL.ThreeQuarter != "Yes" || SiteStatusTwo.Slave.NFL.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.ThreeQuarter = "No"
	}

	if SiteStatusOne.Slave.NFL.FourQuarter != "Yes" || SiteStatusTwo.Slave.NFL.FourQuarter != "Yes" {
		returnSiteStatus.Slave.NFL.FourQuarter = "No"
	}

	// --------------- Slave Values - College Basketball

	if SiteStatusOne.Slave.CollegeBasketball.MoneyLine != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.MoneyLine != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.MoneyLine = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.Spread != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.Spread != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.Spread = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.Total != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.Total != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.Total = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.TeamTotal != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.TeamTotal != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.TeamTotal = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.Game != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.Game != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.Game = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.OneHalf != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.OneHalf != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.OneHalf = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.TwoHalf != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.TwoHalf != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.TwoHalf = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.OneQuarter != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.OneQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.OneQuarter = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.TwoQuarter != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.TwoQuarter = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.ThreeQuarter != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.ThreeQuarter = "No"
	}

	if SiteStatusOne.Slave.CollegeBasketball.FourQuarter != "Yes" || SiteStatusTwo.Slave.CollegeBasketball.FourQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeBasketball.FourQuarter = "No"
	}

	// --------------- Slave Values - College Football

	if SiteStatusOne.Slave.CollegeFootball.MoneyLine != "Yes" || SiteStatusTwo.Slave.CollegeFootball.MoneyLine != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.MoneyLine = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.Spread != "Yes" || SiteStatusTwo.Slave.CollegeFootball.Spread != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.Spread = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.Total != "Yes" || SiteStatusTwo.Slave.CollegeFootball.Total != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.Total = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.TeamTotal != "Yes" || SiteStatusTwo.Slave.CollegeFootball.TeamTotal != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.TeamTotal = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.Game != "Yes" || SiteStatusTwo.Slave.CollegeFootball.Game != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.Game = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.OneHalf != "Yes" || SiteStatusTwo.Slave.CollegeFootball.OneHalf != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.OneHalf = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.TwoHalf != "Yes" || SiteStatusTwo.Slave.CollegeFootball.TwoHalf != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.TwoHalf = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.OneQuarter != "Yes" || SiteStatusTwo.Slave.CollegeFootball.OneQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.OneQuarter = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.TwoQuarter != "Yes" || SiteStatusTwo.Slave.CollegeFootball.TwoQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.TwoQuarter = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.ThreeQuarter != "Yes" || SiteStatusTwo.Slave.CollegeFootball.ThreeQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.ThreeQuarter = "No"
	}

	if SiteStatusOne.Slave.CollegeFootball.FourQuarter != "Yes" || SiteStatusTwo.Slave.CollegeFootball.FourQuarter != "Yes" {
		returnSiteStatus.Slave.CollegeFootball.FourQuarter = "No"
	}

	var helper Helper
	fmt.Println("SiteStatusOne:")
	helper.PrintStructInJSON(SiteStatusOne)
	fmt.Println("SiteStatusTwo:")
	helper.PrintStructInJSON(SiteStatusTwo)
	fmt.Println("ReturnSiteStatus:")
	helper.PrintStructInJSON(returnSiteStatus)

	return returnSiteStatus

}
