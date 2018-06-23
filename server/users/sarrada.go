package users

import (
	"hack-health-solution/server/dbs"
	"github.com/kr/pretty"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Sarrada struct {
	Id       string `json:"-" bson:"_id"`
	UserId     string        `json:"userid"`
	Camisinha  bool          `json:"camisinha"`
	Oral       bool          `json:"oral"`
	Pessoa     string        `json:"pessoa"`
	Quantidade int           `json:"quantidade"`
	Drogas     bool          `json:"drogas"`
	Ejaculou   bool          `json:"ejaculou"`
}

func (sarrada *Sarrada) CreateSarrada() error {
	sarrada.Id = bson.NewObjectId().String()
	c := dbs.Session.DB(dbs.Database).C("sarradas")

	if err := c.Insert(sarrada); err != nil {
		pretty.Log(err)
		return err
	}
	return nil
}

func GetSarrada(id string) (*Sarrada, error) {
	var data Sarrada
	c := dbs.Session.DB(dbs.Database).C("sarradas")
	if err := c.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&data); err != nil {
		pretty.Log("error: ", err)
		return nil, errors.New("sarrada not found")
	}
	data.Id = bson.ObjectId(data.Id).Hex()
	pretty.Log(data)
	return &data, nil
}

func GetSarradas(id string)([]Sarrada, error){
	var data []Sarrada
	c := dbs.Session.DB(dbs.Database).C("sarradas")
	if err := c.Find(bson.M{"userid": id}).All(&data); err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	for k, v := range data {
		data[k].Id = bson.ObjectId(v.Id).Hex()
	}
	return data, nil
}
