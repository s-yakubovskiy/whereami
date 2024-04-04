package whereami

import (
	"github.com/s-yakubovskiy/whereami/internal/common"
)

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
	// NOTE: right now no difference between full and short `show` output
	l.ShowFull()
}

func (l *Locator) ShowFull() {
	var ip string
	var err error

	if l.iplookup == "" {
		ip, err = l.client.GetIP()
		if err != nil {
			common.Errorln(err.Error())
		}
	} else {
		ip = l.iplookup
	}
	// Fetching data from IP API
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

		// update comment for location
		location.Comment += ". Using public ip provider: " + l.client.ShowIpProvider()

		location.Output(categories, orderedCategories)
	}
}
