package contracts

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

type Location struct {
	Status      string  `json:"status"`
	Country     string  `json:"country"`
	CountryCode string  `json:"countryCode"`
	Region      string  `json:"region"`
	RegionName  string  `json:"regionName"`
	City        string  `json:"city"`
	Zip         string  `json:"zip"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Timezone    string  `json:"timezone"`
	Isp         string  `json:"isp"`
	Org         string  `json:"org"`
	As          string  `json:"as"`
	IP          string  `json:"query"`
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
