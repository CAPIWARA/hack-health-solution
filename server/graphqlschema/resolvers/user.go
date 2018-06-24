package resolvers

import (

	"github.com/graphql-go/graphql"
	"hack-health-solution/server/users"
	"github.com/kr/pretty"
	"time"
	"github.com/pkg/errors"
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

func CreateExam(params graphql.ResolveParams)(interface{}, error){
	userId := params.Context.Value("id").(string)

	user, err := users.GetUser(userId)
	if err != nil {
		pretty.Log(err)
		return nil, err
	}

	lastexam, err := users.ConvertStringToDate(user.LastExam)
	if err != nil {
		pretty.Log(err)
	}
	if lastexam.AddDate(0,3,0).After(time.Now().AddDate(0,3,0).Add(-2 * time.Minute)) {
		return nil, errors.New("Voce precisa esperar pelo menos 3 meses para refazer o teste")
	}

	if err := user.CreateExam(); err != nil {
		pretty.Log(err)
		return nil, err
	}

	return user, nil
}