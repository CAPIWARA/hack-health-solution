package resolvers

import (
	"github.com/graphql-go/graphql"
	"hack-health-solution/server/users"
	"github.com/kr/pretty"
)

func CreateSarrada(params graphql.ResolveParams) (interface{}, error) {
	var sarrada users.Sarrada

	sarrada.Camisinha = params.Args["camisinha"].(bool)
	sarrada.Drogas = params.Args["drogas"].(bool)
	sarrada.Ejaculou = params.Args["ejaculou"].(bool)
	sarrada.Oral = params.Args["oral"].(bool)
	sarrada.Quantidade = params.Args["quantidade"].(int)
	sarrada.Pessoa = params.Args["pessoa"].(string)
	sarrada.UserId = params.Context.Value("id").(string)

	if err := sarrada.CreateSarrada(); err != nil {
		pretty.Log(err)
		return nil, err
	}
	return sarrada, nil
}