// Package expense provides functionality to manage and track expenses.
package expense

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

// expense represents a single expense entry with an ID, date, description, and amount.
type expense struct {
	ID          int       `json:"id"`          // Unique identifier for the expense
	Date        time.Time `json:"date"`        // Date when the expense was incurred
	Description string    `json:"description"` // Description of the expense
	Amount      float64   `json:"amount"`      // Amount of the expense
}

// ExpenseList represents a list of expenses.
type ExpenseList []expense

// Load reads expense data from the specified file and loads it into the ExpenseList.
// The filename parameter specifies the path to the file to be loaded.
// It returns an error if there is any issue reading or parsing the file.
func (e *ExpenseList) Load(filename string) error {
	// Read the contents of the file using os.ReadFile
	content, err := os.ReadFile(filename)
	if err != nil {
		switch {
		// Skip if the file does not exist.
		case errors.Is(err, os.ErrNotExist):
			return nil
		default:
			return err
		}
	}

	// Simply skip if the contents of the file is empty.
	if len(content) == 0 {
		return nil
	}

	// Parse the json contents into list of expense struct.
	return json.Unmarshal(content, e)
}

// Save serializes the ExpenseList to JSON format and writes it to the specified file.
// The JSON data is indented for readability.
// The file is created with read-write permissions for the owner and read-only permissions for others.
//
// Parameters:
//   - filename: The name of the file where the JSON data will be saved.
//
// Returns:
//   - error: An error if the JSON marshaling or file writing fails, otherwise nil.
func (e *ExpenseList) Save(filename string) error {
	js, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}
