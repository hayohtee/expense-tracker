// Package expense provides functionality to manage and track expenses.
package expense

import "time"

// expense represents a single expense entry with an ID, date, description, and amount.
type expense struct {
	ID          int       `json:"id"`          // Unique identifier for the expense
	Date        time.Time `json:"date"`        // Date when the expense was incurred
	Description string    `json:"description"` // Description of the expense
	Amount      float64   `json:"amount"`      // Amount of the expense
}
