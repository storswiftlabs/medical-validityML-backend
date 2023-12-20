package handlers

import (
	"medical-zkml-backend/internal/db"
	"medical-zkml-backend/internal/module"
)

type User struct{}

func (u *User) IsRegistered(user string) (bool, error) {
	return db.UserQuery(user)
}

func (u *User) UserRegistration(user string) error {
	return db.UserRegistration(user)
}

func (u *User) IsCollected(user, url string) (bool, error) {
	return db.CheckCollection(user, url)
}

func (u *User) CollectArticles(user, url string) error {
	return db.CollectArticles(user, url)

}

func (u *User) FavoriteArticleList(user string) ([]module.ArticleResult, error) {
	return db.ArticlesForUser(user)
}

func (u *User) CancelArticleCollection(user, url string) error {
	return db.DeleteCollectArticles(user, url)
}
