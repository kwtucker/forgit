package lib

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func isPath(p string) bool {
	switch string(p[0]) {
	case "/":
		return true
	}
	return false
}

// Init ...
func Init() {
	var (
		err    error
		reader *bufio.Reader
		path   string
	)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	reader = bufio.NewReader(os.Stdin)
	fmt.Println("Your Current Absolute Path is ->", currentDir)
	fmt.Println("-=-=-")
	fmt.Print("Enter Absolute Path to your Forgit Directory: ")

	path, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	valid := isPath(path)
	if !valid {
		fmt.Println()
		fmt.Println()
		fmt.Println("NOT VALID PATH. Absolute Path ONLY")
		fmt.Println("Example: /Users/kevintucker/Desktop/")
		fmt.Println("Suggest going to the Forgit Directory and running pwd command to get its path")
		fmt.Println("[ or ]")
		fmt.Println("Run Again --> fgt init")
		return
	}
	fmt.Println(path)
	BuildConfig(path)
}
