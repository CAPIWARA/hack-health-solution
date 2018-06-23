package users

import (
	"errors"
	"github.com/kr/pretty"
	"hack-health-solution/server/dbs"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id         string `json:"-" bson:"_id"`
	Name       string `json:"name"`
	Points     int    `json:"points"`
	SocialName string `json:"socialname,omitempty"`
}

func GetUser(id string) (*User, error) {
	var data User
	c := dbs.Session.DB(dbs.Database).C("users")

	if err := c.Find(bson.M{"_id": id}).One(&data); err != nil {
		pretty.Log("error: ", err)
		return nil, errors.New("user not found")
	}
	pretty.Log(data)
	return &data, nil
}
