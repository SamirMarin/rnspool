package data

type Location struct {
	Id int
	City string
	Province string
	Country string
}

func (location *Location) Create() (err error) {
	statement := `INSERT INTO location(city, province, country)
	VALUES ($1, $2, $3)
	RETURNING id, city, province, country `
	
	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return 
	}
	defer insertStmt.Close()
	
	err = insertStmt.QueryRow(location.City, location.Province, location.Country).
		Scan(&location.Id, &location.City, &location.Province, &location.Country)
	return
}

