package db

import (
	"medical-zkml-backend/pkg/config"
	"testing"
)

func init() {
	InitMysql(config.NewConfig())
}

func TestInitMysql(t *testing.T) {
	conf := config.NewConfig()
	InitMysql(conf)

}

func TestUserQuery(t *testing.T) {
	query, err := UserQuery("0x01f5BB073e893d334FF9b0e239939982c124AF97")
	if err != nil {
		t.Error(err)
	}
	t.Log(query)
}

func TestGetArticle(t *testing.T) {
	articles, err := GetArticle([]string{"Acute_Inflammations"})
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(articles)
}
