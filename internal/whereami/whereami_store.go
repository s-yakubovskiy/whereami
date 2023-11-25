package whereami

func (l *Locator) Store() {
	ip, err := l.client.GetIP()
	if err != nil {
		errorln(err.Error())
	}
	location, err := l.client.GetLocation(ip)
	if err != nil {
		errorln(err.Error())
	}

	if location != nil && ip != "" {
		vpninterfaces, err := l.dbclient.GetVPNInterfaces()
		if err != nil {
			warnln(err.Error())
		}

		vpn, err := l.client.GetVPN(vpninterfaces)
		if err != nil {
			warnln(err.Error())
		}
		if vpn {
			location.Vpn = true
		}
		if l.ipquality {
			l.client.AddIPQuality(location, ip)
		}
		if err := l.dbclient.StoreLocation(location); err != nil {
			if err.Error() == "The database is already contains this record." {
				warnln(err.Error())
			} else {
				errorln(err.Error())
			}
		}
	}
}
