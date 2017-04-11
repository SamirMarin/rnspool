package data

import (
	"time"
)

type User struct {
	Id        int       `json:"id"`
	Uuid      string    `json:"uuid"`
	Email     string    `json: "email"`
	Password  string    `josn: "password"`
	FirstName string    `json: "firstname"`
	LastName  string    `json: "lastname"`
	AboutMe   string    `json: "aboutme"`
	AptNum    int       `json: "aptnum"`
	HouseNum  int       `json: "housenum"`
	Street    string    `json: "street"`
	CreatedAt time.Time `json: "createdat"`
}
type Session struct {
	Id        int
	Uuid      string
	Email     string
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

/**
User creates an account
*/
func (user *User) Create() (err error) {
	statement := "INSERT INTO user_table ( uuid, email, password, firstName, lastName, " +
		"aboutMe, aptNum, houseNum, street, createdAt )" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)" +
		"RETURNING id, uuid, createdAt"

	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(createUUID(), user.Email, Encrypt(user.Password),
		user.FirstName, user.LastName, user.AboutMe,
		user.AptNum, user.HouseNum, user.Street, time.Now()).
		Scan(&user.Id, &user.Uuid, &user.CreatedAt)
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

/**
User login, creates a session for the user
*/
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO session_table (uuid, email, userId, createAt)" +
		"VALUES ($1, $2, $3, $4)" +
		"RETURNING id, uuid, email, userId, createAt"
	insertStmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer insertStmt.Close()

	err = insertStmt.QueryRow(createUUID(), user.Email, user.Id, time.Now()).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

/**
check in user session is still active
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
		`SELECT id, uuid, email, firstName, lastName, aboutMe
	FROM user_table
	WHERE email = $1`,
		email).
		Scan(&user.Id, user.Uuid, user.Email, user.FirstName, user.LastName, user.AboutMe)
	return
}
