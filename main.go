package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hayohtee/expense-tracker/internal/expense"
)

const filename = ".expense_list.json"

func main() {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(0)
	}

	// Load the expense list from the file.
	var expenseList expense.ExpenseList
	if err := expenseList.Load(filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "add":
		description := addCmd.String("description", "", "The description for the expense")
		amount := addCmd.Float64("amount", 0.0, "The amount for the expense")
		if err := addCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Add new expense to the list.
		if err := expenseList.Add(*description, *amount); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Write successful message to the STDOUT.
		fmt.Printf("Expense added successfully (ID: %d)\n", expenseList[len(expenseList)-1].ID)

		// Save the new list.
		if err := expenseList.Save(filename); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case "list":
		if err := listCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		// Write the list of expense to the STDOUT.
		expenseList.List(os.Stdout)
	case "summary":
		month := summaryCmd.Int("month", 0, "The month to generate the summary for")
		if err := summaryCmd.Parse(os.Args[2:]); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// If month was not specified, simply generate the summary for all expenses
		// and write to the STDOUT.
		if *month == 0 {
			expenseList.Summary(os.Stdout)
			return
		}

		// If month was specified then generate summary for the provided month and
		// write to the STDOUT.
		if err := expenseList.SummaryForMonth(os.Stdout, *month); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
