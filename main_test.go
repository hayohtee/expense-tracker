package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
	"time"
)

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

			expected := fmt.Sprintf("Expense added successfully (ID: %d)\n", index+1)
			if string(out) != expected {
				t.Errorf("expected %q, but got %q instead", expected, string(out))
			}
		}
	})

	t.Run("TestSummaryCMD", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "summary")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		var totalExpenses float64 = 0

		for _, v := range expenses {
			totalExpenses += v.amount
		}

		expected := fmt.Sprintf("Total expenses: $%.2f\n", totalExpenses)
		if string(out) != expected {
			t.Errorf("expected %q, but got %q instead", expected, string(out))
		}

		currentMonth := time.Now().Month()
		cmd = exec.Command(cmdPath, "summary", "-month", fmt.Sprint(int(currentMonth)))
		out, err = cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		expected = fmt.Sprintf("Total expenses for %s: $%.2f\n", time.Month(currentMonth).String(), totalExpenses)
		if string(out) != expected {
			t.Errorf("expected %q, but got %q instead", expected, string(out))
		}
	})

	t.Run("TestUpdateCMD", func(t *testing.T) {
		cases := []struct {
			id          string
			amount      string
			description string
			expected    string
		}{
			{
				id:          "1",
				amount:      "100",
				description: "new expense 1",
				expected:    "Expense updated successfully (ID: 1)\n",
			},
			{
				id:       "4",
				amount:   "1000",
				expected: "Expense updated successfully (ID: 4)\n",
			},
			{
				id:          "5",
				description: "new expense 5",
				expected:    "Expense updated successfully (ID: 5)\n",
			},
		}

		for _, v := range cases {
			var cmd *exec.Cmd
			switch {
			case v.amount != "" && v.description != "":
				cmd = exec.Command(cmdPath, "update", "--id", v.id, "--amount", v.amount, "--description", v.description)
			case v.amount != "":
				cmd = exec.Command(cmdPath, "update", "--id", v.id, "--amount", v.amount)
			case v.description != "":
				cmd = exec.Command(cmdPath, "update", "--id", v.id, "--description", v.description)
			}
			out, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println(err)
				t.Fatal(err)
			}

			if string(out) != v.expected {
				t.Errorf("expected %q, but got %q instead", v.expected, string(out))
			}
		}

	})

	t.Run("TestDeleteCMD", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "delete", "-id", "5")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := "Expense deleted successfully\n"
		if string(out) != expected {
			t.Errorf("expected %q, but got %q instead", expected, string(out))
		}

		cmd = exec.Command(cmdPath, "delete", "-id", "4")
		out, err = cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != expected {
			t.Errorf("expected %q, but got %q instead", expected, string(out))
		}

		cmd = exec.Command(cmdPath, "delete", "-id", "1")
		out, err = cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		if string(out) != expected {
			t.Errorf("expected %q, but got %q instead", expected, string(out))
		}

	})

}
