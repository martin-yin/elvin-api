package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	NickName   string `json:"nick_name"`
	BufferTime int64  `json:"buffer_time"`
	jwt.StandardClaims
}
