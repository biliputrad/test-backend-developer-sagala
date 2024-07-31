package jwt_token

import (
	"github.com/dgrijalva/jwt-go"
	"test-backend-developer-sagala/config/env"
	"time"
)

type JwtClaim struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

// GenerateToken is used to create new token and will return token and time expired token
func GenerateToken(claim JwtClaim) (string, time.Time, error) {
	jwtKey := []byte(env.GlobalConfiguration.SecretKey)
	expirationTime := time.Now().Add(time.Duration(env.GlobalConfiguration.ExpiredTime) * time.Hour).Local()

	claims := claim
	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	return tokenString, expirationTime, err
}
