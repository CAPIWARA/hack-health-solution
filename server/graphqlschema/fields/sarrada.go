package fields

import (
	"github.com/graphql-go/graphql"
	"hack-health-solution/server/graphqlschema/types"
	"hack-health-solution/server/graphqlschema/resolvers"
)

var CreateSarrada = &graphql.Field{
	Type:        types.TotalSarradinhas,
	Description: "create new sarrada",
	Args: graphql.FieldConfigArgument{
		"camisinha": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"oral": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"drogas": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"ejaculou": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Boolean),
		},
		"pessoa": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"quantidade": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve:     resolvers.CreateSarrada,
}

var SarradaQuery = &graphql.Field{
	Type:        types.TotalSarradinhas,
	Description: "get sarrada",
	Args: graphql.FieldConfigArgument{
		"sarradaId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: resolvers.GetSarrada,
}

var SarradasQuery = &graphql.Field{
	Type:        graphql.NewList(types.Sarrada),
	Description: "get sarradas",
	Resolve: resolvers.GetSarradas,
}
