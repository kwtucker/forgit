package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPath(p string) bool {
	// Get stat on the path
	isDirStat, err := os.Stat(p)
	if err != nil {
		fmt.Println("Path is not valid! Try again.")
		return false
	}

	switch {
	case strings.Contains(p, "Forgit") || strings.Contains(p, "forgit"):
		return false
	case string(p[0]) == "/" && string(p[len(p)-1:]) == "/" && isDirStat.IsDir() && string(p) != "/dev/null":
		return true
	}
	return false
}

// Init ...
func Init() {
	var (
		err        error
		scanner    *bufio.Scanner
		path, uuid string
	)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("<> Your Current Absolute Path is ->", currentDir)
	fmt.Println("<> Path cannot contain Forgit name.")
	fmt.Print("<> Enter Absolute path where you want the Forgit directory: ")
	scanner.Scan()
	path = scanner.Text()

	valid := isPath(path)
	if !valid {
		fmt.Println()
		fmt.Println("[] NOT VALID PATH. Absolute Path ONLY")
		fmt.Println("** Example: /Users/CURRENT-USER/Desktop/")
		fmt.Println("[]")
		fmt.Println("[] Suggest going to the directory and running pwd command to get its path.")
		fmt.Println("   [ or ]")
		fmt.Println("[] Try Again --> fgt init")
		return
	}

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("<> Enter UserId from Forgit Online Terminal Page: ")
	scanner.Scan()
	uuid = scanner.Text()

	// build the config
	BuildConfig(string(path), string(uuid))
}
