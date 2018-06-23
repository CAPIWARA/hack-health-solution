package graphqlschema

import (

	"github.com/graphql-go/graphql"
	"hack-health-solution/server/graphqlschema/fields"
)

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "query",
		Description: "query",
		Fields: graphql.Fields{
			"user": fields.UserQuery,
			"sarrada": fields.SarradaQuery,
			"sarradas": fields.SarradasQuery,
		},
	},
)
