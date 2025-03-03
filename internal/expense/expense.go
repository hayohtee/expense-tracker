// Package expense provides functionality to manage and track expenses.
package expense

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
	"slices"
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
// Returns an error if the provided position is out of range.
//
// Parameters:
//   - pos: The 1-based position of the expense to be updated.
//   - description: The new description for the expense item. If empty, the description is not updated.
//   - amount: The new amount for the expense item. If negative, the amount is not updated.
//
// Returns:
//   - error: An error if the provided position is out of range, otherwise nil.
func (e *ExpenseList) Update(pos int, description string, amount float64) error {
	expenseList := *e

	// Check if the provided id is within the range.
	if pos < 1 || pos > len(expenseList) {
		return errors.New("invalid position: position is out of range")
	}

	// Update description only if new non-empty description is provided.
	if description != "" && expenseList[pos-1].Description != strings.ToLower(description) {
		expenseList[pos-1].Description = strings.ToLower(description)
		expenseList[pos-1].Date = time.Now()
	}

	// Update amount only if new non-negative amount is provided.
	if amount >= 0 && expenseList[pos-1].Amount != amount {
		expenseList[pos-1].Amount = amount
		expenseList[pos-1].Date = time.Now()
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

	*e = slices.Delete(expenseList, pos-1, pos)
	return nil
}

// List writes the expense list to the provided io.Writer in a tabular format.
func (e *ExpenseList) List(w io.Writer) {
	header := "ID    Date        Description   Amount\n"
	var buf bytes.Buffer
	buf.WriteString(header)

	for _, item := range *e {
		value := fmt.Sprintf("%6d%12s%14s$%.2f\n", item.ID, item.Date.Format("2006-01-02"), item.Description, item.Amount)
		buf.WriteString(value)
	}

	w.Write(buf.Bytes())
}

// Summary writes a summary of the total expenses to the provided io.Writer.
// It calculates the total amount of all expenses in the ExpenseList and
// writes the formatted summary string to the output writer.
//
// Parameters:
//
//	w (io.Writer): The writer to which the summary will be written.
func (e *ExpenseList) Summary(w io.Writer) {
	var total float64 = 0

	for _, item := range *e {
		total += item.Amount
	}

	summary := fmt.Sprintf("Total expenses: $%.2f\n", total)
	w.Write([]byte(summary))
}

// SummaryForMonth writes a summary of the total expenses for a given month to the provided writer.
// The month parameter should be an integer between 1 and 12, representing the months January to December.
// If the month is out of range, an error is returned.
// The summary includes the total amount of expenses for the specified month.
//
// Parameters:
//   w - an io.Writer where the summary will be written
//   month - an integer representing the month (1 for January, 12 for December)
//
// Returns:
//   error - an error if the month is out of range, otherwise nil
func (e *ExpenseList) SummaryForMonth(w io.Writer, month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month: month is out of range")
	}

	var total float64 = 0
	for _, item := range *e {
		if item.Date.Month() == time.Month(month) {
			total += item.Amount
		}
	}

	summary := fmt.Sprintf("Total expenses for %s: $%.2f\n", time.Month(month).String(), total)
	w.Write([]byte(summary))
	return nil
}
