package jwt_token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"test-backend-developer-sagala/common/constants"
	responseMessage "test-backend-developer-sagala/common/response-message"
	"test-backend-developer-sagala/config/database/postgres"
	"test-backend-developer-sagala/config/env"
	"test-backend-developer-sagala/domain"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		bearerToken := context.GetHeader(constants.Authorization)
		if !strings.Contains(bearerToken, constants.Bearer) {
			res := responseMessage.GetResponse(http.StatusUnauthorized, false, constants.ResponseInvalidToken, nil)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		tokenString := ""
		arrayToken := strings.Split(bearerToken, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		id, err := ValidateUserToken(tokenString)
		if err != nil {
			res := responseMessage.GetResponse(http.StatusUnauthorized, false, err.Error(), nil)
			context.JSON(http.StatusUnauthorized, res)
			context.Abort()
			return
		}

		context.Set("id", id)
		context.Next()
	}
}

func ValidateUserToken(signedToken string) (id int64, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(env.GlobalConfiguration.SecretKey), nil
		},
	)
	if err != nil {
		return id, err
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New(constants.MsgParseErr)
		return id, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New(constants.MsgTokenExpired)
		return id, err
	}

	return ValidateUserAccount(claims.ID, claims.Username)
}

func ValidateUserAccount(userId int64, username string) (int64, error) {
	var user domain.User
	postgres.GlobalDatabase.Where(constants.ById, userId).First(&user)
	if user.ID == 0 {
		return 0, errors.New(constants.ResponseInvalidToken)
	}

	if user.Username != username {
		return 0, errors.New(constants.ResponseInvalidToken)
	}

	return user.ID, nil
}
