package db

func (s *LocationKeeper) AddVPNInterface(interfaceName string) error {
	stmt, err := s.db.Prepare("INSERT INTO vpn_interfaces (interface_name) VALUES (?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(interfaceName)
	if err != nil {
		return err
	}

	return nil
}

func (s *LocationKeeper) ListVPNInterfaces() ([]string, error) {
	rows, err := s.db.Query("SELECT interface_name FROM vpn_interfaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var interfaces []string
	for rows.Next() {
		var interfaceName string
		if err := rows.Scan(&interfaceName); err != nil {
			return nil, err
		}
		interfaces = append(interfaces, interfaceName)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return interfaces, nil
}
