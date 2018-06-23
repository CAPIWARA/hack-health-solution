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
	}})