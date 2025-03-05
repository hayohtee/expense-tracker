package main

import (
	"fmt"
	"os"
	"os/exec"
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

	build := exec.Command("go", "build", "-o", binName, "./...")
	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool %s:%s", binName, err.Error())
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	os.Remove(binName)
	os.Remove(filename)

	os.Exit(result)
}
