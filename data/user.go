package data

import (
	"time"
	"database/sql"
)

type User struct {
	Id         int       `json:"id"`
	Uuid       string    `json:"uuid"`
	Email      string    `json:"email"`
	Password   string    `josn:"password"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	AboutMe    string    `json:"aboutMe"`
	AptNum     int       `json:"aptNum"`
	HouseNum   int       `json:"houseNum"`
	Street     string    `json:"street"`
	PostalCode string    `json:"postalCode"`
	City       string    `json:"city"`
	Province   string    `json:"province"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"createdat"`
	AddressId  int       `json:"addressId"`
}
type Session struct {
	Id        int
	Uuid      string
	Email     string
	FirstName string
	LastName  string
	AboutMe   string
	UserId    int
	CreatedAt time.Time
}
type Rider struct {
	UserId      int
	RiderRating float64
}
type Driver struct {
	UserId       int
	DriverRating float64
}

type Address struct {
	Id         int
	AptNum     int
	HouseNum   int
	Street     string
	PostalCode string
	LocationId int
}

/**
User creates an account
*/
func (user *User) Create() (err error) {
	statement := "INSERT INTO user_table ( uuid, email, password, firstName, lastName, " +
		"aboutMe, createdAt, addressId )" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)" +
		"RETURNING id, uuid, email, firstName, lastName, aboutMe, createdAt, addressId"

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(createUUID(), user.Email, Encrypt(user.Password),
		user.FirstName, user.LastName, user.AboutMe, time.Now(), user.AddressId).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.FirstName, &user.LastName,
			&user.AboutMe, &user.CreatedAt, &user.AddressId)
	return
}
func (driver *Driver) Create() (err error) {
	statement := "INSERT INTO driver (userId)" +
		"VALUES($1)" +
		"RETURNING userId"
	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(driver.UserId).Scan(&driver.UserId)
	return
}
func (rider *Rider) Create() (err error) {
	statement := "INSERT INTO rider (userId)" +
		"VALUES($1)" +
		"RETURNING userId"
	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(rider.UserId).Scan(&rider.UserId)
	return
}

func (address *Address) Create() (err error) {
	err = Db.QueryRow(`SELECT id
	FROM address
	WHERE aptNum = $1 AND houseNume = $2 AND street = $3 AND postalCode = $4`,
	address.AptNum, address.HouseNum, address.Street, address.PostalCode).
		Scan(&address.Id)

	if address.Id == 0 {
		statement := `INSERT INTO address (aptNum, houseNum, street, postalCode, locationId)
		VALUES($1, $2, $3, $4, $5)
		RETURNING id, aptNum, houseNum, street, postalCode, locationId`

		var insertStmt *sql.Stmt
		insertStmt, err = Db.Prepare(statement)
		if err != nil {
			return
		}
		defer insertStmt.Close()

		err = insertStmt.QueryRow(address.AptNum, address.HouseNum, address.Street,
			address.PostalCode, address.LocationId).
			Scan(&address.Id, &address.AptNum, &address.HouseNum, &address.Street,
			&address.PostalCode, &address.LocationId)
		return
	}
	return
}

/**
User login, creates a session for the user
*/
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO session_table (uuid, email, firstName, lastName," +
		"aboutMe, userId, createdAt)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)" +
		"RETURNING id, uuid, email, firstName, lastName, aboutMe, userId, createdAt"
	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(createUUID(), user.Email, user.FirstName, user.LastName,
		user.AboutMe, user.Id, time.Now()).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.FirstName, &session.LastName,
			&session.AboutMe, &session.UserId, &session.CreatedAt)
	return
}

/**
check if user session is still active
*/
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow(
		`SELECT id, uuid, email, creatAt
	FROM session_table
	WHERE uuid = $1`,
		session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}
func UserByEmail(email string) (user User, err error) {
	err = Db.QueryRow(
		`SELECT id, uuid, email, password, firstName, lastName, aboutMe
	FROM user_table
	WHERE email = $1`,
		email).
		Scan(&user.Id, &user.Uuid, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.AboutMe)
	return
}
