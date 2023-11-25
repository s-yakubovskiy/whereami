package contracts

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

type Location struct {
	IP          string  `json:"ip"`
	Country     string  `json:"country"`
	CountryCode string  `json:"country_code"`
	Region      string  `json:"region"`
	RegionCode  string  `json:"region_code"`
	City        string  `json:"city"`
	Timezone    string  `json:"timezone"`
	Zip         string  `json:"zip"`
	Postal      string  `json:"postal"`
	Flag        string  `json:"flag"`
	EmojiFlag   string  `json:"emoji_flag"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	Asn         string  `json:"asn"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Date        string  `json:"date"`
	Vpn         bool    `json:"vpn"`
}

func (l *Location) Output(fields ...string) {
	val := reflect.ValueOf(*l)
	typ := val.Type()

	fieldMap := make(map[string]reflect.Value)
	for i := 0; i < val.NumField(); i++ {
		fieldMap[strings.ToLower(typ.Field(i).Name)] = val.Field(i)
	}

	for _, field := range fields {
		fieldValue, exists := fieldMap[strings.ToLower(field)]
		if !exists {
			color.Red("Field %s does not exist", field)
			continue
		}
		// color.Cyan("%s: ", field)

		cyan := color.New(color.FgCyan)
		cyan.Add(color.Bold)

		cyanP := cyan.PrintfFunc()
		cyanP("%s: ", field)
		fmt.Println(fieldValue.Interface())
	}
}
