package data

type Vehicle struct {
	Id               int    `json:"id"`
	Licence          string `json:"licence"`
	Make             string `json:"make"`
	Model            string `json: "model"`
	Year             int    `json: "year"`
	NumberPassengers int    `json: "numPassengers"`
	Type             string `json: "type"`
	DriverId         int    `json: "driverId"`
	Uuid             string `json: "uuid"`
}

func (vehicle *Vehicle) Create() (err error) {
	statement := `INSERT INTO Vehicle (licence, make, model, year, numPassengers, type, driverId)
	VALUES($1, $2, $3, $4, $5, $6, $7)
	RETURNING id, licence, make, model, year, numPassengers, type, driverId`

	insertStmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(vehicle.Licence, vehicle.Make, vehicle.Model, vehicle.Year,
		vehicle.NumberPassengers, vehicle.Type, vehicle.DriverId).
		Scan(&vehicle.Id, &vehicle.Licence, &vehicle.Make, &vehicle.Model, &vehicle.Year,
			&vehicle.NumberPassengers, &vehicle.Type, &vehicle.DriverId)
	return

}
