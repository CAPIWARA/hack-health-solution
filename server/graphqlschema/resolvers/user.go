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

func CreateExam(params graphql.ResolveParams) (interface{}, error) {
	userId := params.Context.Value("id").(string)

	user, err := users.GetUser(userId)
	if err != nil {
		pretty.Log(err)
		return nil, err
	}

	lastexam, err := users.ConvertStringToDate(user.LastExam)
	if err != nil {
		pretty.Log(err)
		return nil, err
	}
	if lastexam.AddDate(0, 3, 0).After(time.Now().AddDate(0, 3, 0).Add(-1 * time.Minute)) {
		return nil, errors.New("Voce precisa esperar pelo menos 3 meses para refazer o teste")
	}

	if err := user.CreateExam(); err != nil {
		pretty.Log(err)
		return nil, err
	}

	return user, nil
}

func GetFriends(params graphql.ResolveParams) (interface{}, error) {
	userId := params.Context.Value("id").(string)

	user, err := users.GetUser(userId)

	if err != nil {
		pretty.Log(err)
		return nil, err
	}
	var friends []users.Friend
	for _, v := range user.Friends {
		item, _ := users.GetUser(v)
		friend := users.Friend{
			Name:        item.Name,
			Sarradinhas: item.Sarradinhas,
			Image: item.Image,
		}
		friends = append(friends, friend)
	}
	return friends, nil
}
