package domain

import baseEntities "test-backend-developer-sagala/common/base-entities"

type User struct {
	baseEntities.Base
	Username string `json:"username"`
	Password string `json:"-"`
}
