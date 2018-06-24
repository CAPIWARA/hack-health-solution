package users

import (
	"errors"
	"github.com/kr/pretty"
	"hack-health-solution/server/dbs"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id             string   `json:"-" bson:"_id"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Birthday       string   `json:"birthday"`
	Sarradinhas    int      `json:"sarradinhas"`
	SocialName     string   `json:"socialName"`
	Orientation    string   `json:"orientation"`
	Ethnicity      string   `json:"ethnicity"`
	GenderIdentity string   `json:"genderIdentity"`
	Image          string   `json:"image"`
	LastExam       string   `json:"lastExam"`
	Friends        []string `json:"friends"`
}

type Friend struct {
	Name        string `json:"name"`
	Sarradinhas int    `json:"sarradinhas"`
	Image       string `json:"image"`
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

func (user *User) CreateExam() (error) {
	c := dbs.Session.DB(dbs.Database).C("users")
	if err := c.Update(bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"lastexam": GetCurrentGmtDate()}}); err != nil {
		pretty.Log(err)
		return err
	}
	user.LastExam = GetCurrentGmtDate()
	return nil
}

func AddSarradinhas(total int, userid string) (error) {
	//user, err := GetUser(userid)
	//if err != nil {
	//	pretty.Log(err)
	//	return err
	//}
	c := dbs.Session.DB(dbs.Database).C("users")
	if err := c.Update(bson.M{"_id": userid}, bson.M{"$inc": bson.M{"sarradinhas": total}}); err != nil {
		pretty.Log(err)
		return err
	}
	return nil
}
