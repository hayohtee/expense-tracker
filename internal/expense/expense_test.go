package expense_test

import (
	"os"
	"testing"

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
