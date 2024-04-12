package whereami

import "github.com/s-yakubovskiy/whereami/internal/common"

func (l *Locator) Store() {
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
		if l.cfg.IpQuality {
			l.client.AddIPQuality(location, ip)
		}
		if err := l.dbclient.StoreLocation(location); err != nil {
			if err.Error() == "The database is already contains this record." {
				common.Warnln(err.Error())
			} else {
				common.Errorln(err.Error())
			}
		}
	}
}
