package authentication_dto

import "time"

type LoginResult struct {
	Type        string    `json:"type"`
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}
