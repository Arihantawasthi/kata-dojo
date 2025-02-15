package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <function>")
        fmt.Println("Example: go run main.go Bubble")
        os.Exit(1)
    }

    function := fmt.Sprintf("Test%s", os.Args[1])
    testCmd := exec.Command("go", "test", "-v", "./tests/...", "-run", function)

    testCmd.Stdout = os.Stdout
    testCmd.Stdin = os.Stdin

    err := testCmd.Run()
    if err != nil {
        if strings.Contains(err.Error(), "exit status 1") {
            fmt.Println("Some tests failed xxx")
        } else {
            fmt.Printf("Error running tests: %v\n", err)
        }
        os.Exit(1)
    } else {
        fmt.Println("All tests passed!!!")
    }
}
