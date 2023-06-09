package models

import "time"

//Consultation Model
type Consultation struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Date        string
	Time        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

//User Model
type Users struct {
	ID          int
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Password    string
	AccessLevel int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
