package graphqlschema

import (
	"comunicode/server/gql/fields"

	"github.com/graphql-go/graphql"
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mutation",
		Description: "mutation",
		Fields: graphql.Fields{
			"checkout": fields.CheckoutMutation,
		},
	})
