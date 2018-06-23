package dbs

import (
	"gopkg.in/mgo.v2"
	"log"
)

var Session *mgo.Session
var Database = "hackhealth"

func NewSession() error {
	session, err := mgo.Dial("mongodb://cardosomarcos:cardosomarcos10@ds117431.mlab.com:17431/hackhealth")

	if err != nil {
		log.Printf("db error: %v", err)
		return err
	}
	session.SetMode(mgo.Monotonic, true)
	Session = session
	return nil
}
