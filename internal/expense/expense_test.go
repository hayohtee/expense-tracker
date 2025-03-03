package expense_test

import (
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
			expected:    fmt.Sprintf("%d\t%s\t%s\t%.2f", 1, "demo expense 1", time.Now().Format("2006-01-02"), 100.0),
		},

		{
			description: "Demo Expense 2",
			amount:      150,
			expected:    fmt.Sprintf("%d\t%s\t%s\t%.2f", 2, "demo expense 2", time.Now().Format("2006-01-02"), 150.0),
		},
		{
			description: "Demo Expense 3",
			amount:      250,
			expected:    fmt.Sprintf("%d\t%s\t%s\t%.2f", 3, "demo expense 3", time.Now().Format("2006-01-02"), 250.0),
		},
		{
			description: "Demo Expense 4",
			amount:      1150,
			expected:    fmt.Sprintf("%d\t%s\t%s\t%.2f", 4, "demo expense 4", time.Now().Format("2006-01-02"), 1150.0),
		},
		{
			description: "Demo Expense 5",
			amount:      50.90,
			expected:    fmt.Sprintf("%d\t%s\t%s\t%.2f", 5, "demo expense 5", time.Now().Format("2006-01-02"), 50.90),
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
