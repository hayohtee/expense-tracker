# expense-tracker
A simple CLI expense tracker application to manage your finances

## Features
- Adding an expense with a description and amount.
- Updating an expense.
- Deleting an expense.
- Listing all expenses.
- Summary of all expenses.
- Summary of expenses for a specific month (of current year).

## Installing
Ensure the GO SDK is installed
```bash
go install github.com/hayohtee/expense-tracker@latest
```

## Building from source
Ensure the GO SDK is installed
1. Clone the repository
   ```bash
   git clone git@github.com:hayohtee/expense-tracker.git
   ```
3. Change into the project directory
   ```bash
   cd expense-tracker
   ```
4. Compile
   ```bash
   go build -o expense-tracker main.go
   ```

## Usage
```bash
$ expense-tracker add --description "Lunch" --amount 20
# Expense added successfully (ID: 1)

$ expense-tracker add --description "Dinner" --amount 10
# Expense added successfully (ID: 2)

$ expense-tracker list
# ID  Date        Description  Amount
# 1   2024-08-06  Lunch        $20
# 2   2024-08-06  Dinner       $10

$ expense-tracker summary
# Total expenses: $30

$ expense-tracker delete --id 2
# Expense deleted successfully

$ expense-tracker summary
# Total expenses: $20

$ expense-tracker summary --month 8
# Total expenses for August: $20
```

### Challenge URL
Solution to the [Task Tracker](https://roadmap.sh/projects/expense-tracker) project on [roadmap.sh](https://roadmap.sh)

