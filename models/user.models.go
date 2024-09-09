package models

import (
	"time"
)

type User struct {
	ID             string    `json:"id"`
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	DOB            time.Time `json:"dob" validate:"required"`
	Mobile         string    `json:"mobile" validate:"required"`
	Email          string    `json:"email" validate:"required,email"`
	Username       string    `json:"username" validate:"required"`
	PasswordHashed string    `json:"-" validate:"required"`
	WorkUploaded   int       `json:"work_uploaded"`
	WorkDone       int       `json:"work_done"`
	WorkScore      int       `json:"work_score"`
	TotalWorkScore int       `json:"total_work_score"`
	Bio            string    `json:"bio,omitempty"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Gender         string    `json:"gender"`
	Links          []string  `json:"links,omitempty"`
}

type UserUpdate struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DOB            time.Time `json:"dob"`
	Mobile         string    `json:"mobile"`
	Email          string    `json:"email"`
	Username       string    `json:"username"`
	PasswordHashed string    `json:"-"`
	Bio            string    `json:"bio"`
	ProfilePicture string    `json:"profile_picture"`
	Links          []string  `json:"links"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	FirstName      string    `json:"first_name" validate:"required"`
	LastName       string    `json:"last_name" validate:"required"`
	DOB            time.Time `json:"dob" validate:"required"`
	Mobile         string    `json:"mobile" validate:"required"`
	Email          string    `json:"email" validate:"required,email"`
	Username       string    `json:"username" validate:"required"`
	PasswordHash   string    `json:"password" validate:"required"`
	Bio            string    `json:"bio"`
	ProfilePicture string    `json:"profile_picture"`
	Links          []string  `json:"links"`
	Gender         string    `json:"gender"`
}
