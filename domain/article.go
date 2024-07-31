package domain

import baseEntities "test-backend-developer-sagala/common/base-entities"

type Article struct {
	baseEntities.Base
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}
