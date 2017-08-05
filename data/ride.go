package data

import "time"

type Ride struct {
	Id             int `json:"id"`
	StartDescrip   string `json:"startDescrip"`
	EndDescrip     string `json:"endDescrip"`
	CreatedAt      string `json:"createdAt"`
	LocId          int `json:"locId"`
	AvailableSeats int `json:"availableSeats"`
	NeededSeats    int `json:"neededSeats"`
	TimeLeaving    string `json:"timeLeaving"`
	TimePickUp     string `json:"timePickUp"`
	UserId         int `json:"userId"`
	Uuid           string `json:"uuid"`
	VehicleMake    string `json:"carMake"`
	VehicleModel   string `json:"carModel"`
	VehicleYear    int `json:"carYear"`
	VehicleId	  int `json:"carId"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
}

func (ride *Ride) Create(locationId int) (err error ) {
	statement := `INSERT INTO ride ( startDescrip, endDescrip, createdAt, locId )
	 VALUES ($1, $2, $3, $4)
	 RETURNING id, startDescrip, endDescrip, createdAt, locId`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(ride.StartDescrip, ride.EndDescrip, time.Now(), locationId).
		Scan(&ride.Id, &ride.StartDescrip, &ride.EndDescrip, &ride.CreatedAt, &ride.LocId)

	return

}

func (ride *Ride) CreateRideOffered(vehicleId int) (err error) {
	statement := `INSERT INTO rideOffered (rideId, availableSeats, timeLeaving, vehicleId )
		VALUES($1, $2, $3, $4)
		RETURNING availableSeats, timeLeaving, vehicleId`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(ride.Id, ride.AvailableSeats, ride.TimeLeaving, vehicleId).
		Scan(&ride.AvailableSeats, &ride.TimeLeaving, &ride.VehicleId)

	return
}

func (ride *Ride) CreateRideNeeded() (err error) {
	statement := `INSERT INTO rideNeeded (rideId,  neededSeats, timePickUp, riderId)
		VALUES($1, $2, $3, $4)
		RETURNING neededSeats, timePickUP, riderId`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(ride.Id, ride.NeededSeats, ride.TimePickUp, ride.UserId).
		Scan(&ride.NeededSeats, &ride.TimePickUp, &ride.UserId)

	return

}