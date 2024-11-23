package entity

type Location struct {
	IP          string         `json:"ip"`
	Country     string         `json:"country"`
	CountryCode string         `json:"country_code"`
	Region      string         `json:"region"`
	RegionCode  string         `json:"region_code"`
	City        string         `json:"city"`
	Timezone    string         `json:"timezone"`
	Zip         string         `json:"zip"`
	Flag        string         `json:"flag"`
	Isp         string         `json:"isp"`
	Asn         string         `json:"asn"`
	Latitude    float64        `json:"latitude"`
	Longitude   float64        `json:"longitude"`
	Date        string         `json:"date"`
	Vpn         bool           `json:"vpn"`
	Comment     string         `json:"comment"`
	Scores      LocationScores `json:"scores"`
	Gps         GPSReport      `json:"gps"`
	Map         string         `json:"map"`
}

type GPSReport struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Url       string  `json:"url"`
}

type LocationScores struct {
	FraudScore  int    `json:"fraud_score"`
	IsCrawler   bool   `json:"is_crawler"`
	Host        string `json:"host"`
	Proxy       bool   `json:"proxy"`
	VPN         bool   `json:"vpn"`
	Tor         bool   `json:"tor"`
	RecentAbuse bool   `json:"recent_abuse"`
	BotStatus   bool   `json:"bot_status"`
}

func NewLocationScores() *LocationScores {
	return &LocationScores{}
}
