package data

import (
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=rnsdev"+
		" dbname=rnsdev password=rnsdev sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

//create a random UUID from RFC 4122
// obtained from go-web-programming chitchat app
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("cannot generate UUID", err)
	}

	//  0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7f
	// set the four most sig bits (bits 12 through 15) of the
	//time_hi_and_version field to the 4-bit version number
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(str string) (crypstr string) {
	crypstr = fmt.Sprintf("%x", sha1.Sum([]byte(str)))
	return
}
