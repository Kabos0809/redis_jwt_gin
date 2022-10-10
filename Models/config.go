package Models

import "gorm.io/gorm"

type ModelInterface interface {
	CreateUser(username string, password string) error
	CheckPassword(username string, password string) error
	RefreshToken(userid uint64, token string, exp int64) (map[string]string, error)
}

type Model struct {
	Db *gorm.DB
}