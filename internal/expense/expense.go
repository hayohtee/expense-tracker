// Package expense provides functionality to manage and track expenses.
package expense

import (
	"encoding/json"
	"os"
	"time"
)

// defaultFileName is the name of the file where the expense tracker data is stored.
const defaultFileName = ".expense_tracker.json"

// expense represents a single expense entry with an ID, date, description, and amount.
type expense struct {
	ID          int       `json:"id"`          // Unique identifier for the expense
	Date        time.Time `json:"date"`        // Date when the expense was incurred
	Description string    `json:"description"` // Description of the expense
	Amount      float64   `json:"amount"`      // Amount of the expense
}

// ExpenseList represents a list of expenses.
type ExpenseList []expense

// Load reads the expense data from a JSON file and decodes it into the ExpenseList.
// It returns an error if the file cannot be opened or if the decoding fails.
func (e *ExpenseList) Load() error {
	file, err := os.Open(defaultFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(e)
}
