package dbclient

func (s *LocationKeeper) recordExists(query, city string) (bool, error) {
	var exists bool
	stmt, err := s.db.Prepare("SELECT EXISTS(SELECT 1 FROM locations WHERE ip = ? AND city = ? LIMIT 1)")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(query, city).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
