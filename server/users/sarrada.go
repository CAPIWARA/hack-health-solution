package users

import (
	"hack-health-solution/server/dbs"
	"github.com/kr/pretty"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type Sarrada struct {
	Id         string `json:"id" bson:"_id"`
	Stringid   string `json:"id"`
	UserId     string `json:"userid"`
	Camisinha  bool   `json:"camisinha"`
	Oral       bool   `json:"oral"`
	Pessoa     string `json:"pessoa"`
	Quantidade int    `json:"quantidade"`
	Drogas     bool   `json:"drogas"`
	Ejaculou   bool   `json:"ejaculou"`
	Mensagem   string `json:"mensagem"`
	Total      int    `json:"total"`
}

func (sarrada *Sarrada) CreateSarrada() (int, error) {
	sarrada.Id = bson.NewObjectId().Hex()
	c := dbs.Session.DB(dbs.Database).C("sarradas")

	total, err := sarrada.CalcSarrada()
	sarrada.Total = total
	if err != nil {
		pretty.Log(err)
		return 0, err
	}
	var mensagem string
	if total > 27 {
		mensagem = "Deus da sarrada"
	}
	if total > 17 && total < 28{
		mensagem = "Sarrou gostoso"
	}
	if total < 18 {
		mensagem = "Mandou mal"
	}
	sarrada.Mensagem = mensagem
	if err := c.Insert(sarrada); err != nil {
		pretty.Log(err)
		return 0, err
	}

	return total, nil
}

func GetSarrada(id string) (*Sarrada, error) {
	var data Sarrada
	pretty.Log("sarrada id: ", id)
	c := dbs.Session.DB(dbs.Database).C("sarradas")
	if err := c.Find(bson.M{"_id": id}).One(&data); err != nil {
		pretty.Log("error: ", err)
		return nil, errors.New("sarrada not found")
	}
	pretty.Log(data)
	return &data, nil
}

func GetSarradas(id string) ([]Sarrada, error) {
	var data []Sarrada
	c := dbs.Session.DB(dbs.Database).C("sarradas")
	if err := c.Find(bson.M{"userid": id}).All(&data); err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	for k, v := range data {
		data[k].Id = bson.ObjectIdHex(v.Id).Hex()
	}
	return data, nil
}

func (sarrada *Sarrada) CalcSarrada() (int, error) {
	var total int
	total = total + (sarrada.Quantidade)
	if sarrada.Drogas {
		total = total - 2
	}
	if sarrada.Oral {
		total = total + 4
	}
	if sarrada.Pessoa == "FIXO" {
		total = total + 5
	}
	if sarrada.Pessoa == "RECORRENTE" {
		total = total + 4
	}
	if sarrada.Pessoa == "UMA NOITE" {
		total = total + 2
	}
	if sarrada.Pessoa == "PROFISSIONAL" {
		total = total - 1
	}
	if sarrada.Camisinha {
		if sarrada.Ejaculou {
			total = total + 5
		}
		total = total * 2
	}
	total = total + (5 - sarrada.Quantidade)
	return total, AddSarradinhas(total, sarrada.UserId)
}
