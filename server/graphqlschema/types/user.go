package types

import (
	"github.com/graphql-go/graphql"
)

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.String,
			Description: "id",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "name",
		},
		"email": &graphql.Field{
			Type:        graphql.String,
			Description: "email",
		},
		"birthday": &graphql.Field{
			Type:        graphql.String,
			Description: "birthday",
		},
		"sarradinhas": &graphql.Field{
			Type:        graphql.String,
			Description: "sarradinhas",
		},
		"socialName": &graphql.Field{
			Type:        graphql.String,
			Description: "socialName",
		},
		"orientation": &graphql.Field{
			Type:        graphql.String,
			Description: "orientation",
		},
		"ethnicity": &graphql.Field{
			Type:        graphql.String,
			Description: "ethnicity",
		},
		"genderIdentity": &graphql.Field{
			Type:        graphql.String,
			Description: "genderIdentity",
		},
		"image": &graphql.Field{
			Type:        graphql.String,
			Description: "image",
		},
		"lastExam": &graphql.Field{
			Type:        graphql.String,
			Description: "lastExam",
		},
	}})

var Friends = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "name",
		},
		"sarradinhas": &graphql.Field{
			Type:        graphql.Int,
			Description: "sarradinhas",
		},
	}})
