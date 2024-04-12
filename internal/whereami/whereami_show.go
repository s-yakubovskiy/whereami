package whereami

import (
	"log"

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
		"GPS": {
			"gps",
		},
		"Miscellaneous": {
			"map", "date", "comment",
		},
	}

	orderedCategories = []string{
		"Network Information",
		"Geographical Information",
		"Security Assessments",
		"GPS",
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

	if l.cfg.IP == "" {
		ip, err = l.client.GetIP()
		if err != nil {
			common.Errorln(err.Error())
		}
	} else {
		ip = l.cfg.IP
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

		if l.cfg.IpQuality {
			l.client.AddIPQuality(location, ip)
		}

		if l.cfg.GPS {
			// TODO: move to l.client.AddGPSInfo
			report, err := l.FetchGPSReport()
			if err != nil {
				log.Fatal(err)
			}
			location.Gps.Altitude = report.Alt
			location.Gps.Longitude = report.Lon
			location.Gps.Latitude = report.Lat
			location.Comment += ". Enriched by GPSD"
		}

		location.Comment += ". Using public ip provider: " + l.client.ShowIpProvider()
		location.Output(categories, orderedCategories)
	}
}
