package Models

import (
	"time"
	"os"

	"github.com/kabos0809/redis_jwt_gin/Config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
)

func (m Model) CreateUser(username string, password string) error {
	tx := m.Db.Begin()
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err := tx.Create(&User{UName: username, Pass: string(hash)}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (m Model) CheckPassword(username string, password string) error {
	var user User
	m.Db.Find(&User{}, "uname = ?", username).Scan(&user)
	if err := bcrypt.CompareHashAndPassword([]byte(user.Pass), []byte(password)); err != nil {
		return err
	}
	return nil
}

func CreateJWTToken(username string, id uint64) (map[string]string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	Claims := token.Claims.(jwt.MapClaims)
	Claims["sub"] = "AccessToken"
	Claims["id"] = id
	Claims["uName"] = username
	Claims["iat"] = time.Now()
	Claims["exp"] = time.Now().Add(3 * time.Hour).Unix()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	ACCESS_TOKEN_SECRETKEY := os.Getenv("ACCESS_TOKEN_SECRETKEY")
	t, err := token.SignedString([]byte(ACCESS_TOKEN_SECRETKEY))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = "RefreshToken"
	rtClaims["id"] = id
	rtClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	REFRESH_TOKEN_SECRETKEY := os.Getenv("REFRESH_TOKEN_SECRETKEY")
	rt, err := token.SignedString([]byte(REFRESH_TOKEN_SECRETKEY))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"AccessToken": t,
		"RefreshToken": rt,
	}, nil
}

func (m Model) RefreshToken(id uint64, token string, exp int64) (map[string]string, error) {
	var user User
	m.Db.First(&user, id).Scan(&user)
	t, err := CreateJWTToken(user.UName, id) 
	if err != nil {
		return nil, err
	}

	if err := SetBlackList(token, exp); err != nil {
		return nil, err
	}
	return t, nil
}

func SetBlackList(token string, exp int64) error {
	c := Config.ConnRedis()
	defer c.Close()

	nowTime := time.Now()
	expTime := time.Unix(int64(exp), 0)

	timeSub := expTime.Sub(nowTime).Seconds()

	_, err := c.Do("SET", token, byte(exp))
	_, err = c.Do("EXPIRE", token, int64(timeSub))
	if err != nil {
		return err
	}
	return nil
}

func CheckBlackList(token string) error {
	c := Config.ConnRedis()
	defer c.Close()

	_, err := redis.String(c.Do("GET", token))
	if err != nil {
		return err
	}
	return nil
}