package data

import (
	"database/sql"
	"fmt"
)

type Location struct {
	Id       int
	City     string
	Province string
	Country  string
}

func (location *Location) Create() (err error) {
	fmt.Println("here", location.Id)
	fmt.Println(location.City)

	err = Db.QueryRow(`SELECT id
	FROM location
	WHERE city = $1 AND province = $2 AND country = $3`,
		location.City, location.Province, location.Country).
		Scan(&location.Id)
	//if err != nil {
	//	return
	//}
	fmt.Println("you", location.Id)

	if location.Id == 0 {
		statement := `INSERT INTO location(city, province, country)
		VALUES ($1, $2, $3)
		RETURNING id, city, province, country `

		var insertStmt *sql.Stmt
		insertStmt, err = Db.Prepare(statement)
		if err != nil {
			return
		}
		defer insertStmt.Close()

		err = insertStmt.QueryRow(location.City, location.Province, location.Country).
			Scan(&location.Id, &location.City, &location.Province, &location.Country)
		fmt.Println(location.Id)
		return

	}
	return
}
