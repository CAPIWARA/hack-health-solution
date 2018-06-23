package graphqlschema

import (

	"github.com/graphql-go/graphql"
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mutation",
		Description: "mutation",
		Fields: graphql.Fields{},
	})
