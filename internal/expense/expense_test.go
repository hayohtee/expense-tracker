package expense_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/hayohtee/expense-tracker/internal/expense"
)

func TestSaveAndLoad(t *testing.T) {
	// Create a temp file for holding expense list.
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	var expenseList expense.ExpenseList

	// Add new expense to the list.
	if err := expenseList.Add("Demo Expense", 50.55); err != nil {
		t.Fatal(err)
	}

	// Save the expense list into the tempFile.
	if err := expenseList.Save(tempFile.Name()); err != nil {
		t.Fatal(err)
	}

	var newExpenseList expense.ExpenseList

	// Read new expense list from the tempFile.
	if err := newExpenseList.Load(tempFile.Name()); err != nil {
		t.Fatal(err)
	}

	// Assert the new expense list contains the expense inside the tempFile.
	if len(newExpenseList) != 1 {
		t.Errorf("expected length of expense list %d, but got %d instead", 1, len(newExpenseList))
	}

	// Check if the contents of the expense are the same.
	expected := expenseList[0].String()
	if newExpenseList[0].String() != expected {
		t.Errorf("expected %q, but got %q instead", expected, newExpenseList[0].String())
	}
}

func TestAdd(t *testing.T) {
	var expenseList expense.ExpenseList

	testCases := []struct {
		description string
		amount      float64
		expected    string
	}{
		{
			description: "Demo Expense 1",
			amount:      100,
			expected:    fmt.Sprintf("%-6d%-14s%-25s$%.2f", 1, time.Now().Format("2006-01-02"), "demo expense 1", 100.0),
		},

		{
			description: "Demo Expense 2",
			amount:      150,
			expected:    fmt.Sprintf("%-6d%-14s%-25s$%.2f", 2, time.Now().Format("2006-01-02"), "demo expense 2", 150.0),
		},
		{
			description: "Demo Expense 3",
			amount:      250,
			expected:    fmt.Sprintf("%-6d%-14s%-25s$%.2f", 3, time.Now().Format("2006-01-02"), "demo expense 3", 250.0),
		},
		{
			description: "Demo Expense 4",
			amount:      1150,
			expected:    fmt.Sprintf("%-6d%-14s%-25s$%.2f", 4, time.Now().Format("2006-01-02"), "demo expense 4", 1150.0),
		},
		{
			description: "Demo Expense 5",
			amount:      50.90,
			expected:    fmt.Sprintf("%-6d%-14s%-25s$%.2f", 5, time.Now().Format("2006-01-02"), "demo expense 5", 50.90),
		},
	}

	for index, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if err := expenseList.Add(tc.description, tc.amount); err != nil {
				t.Fatal(err)
			}

			if expenseList[index].String() != tc.expected {
				t.Errorf("expected %q, but got %q instead", tc.expected, expenseList[index].String())
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	var expenseList expense.ExpenseList

	// Add some expenses.
	if err := expenseList.Add("Demo Expense 1", 100); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 2", 150); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 3", 150); err != nil {
		t.Fatal(err)
	}

	// Update the second expense item.
	if err := expenseList.Update(2, "New Demo Expense 2", 500); err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf("%-6d%-14s%-25s$%.2f", 2, time.Now().Format("2006-01-02"), "new demo expense 2", 500.0)

	// Assert the second expense item was updated successfully.
	if expenseList[1].String() != expected {
		t.Errorf("expected %q, but got %q instead", expected, expenseList[1].String())
	}
}

func TestDelete(t *testing.T) {
	var expenseList expense.ExpenseList

	// Add some expenses.
	if err := expenseList.Add("Demo Expense 1", 100); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 2", 150); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 3", 150); err != nil {
		t.Fatal(err)
	}

	// Delete the second expense item.
	if err := expenseList.Delete(2); err != nil {
		t.Fatal(err)
	}

	// Assert that the expense item has been deleted.
	if len(expenseList) != 2 {
		t.Errorf("expected length of the expense list: %d, but got %d instead", 2, len(expenseList))
	}

	// Assert the first is the same.
	expected := fmt.Sprintf("%-6d%-14s%-25s$%.2f", 1, time.Now().Format("2006-01-02"), "demo expense 1", 100.0)
	if expenseList[0].String() != expected {
		t.Errorf("expected %q, but got %q instead", expected, expenseList[0].String())
	}

	// Assert the last expense is the same.
	expected = fmt.Sprintf("%-6d%-14s%-25s$%.2f", 3, time.Now().Format("2006-01-02"), "demo expense 3", 150.0)
	if expenseList[1].String() != expected {
		t.Errorf("expected %q, but got %q instead", expected, expenseList[1].String())
	}

	// Delete the last expense item.
	if err := expenseList.Delete(2); err != nil {
		t.Fatal(err)
	}

	// Assert that the expense item has been deleted.
	if len(expenseList) != 1 {
		t.Errorf("expected length of the expense list: %d, but got %d instead", 1, len(expenseList))
	}

	// Delete the remaining expense item.
	if err := expenseList.Delete(1); err != nil {
		t.Fatal(err)
	}

	// Assert the expense list is empty.
	if len(expenseList) != 0 {
		t.Errorf("expected length of the expense list: %d, but got %d instead", 0, len(expenseList))
	}
}

func TestList(t *testing.T) {
	var expenseList expense.ExpenseList

	// Add some expenses.
	if err := expenseList.Add("Demo Expense 1", 100); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 2", 150); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 3", 150); err != nil {
		t.Fatal(err)
	}

	var expectedBuf bytes.Buffer
	expectedBuf.WriteString(fmt.Sprintf("%-6s%-14s%-25s%s\n", "ID", "Date", "Description", "Amount"))
	expectedBuf.WriteString(fmt.Sprintf("%-6d%-14s%-25s$%.2f\n", 1, time.Now().Format("2006-01-02"), "demo expense 1", 100.0))
	expectedBuf.WriteString(fmt.Sprintf("%-6d%-14s%-25s$%.2f\n", 2, time.Now().Format("2006-01-02"), "demo expense 2", 150.0))
	expectedBuf.WriteString(fmt.Sprintf("%-6d%-14s%-25s$%.2f\n", 3, time.Now().Format("2006-01-02"), "demo expense 3", 150.0))

	var gotBuf bytes.Buffer

	expenseList.List(&gotBuf)

	if expectedBuf.String() != gotBuf.String() {
		t.Errorf("expected %q\n, but got %q instead", expectedBuf.String(), gotBuf.String())
	}
}

func TestSummary(t *testing.T) {
	var expenseList expense.ExpenseList

	// Add some expenses.
	if err := expenseList.Add("Demo Expense 1", 100); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 2", 150); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 3", 150); err != nil {
		t.Fatal(err)
	}

	expected := fmt.Sprintf("Total expenses: $%.2f\n", (100.0 + 150.0 + 150.00))

	var buf bytes.Buffer
	expenseList.Summary(&buf)

	if expected != buf.String() {
		t.Errorf("expected %q, but got %q instead", expected, buf.String())
	}
}

func TestSummaryForMonth(t *testing.T) {
	var expenseList expense.ExpenseList
	
	// Add some expenses.
	if err := expenseList.Add("Demo Expense 1", 100); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 2", 150); err != nil {
		t.Fatal(err)
	}
	if err := expenseList.Add("Demo Expense 3", 150); err != nil {
		t.Fatal(err)
	}
	
	currentMonth := time.Now().Month()
	expected := fmt.Sprintf("Total expenses for %s: $%.2f\n", currentMonth.String(), (100.0 + 150.0 + 150.0))
	
	var buf bytes.Buffer
	expenseList.SummaryForMonth(&buf, int(currentMonth))
	
	if expected != buf.String() {
		t.Errorf("expected %q, but got %q instead", expected, buf.String())
	}
}