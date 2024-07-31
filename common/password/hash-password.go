package password

import (
	"golang.org/x/crypto/bcrypt"
	"test-backend-developer-sagala/config/env"
)

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), env.GlobalConfiguration.CostPassword)
	if err != nil {
		return string(hashed), err
	}
	return string(hashed), err
}
