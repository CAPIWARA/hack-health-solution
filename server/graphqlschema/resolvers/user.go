package resolvers

import (

	"github.com/graphql-go/graphql"
	"hack-health-solution/server/users"
	"github.com/kr/pretty"
)

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	token := params.Context.Value("id").(string)
	user, err := users.GetUser(token)
	if err != nil {
		pretty.Log("error: ", err)
		return nil, err
	}
	return user, nil
}

