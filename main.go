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
}
