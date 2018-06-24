package users

import (
	"time"
	"github.com/kr/pretty"
)

var location, _ = time.LoadLocation("America/Sao_Paulo")
var format = "2006-01-02 15:04:05"

func GetCurrentGmtDate() string {
	time := time.Now().In(location)
	return time.Format(format)
}

func ConvertStringToDate(strdate string) (time.Time, error) {
	deadline, err := time.Parse(format, strdate)
	if err != nil {
		pretty.Log(err)
		return time.Time{}, err
	}
	deadline = deadline.Add(3 * time.Hour).In(location)
	return deadline, nil
}
