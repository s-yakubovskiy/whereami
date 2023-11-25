package whereami

import "github.com/s-yakubovskiy/whereami/internal/common"

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
		location.Output(
			"ip",
			"country",
			"region",
			"regioncode",
			"city",
			"timezone",
			"vpn",
			"comment",
		)
	}
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
		location.Output(
			"ip", "country", "countryCode", "region", "regionCode",
			"city", "timezone", "zip", "flag",
			"isp", "asn", "latitude", "longitude", "vpn", "comment", "scores",
		)
	}
}