package model

import "time"

type ExampleStatus string

const (
	ACTIVE   ExampleStatus = "active"
	INACTIVE ExampleStatus = "inactive"
	BANNED   ExampleStatus = "banned"
)

type Example struct {
	ID         int
	Name       string
	Email      string
	Phone      string
	Age        int
	Address    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Status     ExampleStatus
	IsVerified bool
}

func (e *Example) TableName() string {
	return "example"
}
