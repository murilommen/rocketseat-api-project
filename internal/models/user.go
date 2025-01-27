package models

import "github.com/google/uuid"

type UserId uuid.UUID

type User struct {
	FirstName string
	LastName string
	Biography string
}

type UserResponse struct {
	Id string
	User User
}