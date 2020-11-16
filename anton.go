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
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	// This is to mark who created this Master, pulled from Anton Front End
	AntonOwner         string `bson:"antonowner" json:"antonowner"`
	AntonOwnerTelegram string `bson:"antonownertelegram" json:"antonownertelegram"`

	// These are populated when the Master and Lines are received from Scraping Scripts
	MasterName  string `bson:"mastername" json:"mastername"`
	MasterPass  string `bson:"masterpass" json:"masterpass"`
	LoginName   string `bson:"loginname" json:"loginname"`
	LoginPass   string `bson:"loginpass" json:"loginpass"`
	SiteName    string `bson:"sitename" json:"sitename"`
	AccountType string `bson:"accounttype" json:"accounttype"`
	Status      string `bson:"status" json:"status"`

	// Approved Lines sent from the feeder will be in this slice
	MasterLines []Lines `bson:"-" json:"masterlines"`

	// Proxy and HTTPUserAgent is saved in this struct
	Proxy Proxy `bson:"-" json:"-"`

	// If the account type is Agent, we will use a different login, and leverage this slice
	AgentPlayersSlice []string `bson:"-" json:"-"`

	// This slice is populated by pulling the DynamoDB for Slaves to compare against
	Slaves []Slave `bson:"-" json:"-"`

	// These are properties that will change during runtime, Status will be "Active" or "Error"
	ProgramError  []string `bson:"-" json:"-"`
	LoginAttempts int      `bson:"-" json:"-"`
	LinesOnPage   []string `bson:"-" json:"-"`

	// These will be time stamps for comparisons later
	TimeReceived time.Time `bson:"-" json:"-"`
}

// Create struct to hold info about Slaves pulled from the database
type Slave struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	// This is to mark who created this Slave, pulled from Anton Front End
	AntonOwner         string `bson:"antonowner" json:"antonowner"`
	AntonOwnerTelegram string `bson:"antonownertelegram" json:"antonownertelegram"`

	Status      string `bson:"status" json:"status"`
	SlaveName   string `bson:"slavename" json:"slavename"`
	SlavePass   string `bson:"slavepass" json:"slavepass"`
	SiteName    string `bson:"sitename" json:"sitename"`
	WagerAmount string `bson:"wageramount" json:"wageramount"`

	// These are populated by pulling from MongoDB
	SiteDictionary SiteDictionary `bson:"-" json:"-"`
	Profiles       []Profile      `bson:"-" json:"-"`
	Proxy          Proxy          `bson:"-" json:"-"`

	// Sort out the Lines after Creating and Comparing them
	AuthorizedLines []Lines `bson:"-" json:"-"`
}

// Create struct to hold info about lines pulled from scrapped sites
type Lines struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	// This is to mark who created this Line, pulled from Anton Front End
	AntonOwner         string `bson:"antonowner" json:"antonowner"`
	AntonOwnerTelegram string `bson:"antonownertelegram" json:"antonownertelegram"`

	// Created Timestamp
	CreatedTimestamp time.Time `bson:"createdtimestamp" json:"createdtimestamp"`

	// Will need to fill these out based from the scrapper
	RotationNumber  string `bson:"rotationnumber" json:"rotationnumber"`
	LineSpread      string `bson:"linespread" json:"linespread"`
	LineJuice       string `bson:"linejuice" json:"linejuice"`
	OverUnder       string `bson:"overunder,omitempty" json:"overunder,omitempty"`
	FavoredUnderdog string `bson:"favoredunderdog,omitempty" json:"favoredunderdog,omitempty"`
	LineType        string `bson:"linetype" json:"linetype"`
	TicketID        string `bson:"ticketid" json:"ticketid"`
	RiskAmount      string `bson:"riskamount" json:"riskamount"`
	ToWinAmount     string `bson:"towinamount" json:"towinamount"`

	// These are more for documenting purposes for the Database
	MasterName string `bson:"mastername" json:"mastername"`
	MasterPass string `bson:"masterpass" json:"masterpass"`
	MasterSite string `bson:"mastersite" json:"mastersite"`

	SlaveName string `bson:",omitempty" json:"slavename,omitempty"`
	SlavePass string `bson:",omitempty" json:"slavepass,omitempty"`
	SlaveSite string `bson:",omitempty" json:"slavesite,omitempty"`

	// This is only used by SlaveLines, inherited by the MasterLine it is comparing against
	MasterTicketID string `bson:"masterticketID,omitempty" json:"masterticketID,omitempty"`

	// This is Optional holder for any unique values for a particular line
	UniqueID     string `bson:"uniqueid,omitempty" json:"uniqueid"`
	HomeAway     string `bson:"homeaway,omitempty" json:"homeaway"`
	StringHolder string `bson:"stringholder,omitempty" json:"stringholder"`

	// These will be auto-populated by the lines from the scrapper
	Sport  string `bson:"sport" json:"sport"`
	League string `bson:"league" json:"league"`
	Period string `bson:"period" json:"period"`
	Team   string `bson:"team" json:"team"`

	LineCharacteristic string `bson:"linecharacteristic" json:"linecharacteristic"`

	// These can be inherited by Approved Bet, CurrentUser, or have default values for Potential:
	LineStatus         string   `bson:"linestatus" json:"linestatus"`
	CreatedViaFunction string   `bson:"-" json:"-"`
	BetType            string   `bson:"bettype" json:"bettype"`
	LineSpreadFloat    float64  `bson:"linespreadfloat" json:"linespreadfloat"`
	LineJuiceFloat     float64  `bson:"linejuicefloat" json:"linejuicefloat"`
	FunctionLog        string   `bson:"functionlog" json:"functionlog"`
	ErrorLog           []string `bson:"errorlog" json:"errorlog"`

	ComparedLines []Lines `bson:"comparedlines,omitempty"`
}

type Profile struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	// This is to mark who created this Profile, pulled from Anton Front End
	AntonOwner string `bson:"antonowner" json:"antonowner"`

	Status          string `bson:"status" json:"status"`
	MasterName      string `bson:"mastername" json:"mastername"`
	MasterSite      string `bson:"mastersite" json:"mastersite"`
	SlaveName       string `bson:"slavename" json:"slavename"`
	SlaveSite       string `bson:"slavesite" json:"slavesite"`
	SpreadParameter string `bson:"spreadparameter" json:"spreadparameter"`
	JuiceParameter  string `bson:"juiceparameter" json:"juiceparameter"`
	SportsSettings  struct {
		NFL struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"nfl" json:"nfl"`
		CollegeFootball struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"collegefootball" json:"collegefootball"`
		NBA struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"nba" json:"nba"`
		CollegeBasketball struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"collegebasketball" json:"collegebasketball"`
		MLB struct {
			MoneyLine      string `bson:"moneyline" json:"moneyline"`
			Spread         string `bson:"spread" json:"spread"`
			Total          string `bson:"total" json:"total"`
			TeamTotal      string `bson:"teamtotal" json:"teamtotal"`
			Game           string `bson:"game" json:"game"`
			OneFiveInnings string `bson:"onefiveinnings" json:"onefiveinnings"`
		} `bson:"mlb" json:"mlb"`
	} `bson:"sportssettings" json:"sportssettings"`
}

// Struct to hold Proxy Addresses
type Proxy struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	Status string `bson:"status" json:"status"`

	IPAddress     string `bson:"ipaddress" json:"ipaddress"`
	Subnet        string `bson:"subnet" json:"subnet"`
	ProxyAddress  string `bson:"proxyaddress" json:"proxyaddress"`
	HTTPUserAgent string `bson:"httpuseragent" json:"httpuseragent"`

	// Can delete Subnet Number
	SubnetNumber string `bson:"subnetnumber" json:"subnetnumber"`

	// Geolocation
	City      string `bson:"city" json:"city"`
	State     string `bson:"state" json:"state"`
	ZipCode   string `bson:"zipcode" json:"zipcode"`
	Country   string `bson:"country" json:"country"`
	Continent string `bson:"continent" json:"continent"`

	// IP Ownership
	Organization string `bson:"organization" json:"organization"`
	HostName     string `bson:"hostname" json:"hostname"`

	// If Assigned
	AccountType string   `bson:"accounttype" json:"accounttype"`
	LoginName   string   `bson:"loginname" json:"loginname"`
	SiteName    string   `bson:"sitename" json:"sitename"`
	BannedSites struct{} `bson:"bannedsites" json:"bannedsites"`
}

// Anton Users for Front-End Login
type AntonUser struct {
	Username      string `bson:"username" json:"username"`
	Password      string `bson:"password" json:"password"`
	Role          string `bson:"role" json:"role"`
	Telegram      string `bson:"telegram" json:"telegram"`
	Authenticated bool   `bson:"-" json:"-"`
}

// Telegram Message Struct
type TelegramMsg struct {
	ChatID int64  `bson:"-" json:"chat_id"`
	Text   string `bson:"-" json:"text"`
}

// Slave Results
type SlaveResults struct {
	// Section (c) of the Scrapper
	LinesOnPage []Lines `bson:"-" json:"-"`

	// Section (d) of the Scrapper
	ComparedLines []Lines `bson:"-" json:"-"`
	AuthLines     []Lines `bson:"-" json:"-"`
	SkippedLines  []Lines `bson:"-" json:"-"`

	// Section (e) of the Scrapper
	PlacedLines []Lines `bson:"-" json:"-"`
	ErrorLines  []Lines `bson:"-" json:"-"`
}

// Process Struct for Communication between Slave to Master
type Process struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	// Purpose -> Since it's in the Secrets Database, we split them by a purpose
	Purpose string `bson:"purpose" json:"purpose"`

	// Current Process ID
	CurrentID string `bson:"currentid" json:"currentid"`

	// Current Salt
	Salt string `bson:"salt" json:"salt"`
}

// Site Status
type SiteStatus struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	ObjectID string             `bson:"-" json:"-"`

	SiteName string `bson:"sitename" json:"sitename"`
	Master   struct {
		Agent struct {
			AntonEnable string `bson:"antonenable" json:"antonenable"`
			MLB         struct {
				MoneyLine      string `bson:"moneyline" json:"moneyline"`
				Spread         string `bson:"spread" json:"spread"`
				Total          string `bson:"total" json:"total"`
				TeamTotal      string `bson:"teamtotal" json:"teamtotal"`
				Game           string `bson:"game" json:"game"`
				OneFiveInnings string `bson:"onefiveinnings" json:"onefiveinnings"`
			} `bson:"mlb" json:"mlb"`
			NFL struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"nfl" json:"nfl"`
			NBA struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"nba" json:"nba"`
			CollegeFootball struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"collegefootball" json:"collegefootball"`
			CollegeBasketball struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"collegebasketball" json:"collegebasketball"`
		} `bson:"agent" json:"agent"`
		User struct {
			AntonEnable string `bson:"antonenable" json:"antonenable"`
			MLB         struct {
				MoneyLine      string `bson:"moneyline" json:"moneyline"`
				Spread         string `bson:"spread" json:"spread"`
				Total          string `bson:"total" json:"total"`
				TeamTotal      string `bson:"teamtotal" json:"teamtotal"`
				Game           string `bson:"game" json:"game"`
				OneFiveInnings string `bson:"onefiveinnings" json:"onefiveinnings"`
			} `bson:"mlb" json:"mlb"`
			NFL struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"nfl" json:"nfl"`
			NBA struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"nba" json:"nba"`
			CollegeFootball struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"collegefootball" json:"collegefootball"`
			CollegeBasketball struct {
				MoneyLine    string `bson:"moneyline" json:"moneyline"`
				Spread       string `bson:"spread" json:"spread"`
				Total        string `bson:"total" json:"total"`
				TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
				Game         string `bson:"game" json:"game"`
				OneHalf      string `bson:"onehalf" json:"onehalf"`
				TwoHalf      string `bson:"twohalf" json:"twohalf"`
				OneQuarter   string `bson:"onequarter" json:"onequarter"`
				TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
				ThreeQuarter string `bson:"threequarter" json:"threequarter"`
				FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
			} `bson:"collegebasketball" json:"collegebasketball"`
		} `bson:"user" json:"user"`
	} `bson:"master" json:"master"`
	Slave struct {
		AntonEnable string `bson:"antonenable" json:"antonenable"`
		MLB         struct {
			MoneyLine      string `bson:"moneyline" json:"moneyline"`
			Spread         string `bson:"spread" json:"spread"`
			Total          string `bson:"total" json:"total"`
			TeamTotal      string `bson:"teamtotal" json:"teamtotal"`
			Game           string `bson:"game" json:"game"`
			OneFiveInnings string `bson:"onefiveinnings" json:"onefiveinnings"`
		} `bson:"mlb" json:"mlb"`
		NFL struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"nfl" json:"nfl"`
		NBA struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"nba" json:"nba"`
		CollegeFootball struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"collegefootball" json:"collegefootball"`
		CollegeBasketball struct {
			MoneyLine    string `bson:"moneyline" json:"moneyline"`
			Spread       string `bson:"spread" json:"spread"`
			Total        string `bson:"total" json:"total"`
			TeamTotal    string `bson:"teamtotal" json:"teamtotal"`
			Game         string `bson:"game" json:"game"`
			OneHalf      string `bson:"onehalf" json:"onehalf"`
			TwoHalf      string `bson:"twohalf" json:"twohalf"`
			OneQuarter   string `bson:"onequarter" json:"onequarter"`
			TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
			ThreeQuarter string `bson:"threequarter" json:"threequarter"`
			FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
		} `bson:"collegebasketball" json:"collegebasketball"`
	} `bson:"slave" json:"slave"`
}

// This is a Struct to hold the Site's Dict, to help navigate through the different Sites, easier to read this way
type SiteDictionary struct {
	SiteName   string `bson:"sitename" json:"sitename"`
	SportsDict struct {
		Baseball struct {
			ID  string `bson:"id" json:"id"`
			MLB struct {
				ID     string `bson:"id" json:"ID"`
				Period struct {
					OneFiveInnings string `bson:"onefiveinnings" json:"onefiveinnings"`
					Game           string `bson:"game" json:"game"`
				} `bson:"period" json:"period"`
			} `bson:"mlb" json:"mlb"`
		} `bson:"baseball" json:"baseball"`
		Basketball struct {
			ID  string `bson:"id" json:"id"`
			NBA struct {
				ID     string `bson:"id" json:"id"`
				Period struct {
					OneHalf      string `bson:"onehalf" json:"onehalf"`
					OneQuarter   string `bson:"onequarter" json:"onequarter"`
					TwoHalf      string `bson:"twohalf" json:"twohalf"`
					TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
					ThreeQuarter string `bson:"threequarter" json:"threequarter"`
					FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
					Game         string `bson:"game" json:"game"`
				} `bson:"period" json:"period"`
			} `bson:"nba" json:"nba"`
			CollegeBasketball struct {
				ID     string `bson:"id" json:"id"`
				Period struct {
					OneHalf      string `bson:"onehalf" json:"onehalf"`
					OneQuarter   string `bson:"onequarter" json:"onequarter"`
					TwoHalf      string `bson:"twohalf" json:"twohalf"`
					TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
					ThreeQuarter string `bson:"threequarter" json:"threequarter"`
					FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
					Game         string `bson:"game" json:"game"`
				} `bson:"period" json:"period"`
			} `bson:"collegebasketball" json:"collegebasketball"`
		} `bson:"basketball" json:"basketball"`

		Football struct {
			ID  string `bson:"id" json:"id"`
			NFL struct {
				ID     string `bson:"id" json:"id"`
				Period struct {
					OneHalf      string `bson:"onehalf" json:"onehalf"`
					TwoHalf      string `bson:"twohalf" json:"twohalf"`
					OneQuarter   string `bson:"onequarter" json:"onequarter"`
					TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
					ThreeQuarter string `bson:"threequarter" json:"threequarter"`
					FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
					Game         string `bson:"game" json:"game"`
				} `bson:"period" json:"period"`
			} `bson:"nfl" json:"nfl"`
			CollegeFootball struct {
				ID     string `bson:"id" json:"id"`
				Period struct {
					OneHalf      string `bson:"onehalf" json:"onehalf"`
					TwoHalf      string `bson:"twohalf" json:"twohalf"`
					OneQuarter   string `bson:"onequarter" json:"onequarter"`
					TwoQuarter   string `bson:"twoquarter" json:"twoquarter"`
					ThreeQuarter string `bson:"threequarter" json:"threequarter"`
					FourQuarter  string `bson:"fourquarter" json:"fourquarter"`
					Game         string `bson:"game" json:"game"`
				} `bson:"period" json:"period"`
			} `bson:"collegefootball" json:"collegefootball"`
		} `bson:"football" json:"football"`
	} `bson:"sportsdict" json:"sportsdict"`
}
