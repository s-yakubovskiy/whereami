package whereami

import "github.com/s-yakubovskiy/whereami/internal/common"

var (
	categories = map[string][]string{
		"Geographical Information": {
			"country", "countryCode", "region", "regionCode",
			"city", "timezone", "zip", "latitude", "longitude",
		},
		"Network Information": {
			"ip", "isp", "asn", "flag",
		},
		"Security Assessments": {
			"vpn", "scores",
		},
		"Miscellaneous": {
			"map", "date", "comment",
		},
	}

	orderedCategories = []string{
		"Network Information",
		"Geographical Information",
		"Security Assessments",
		"Miscellaneous",
	}
)

func (l *Locator) Show() {
	// Fetching data from IP API
	ip, err := l.client.GetIP()
	if err != nil {
		common.Errorln(err.Error())
	}
	location, err := l.client.GetLocation(ip)
	if err != nil {
		common.Errorln(err.Error())
	}
	if location != nil && ip != "" {
		vpninterfaces, err := l.dbclient.GetVPNInterfaces()
		if err != nil {
			common.Warnln(err.Error())
		}

		vpn, err := l.client.GetVPN(vpninterfaces)
		if err != nil {
			common.Warnln(err.Error())
		}

		if vpn {
			location.Vpn = true
		}

		if l.ipquality {
			l.client.AddIPQuality(location, ip)
		}

		// output to stding colorized
		// TODO: output will be always full for now (reconsider adding --short flag later)
		// location.Output(
		// 	"ip",
		// 	"country",
		// 	"region",
		// 	"regioncode",
		// 	"city",
		// 	"timezone",
		// 	"vpn",
		// 	"comment",
		// )
	}

	// location.Output(
	// 	"ip", "country", "countryCode", "region", "regionCode",
	// 	"city", "timezone", "zip", "flag",
	// 	"isp", "asn", "latitude", "longitude", "vpn", "comment", "scores",
	// )

	location.Output(categories, orderedCategories)
}

func (l *Locator) ShowFull() {
	// Fetching data from IP API
	ip, err := l.client.GetIP()
	if err != nil {
		common.Errorln(err.Error())
	}
	location, err := l.client.GetLocation(ip)
	if err != nil {
		common.Errorln(err.Error())
	}
	if location != nil && ip != "" {
		vpninterfaces, err := l.dbclient.GetVPNInterfaces()
		if err != nil {
			common.Warnln(err.Error())
		}

		vpn, err := l.client.GetVPN(vpninterfaces)
		if err != nil {
			common.Warnln(err.Error())
		}

		if vpn {
			location.Vpn = true
		}

		if l.ipquality {
			l.client.AddIPQuality(location, ip)
		}

		// output to stding colorized
		// location.Output(
		// 	"ip", "country", "countryCode", "region", "regionCode",
		// 	"city", "timezone", "zip", "flag",
		// 	"isp", "asn", "latitude", "longitude", "vpn", "comment", "scores",
		// )
		location.Output(categories, orderedCategories)
	}
}
