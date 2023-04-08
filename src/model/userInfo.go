package model

import (
	"golang.org/x/text/language"
)

type UserInfo struct {
	ID                string
	Username          string
	Password          string
	FirstName         string
	LastName          string
	Email             string
	EmailVerified     bool
	Phone             string
	PhoneVerified     bool
	PreferredLanguage language.Tag
	IsAdmin           bool
}

func NewUserInfo() map[string]*UserInfo {

	return map[string]*UserInfo{
		"id1": {
			ID:                "id1",
			Username:          "test-user1",
			Password:          "verysecure",
			FirstName:         "Test",
			LastName:          "User",
			Email:             "test-user1@gy.com",
			EmailVerified:     true,
			Phone:             "",
			PhoneVerified:     false,
			PreferredLanguage: language.Chinese,
			IsAdmin:           true,
		},
		"id2": {
			ID:                "id2",
			Username:          "test-user2",
			Password:          "verysecure",
			FirstName:         "Test",
			LastName:          "User2",
			Email:             "test-user2@gy.com",
			EmailVerified:     true,
			Phone:             "",
			PhoneVerified:     false,
			PreferredLanguage: language.Chinese,
			IsAdmin:           false,
		},
	}
}