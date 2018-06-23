package graphqlschema

import (

	"github.com/graphql-go/graphql"
	"hack-health-solution/server/graphqlschema/fields"
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mutation",
		Description: "mutation",
		Fields: graphql.Fields{
			"createSarrada" : fields.CreateSarrada,
		},
	})
