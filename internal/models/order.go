package models

import "time"

type Order struct {
	Id         int
	CustomerId int
	CreatedAt  time.Time
	Status     string
}
