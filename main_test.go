package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

const testFileName = ".expense_tracker_test.json"

var binName = "expense-tracker"

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	// Append .exe extension to binName if the operating
	// system that is running the test is windows.
	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName, "main.go")
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool %s:%s\n", binName, err.Error())
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	os.Remove(binName)
	os.Remove(filename)

	os.Exit(result)
}

func TestExpenseTracker(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	expenses := []struct {
		description string
		amount      float64
	}{
		{description: "new expense 1", amount: 20.00},
		{description: "new expense 2", amount: 20.60},
		{description: "new expense 3", amount: 200.50},
		{description: "new expense 4", amount: 200.00},
		{description: "new expense 5", amount: 120.00},
	}

	t.Run("TestAddCMD", func(t *testing.T) {
		for index, item := range expenses {
			cmd := exec.Command(cmdPath, "add", "--description", item.description, "--amount", fmt.Sprint(item.amount))
			out, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatal(err)
			}

			expected := fmt.Sprintf("Expense added successfully (ID: %d)\n", index + 1)
			if string(out) != expected {
				t.Errorf("expected %q, but got %q instead", expected, string(out))
			}
		}
	})

}
