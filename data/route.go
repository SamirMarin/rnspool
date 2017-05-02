package data

import (
	"database/sql"
	"errors"
)

type Route struct {
	Id int
	StartDescrip string
	EndDescrip string
	RideId int
	Description string
	Legs []Leg
}

type Leg struct {
	StartPointLat float64
	StartPointLon float64
	EndPointLat float64
	EndPointLon float64
	HtmlInstr string
	Duration int64
	Distance int
	RouteId int
}

func (route *Route) Create() (err error) {
	err = Db.QueryRow(`SELECT id
	FROM route
	WHERE startDescrip = $1 AND endDescrip = $2 AND description = $3`,
		route.StartDescrip, route.EndDescrip, route.Description).
		Scan(&route.Id)
	if err != nil {
		return
	}
	if route.Id == 0 {
		statement := `INSTERT INTO route (startDescrip, endDescrip, description)
	 	VALUES($1, $2, $3)
	 	RETURNING id, startDescrip, endDescrip, description`

		var insertStmt *sql.Stmt
		insertStmt, err = Db.Prepare(statement)
		if err != nil {
			return
		}
		defer insertStmt.Close()

		err = insertStmt.QueryRow(route.StartDescrip, route.EndDescrip, route.Description).
			Scan(&route.Id, &route.StartDescrip, &route.EndDescrip, &route.Description)
		return

	}
	return
}



func (leg *Leg) Create(routeId int) (err error) {
	statement := `INSERT INTO leg ( startPointLat, startPointLon, endPointLat, endPointLon,
		htmlInstr, duration, distance, routeId
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING startPointLat, startPointLon, endPointLat, endPointLon`

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(leg.StartPointLat, leg.StartPointLon, leg.EndPointLat, leg.EndPointLon,
		leg.HtmlInstr, leg.Duration, leg.Distance, routeId).
		Scan(&leg.StartPointLat, &leg.StartPointLon, &leg.EndPointLat, &leg.EndPointLon)

	return
}

func CreateRideHasRouteByIds(rideId int, routeId int) (err error) {
	statement := `INSERT INTO rideHasRoute ( rideId, routeId )
	 VALUES ($1, $2)`
	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	result, err := insertStmt.Exec(rideId, routeId)
	rows, err := result.RowsAffected()
	if err != nil {
		return
	}
	if rows < 0 {
		err = errors.New("no rows changed on rideId, routeId insert")
		if err != nil {
			return
		}
	}
	return
}
