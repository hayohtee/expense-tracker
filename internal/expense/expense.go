// Package expense provides functionality to manage and track expenses.
package expense

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
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

// Add adds a new expense to the ExpenseList with the given description and amount.
// It returns an error if the description is empty or the amount is negative.
//
// Parameters:
//   - description: A string representing the description of the expense.
//   - amount: A float64 representing the amount of the expense.
//
// Returns:
//   - error: An error if the description is empty or the amount is negative, otherwise nil.
func (e *ExpenseList) Add(description string, amount float64) error {
	if description == "" {
		return errors.New("description is empty")
	}

	if amount < 0 {
		return errors.New("negative amount")
	}

	expenseList := *e

	// Calculate the next id by adding the last expense id + 1
	var id int
	if len(expenseList) == 0 {
		id = 1
	} else {
		id = expenseList[len(expenseList)-1].ID + 1
	}

	item := expense{
		ID:          id,
		Date:        time.Now(),
		Description: strings.ToLower(description),
		Amount:      amount,
	}

	*e = append(*e, item)
	return nil
}

// Update modifies the description and/or amount of an expense item in the ExpenseList.
// The expense item to be updated is identified by its id (1-based index).
// If a new non-empty description is provided, it updates the description and sets the current date and time.
// If a new non-negative amount is provided, it updates the amount and sets the current date and time.
// Returns an error if the provided id is out of range.
//
// Parameters:
//   - id: The 1-based index of the expense item to be updated.
//   - description: The new description for the expense item. If empty, the description is not updated.
//   - amount: The new amount for the expense item. If negative, the amount is not updated.
//
// Returns:
//   - error: An error if the provided id is out of range, otherwise nil.
func (e *ExpenseList) Update(id int, description string, amount float64) error {
	expenseList := *e

	// Check if the provided id is within the range.
	if id < 1 || id > len(expenseList) {
		return errors.New("invalid id")
	}

	// Update description only if new non-empty description is provided.
	if description != "" && expenseList[id-1].Description != strings.ToLower(description) {
		expenseList[id-1].Description = strings.ToLower(description)
		expenseList[id-1].Date = time.Now()
	}

	// Update amount only if new non-negative amount is provided.
	if amount >= 0 && expenseList[id-1].Amount != amount {
		expenseList[id-1].Amount = amount
		expenseList[id-1].Date = time.Now()
	}

	return nil
}

// Delete removes an expense from the ExpenseList at the specified position.
// The position is 1-based, meaning the first element is at position 1.
// If the position is out of range, it returns an error indicating an invalid position.
//
// Parameters:
// - pos: The 1-based position of the expense to be removed.
//
// Returns:
// - error: An error if the position is out of range, otherwise nil.
func (e *ExpenseList) Delete(pos int) error {
	expenseList := *e
	if pos < 1 || pos > len(expenseList) {
		return errors.New("invalid position: position is out of range")
	}

	*e = append(expenseList[:pos-1], expenseList[pos:]...)
	return nil
}
