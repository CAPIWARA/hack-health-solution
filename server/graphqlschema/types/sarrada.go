package types

import "github.com/graphql-go/graphql"

var Sarrada = graphql.NewObject(graphql.ObjectConfig{
	Name: "sarrada",
	Fields: graphql.Fields{
		"pessoa": &graphql.Field{
			Type:        graphql.String,
			Description: "pessoa string",
		},
		"camisinha": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "camisinha bool",
		},
		"oral": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "oral bool",
		},
		"quantidade": &graphql.Field{
			Type:        graphql.Int,
			Description: "quantidade int",
		},
		"drogas": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "drogas bool",
		},
		"ejaculou": &graphql.Field{
			Type:        graphql.Boolean,
			Description: "ejaculou bool",
		},
		"userid": &graphql.Field{
			Type:        graphql.String,
			Description: "userid string",
		},
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id string",
		},
		"total": &graphql.Field{
			Type:        graphql.Int,
			Description: "total int",
		},
		"mensagem": &graphql.Field{
			Type:        graphql.String,
			Description: "mensagem string",
		},
	}})

var TotalSarradinhas = graphql.NewObject(graphql.ObjectConfig{
	Name: "sarradinhas",
	Fields: graphql.Fields{
		"total": &graphql.Field{
			Type:        graphql.Int,
			Description: "total de sarradinhas",
		},
		"mensagem": &graphql.Field{
			Type:        graphql.String,
			Description: "mensagem da sarrada",
		},
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id",
		},
	}})
