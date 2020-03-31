package handlers

import (
	"time"

	"katea_blog/conf"
	jwt "github.com/dgrijalva/jwt-go"
)

//Token the jwt custon claims
type Token struct {
	UserName string `json:"user_name"`
	Level    int    `json:"level"`
	UserID   string `json:"user_id"`
	jwt.StandardClaims
}

// Get ...
func (t *Token) Get() string {

	t.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	s, _ := conf.Config.String("secretkey")

	// Generate encoded token and send it as response.
	ts, _ := token.SignedString([]byte(s))
	return ts
}
