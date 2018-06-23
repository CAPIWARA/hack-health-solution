package users

import (
	"hack-health-solution/server/dbs"
	"github.com/kr/pretty"
	"gopkg.in/mgo.v2/bson"
)

type Sarrada struct {
	Id         bson.ObjectId `json:"-" bson:"_id"`
	UserId     string        `json:"userid"`
	Camisinha  bool          `json:"camisinha"`
	Oral       bool          `json:"oral"`
	Pessoa     string        `json:"pessoa"`
	Quantidade int           `json:"quantidade"`
	Drogas     bool          `json:"drogas"`
	Ejaculou   bool          `json:"ejaculou"`
}

func (sarrada *Sarrada) CreateSarrada() error {
	sarrada.Id = bson.NewObjectId()
	c := dbs.Session.DB(dbs.Database).C("sarradas")

	if err := c.Insert(sarrada); err != nil {
		pretty.Log(err)
		return err
	}
	
	return nil
}
