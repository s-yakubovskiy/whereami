package contracts

import (
	"reflect"

	"github.com/fatih/color"
)

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

// Output reworked to support ordered categories with conditional field output and correct struct handling
func (l *Location) Output(categories map[string][]string, orderedCategories []string) {
	val := reflect.ValueOf(*l)

	// Prepare color outputs
	cyan := color.New(color.FgCyan).Add(color.Bold)
	white := color.New(color.FgWhite)                     // For field values
	magenta := color.New(color.FgMagenta).Add(color.Bold) // For categories
	cyanP := cyan.PrintfFunc()
	whiteP := white.PrintfFunc()
	magentaP := magenta.PrintfFunc() // Print function for category names in magenta

	// Iterate through each category in the ordered list
	for _, category := range orderedCategories {
		fields, exists := categories[category]
		if !exists {
			continue // Skip categories not defined in the map
		}
		magentaP("\n%s\n", capitalizeFirst(category)) // Print category name in magenta
		for _, field := range fields {
			fieldVal, found := findField(val, field)
			if !found {
				continue // Skip fields not found
			}
			if field == "date" {
				handleDateField(fieldVal, cyanP, whiteP)
				continue
			}

			if field == "map" {
				handleMapField(fieldVal, cyanP, whiteP, l.Latitude, l.Longitude)
			}
			if field == "gps" {
				printStructFields(fieldVal, cyanP, whiteP)
				continue
			}

			fieldValue := fieldVal.Interface()
			if reflect.TypeOf(fieldValue).Kind() == reflect.Struct && field == "scores" {
				// Handle struct types specifically for 'scores'
				if field != "scores" {
					cyanP("  %s:\n", capitalizeFirst(field))
				}
				printStructFields(fieldVal, cyanP, whiteP)
			} else if !isEmpty(fieldValue) {
				// Skip empty fields except for 'date', which is handled separately
				cyanP("  %s: ", capitalizeFirst(field))
				whiteP("%v\n", fieldValue)
			}
		}
	}
}
