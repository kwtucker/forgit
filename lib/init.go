package lib

import (
	"bufio"
	"fmt"
	"os"
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
		err     error
		scanner *bufio.Scanner
		path    string
	)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("<> Your Current Absolute Path is ->", currentDir)
	fmt.Println("<>")
	fmt.Print("<> Enter Absolute Path to your Forgit Directory: ")
	scanner.Scan()
	path = scanner.Text()

	valid := isPath(path)
	if !valid {
		fmt.Println("[]")
		fmt.Println("[]")
		fmt.Println("[] NOT VALID PATH. Absolute Path ONLY")
		fmt.Println("** Example: /Users/CURRENT-USER/Desktop/Forgit/")
		fmt.Println("[]")
		fmt.Println("[] Suggest going to the Forgit Directory and running pwd command to get its path.")
		fmt.Println("   [ or ]")
		fmt.Println("[] Try Again --> fgt init")
		return
	}

	// build the config
	BuildConfig(string(path))
}
