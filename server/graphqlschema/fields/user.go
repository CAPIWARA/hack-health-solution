package fields

import (

	"github.com/graphql-go/graphql"
	"hack-health-solution/server/graphqlschema/types"
	"hack-health-solution/server/graphqlschema/resolvers"

)

var UserQuery = &graphql.Field{
	Type:        types.User,
	Description: "get user",
	Resolve:     resolvers.GetUser,
}

var CreateExam = &graphql.Field{
	Type:        types.User,
	Description: "create exam",
	Resolve:     resolvers.CreateExam,
}
