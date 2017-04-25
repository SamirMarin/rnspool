package data

import "time"

type Ride struct {
	Id             int `json: "id"`
	StartDescrip   string `json: startDescrip`
	EndDescrip     string `json: endDescrip`
	CreatedAt      time.Time `json: createdAt`
	LocId          int `json: locId`
	AvailableSeats int `json: availableSeats`
	NeededSeats    int `json: neededSeats`
	TimeLeaving    time.Time `json: timeLeaving`
	TimePickUp     time.Time `json: timePickUp`
	UserId         int `json: userId`
	Uuid           string `json: uuid`
	VehicleMake    string `json: carMake`
	VehicleModel   string `json: carModel`
	VehicleYear    int `json: carYear`
	City           string `json: city`
	Province       string `json: province`
	Country        string `json: province`
}

func (ride *Ride) Create(locationId int) (err error ) {
	statement := `INSERT INTO ride ( startDescrip, endDescrip, createdAt, locId )
	 VALUES ($1, $2, $3, $4)
	 RETURNING startDescrip, endDescrip, createdAt, locId`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(ride.StartDescrip, ride.EndDescrip, time.Now, locationId).
		Scan(&ride.StartDescrip, &ride.EndDescrip, &ride.CreatedAt, &ride.LocId)

	return

}

func (ride *Ride) CreateRideOffered(vehicleId int) (err error) {
	statement := `INSERT INTO rideOffered (startDescrip, endDescrip, availableSeats, timeLeaving, vehicleId )
		VALUES($1, $2, $3, $4, $5)
		RETURNING startDescrip, endDescrip, availableSeats, timeLeaving`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(ride.StartDescrip, ride.EndDescrip, ride.AvailableSeats, ride.TimeLeaving).
		Scan(&ride.StartDescrip, &ride.EndDescrip, &ride.AvailableSeats, &ride.TimeLeaving)

	return
}

func (ride *Ride) CreateRideNeeded() (err error) {
	statement := `INSERT INTO rideNeeded (startDescrip, endDescrip, neededSeats, timePickUp, rideId)
		VALUES($1, $2, $3, $4, $5)
		RETURNING startDescrip, endDescrip, neededSeats, timePickUP, rideId`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(ride.StartDescrip, ride.EndDescrip, ride.NeededSeats, ride.TimePickUp, ride.UserId).
		Scan(&ride.StartDescrip, &ride.EndDescrip, &ride.NeededSeats, &ride.TimePickUp, &ride.UserId)

	return

}