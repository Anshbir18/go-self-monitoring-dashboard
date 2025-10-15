
package models

import "time"

type Order struct {
	ID        int
	UserID    int
	Amount    float64
	Status    string
	CreatedAt time.Time
}

// Order{ID int, UserID int, Amount float64, Status string, CreatedAt time.Time}