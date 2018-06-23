package types

import "github.com/graphql-go/graphql"

var User = graphql.NewObject(graphql.ObjectConfig{
	Name: "user",
	Fields: graphql.Fields{
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
		"points": &graphql.Field{
			Type:        graphql.String,
			Description: "points",
		},
		"socialName": &graphql.Field{
			Type:        graphql.String,
			Description: "socialName",
		},
		"gender": &graphql.Field{
			Type:        graphql.String,
			Description: "gender",
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

	}})