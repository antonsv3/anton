package anton

// The different Types are:
// Master
// Slave
// Profile
// SiteDictionary
// Lines
// TelegramMsg
// SlaveResults

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create struct to hold info about Masters
type Master struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ObjectID string             `bson:"-"`

	// These are populated when the Master and Lines are received from Scraping Scripts
	MasterName    string `json:"MasterName"`
	MasterPass    string `json:"MasterPass"`
	LoginName     string `json:"LoginName"`
	LoginPass     string `json:"LoginPass"`
	SiteName      string `json:"SiteName"`
	HTTPUserAgent string `json:"HTTPUserAgent"`
	ProxyAddress  string `json:"ProxyAddress"`
	AccountType   string `json:"AccountType"`

	// If the account type is Agent, we will use a different login, and leverage these fields
	AgentPlayers      string   `json:"AgentPlayers" bson:"-"`
	AgentPlayersSlice []string `bson:"-"`

	// Approved Lines sent from the feeder will be in this slice
	MasterLines []Lines `json:"MasterLines" bson:"-"`

	// This slice is populated by pulling the DynamoDB for Slaves to compare against
	Slaves []Slave `bson:"-"`

	// These are properties that will change during runtime, Status will be "Active" or "Error"
	Status        string   `json:"Status"`
	ProgramError  []string `json:"-" bson:"-"`
	LoginAttempts int      `json:"-" bson:"-"`
	LinesOnPage   []string `json:"-" bson:"-"`

	// These will be time stamps for comparisons later
	TimeReceived time.Time `json:"-" bson:"-"`
}

// Create struct to hold info about Slaves pulled from the database
type Slave struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	ObjectID string             `bson:"-"`

	Status      string `bson:"Status" json:"Status"`
	SlaveName   string `json:"SlaveName"`
	SlavePass   string `json:"SlavePass"`
	SiteName    string `json:"SiteName"`
	WagerAmount string `json:"WagerAmount"`

	// Sort out the Lines after Creating and Comparing them
	AuthorizedLines []Lines

	// These are for the UserAgent and the Proxy
	Proxy Proxy `bson:"-" json:"-"`

	// This is populated by pulling from DynamoDB, the structure can be found below
	SiteDictionary SiteDictionary `bson:"-"`
	Profiles       []Profile      `bson:"-"`
}

// Create struct to hold info about lines pulled from scrapped sites
type Lines struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ObjectID string             `bson:"-"`

	// Will need to fill these out based from the scrapper
	RotationNumber  string `json:"RotationNumber"`
	LineSpread      string `json:"LineSpread"`
	LineJuice       string `json:"LineJuice"`
	OverUnder       string `json:"OverUnder"`
	FavoredUnderdog string `json:"FavoredUnderdog"`
	LineType        string `json:"LineType"`
	TicketID        string `json:"TicketID"`
	RiskAmount      string `json:"RiskAmount"`
	ToWinAmount     string `json:"ToWinAmount"`

	// These are more for documenting purposes for the Database
	MasterName string `json:"MasterName"`
	MasterPass string `json:"MasterPass"`
	MasterSite string `json:"MasterSite"`

	SlaveName string `json:"SlaveName"`
	SlavePass string `json:"SlavePass"`
	SlaveSite string `json:"SlaveSite"`

	// This is Optional holder for any unique values for a particular line
	UniqueID string `json:"UniqueID"`
	HomeAway string `json:"HomeAway"`

	// These will be auto-populated by the lines from the scrapper
	Sport  string `json:"Sport"`
	League string `json:"League"`
	Period string `json:"Period"`
	Team   string `json:"Team"`

	LineCharacteristic string `json:"LineCharacteristic"`

	// These can be inherited by Approved Bet, CurrentUser, or have default values for Potential:
	LineStatus         string
	CreatedViaFunction string
	BetType            string
	BettingSite        string
	BettingUser        string
	LineSpreadFloat    float64
	LineJuiceFloat     float64
	FunctionLog        string
	ErrorLog           []string

	ComparedLines []Lines // Line Struct of ApprovedBet			Populated by ApprovedBet
}

type Profile struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ObjectID string             `bson:"-"`

	Status          string `json:"Status"`
	SiteName        string `json:"SiteName"`
	MasterName      string `json:"MasterName"`
	MasterSite      string `json:"MasterSite"`
	SlaveName       string `json:"SlaveName"`
	SlaveSite       string `json:"SlaveSite"`
	SpreadParameter string `json:"SpreadParameter"`
	JuiceParameter  string `json:"JuiceParameter"`
	SportsSettings  struct {
		NFL struct {
			MoneyLine    string `json:"MoneyLine"`
			Spread       string `json:"Spread"`
			Total        string `json:"Total"`
			TeamTotal    string `json:"TeamTotal"`
			Game         string `json:"Game"`
			OneHalf      string `json:"OneHalf"`
			TwoHalf      string `json:"TwoHalf"`
			OneQuarter   string `json:"OneQuarter"`
			TwoQuarter   string `json:"TwoQuarter"`
			ThreeQuarter string `json:"ThreeQuarter"`
			FourQuarter  string `json:"FourQuarter"`
		} `json:"NFL"`
		NBA struct {
			MoneyLine    string `json:"MoneyLine"`
			Spread       string `json:"Spread"`
			Total        string `json:"Total"`
			TeamTotal    string `json:"TeamTotal"`
			Game         string `json:"Game"`
			OneHalf      string `json:"OneHalf"`
			TwoHalf      string `json:"TwoHalf"`
			OneQuarter   string `json:"OneQuarter"`
			TwoQuarter   string `json:"TwoQuarter"`
			ThreeQuarter string `json:"ThreeQuarter"`
			FourQuarter  string `json:"FourQuarter"`
		} `json:"NBA"`
		MLB struct {
			MoneyLine      string `json:"MoneyLine"`
			Spread         string `json:"Spread"`
			Total          string `json:"Total"`
			TeamTotal      string `json:"TeamTotal"`
			Game           string `json:"Game"`
			OneFiveInnings string `json:"OneFiveInnings"`
		} `json:"MLB"`
	} `json:"SportsSettings"`
}

// Struct to hold Proxy Addresses
type Proxy struct {
	// MongoDB ID's
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	ObjectID string             `bson:"-"`

	ProxyAddress  string   `json:"ProxyAddress"`
	HTTPUserAgent string   `json:"HTTPUserAgent"`
	SubnetNumber  string   `json:"SubnetNumber"`
	UserName      string   `json:"UserName"`
	SiteName      string   `json:"SiteName"`
	BannedSites   struct{} `json:"BannedSites"`
}

// Anton Users for Front-End Login
type FrontEndUser struct {
	User          string `json:"User"`
	Password      string `json:"Password"`
	Authenticated bool   `bson:"-"`
}

// Telegram Message Struct
type TelegramMsg struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

// Slave Results for return when
type SlaveResults struct {
	PlacedLines  []Lines
	SkippedLines []Lines
	ErrorLines   []Lines
}

// This is a Struct to hold the Site's Dict, to help navigate through the different Sites, easier to read this way
type SiteDictionary struct {
	SiteName   string `json:"SiteName"`
	SportsDict struct {
		Baseball struct {
			ID  string `json:"ID"`
			MLB struct {
				ID     string `json:"ID"`
				Period struct {
					OneFiveInnings string `json:"1st 5 Innings"`
					Game           string `json:"Game"`
				} `json:"Period"`
			} `json:"MLB"`
		} `json:"Baseball"`
		Basketball struct {
			ID  string `json:"ID"`
			NBA struct {
				ID     string `json:"ID"`
				Period struct {
					OneHalf      string `json:"1st Half"`
					OneQuarter   string `json:"1st Quarter"`
					TwoHalf      string `json:"2nd Half"`
					TwoQuarter   string `json:"2nd Quarter"`
					ThreeQuarter string `json:"3rd Quarter"`
					FourQuarter  string `json:"4th Quarter"`
					Game         string `json:"Game"`
				} `json:"Period"`
			} `json:"NBA"`
		} `json:"Basketball"`
		Football struct {
			ID  string `json:"ID"`
			NFL struct {
				ID     string `json:"ID"`
				Period struct {
					OneHalf      string `json:"1st Half"`
					OneQuarter   string `json:"1st Quarter"`
					TwoHalf      string `json:"2nd Half"`
					TwoQuarter   string `json:"2nd Quarter"`
					ThreeQuarter string `json:"3rd Quarter"`
					FourQuarter  string `json:"4th Quarter"`
					Game         string `json:"Game"`
				} `json:"Period"`
			} `json:"NFL"`
		} `json:"Football"`
	} `json:"SportsDict"`
}
