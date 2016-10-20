package lib

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isPath(p string) (string, bool) {
	if len(p) == 0 {
		return p, true
	}
	// Get stat on the path
	isDirStat, err := os.Stat(p)
	if err != nil {
		fmt.Println("Path is not valid! Try again.")
		return p, false
	}
	switch {
	case strings.Contains(p, "Forgit") || strings.Contains(p, "forgit"):
		return p, false
	case string(p[0]) == "/" && string(p[len(p)-1:]) == "/" && isDirStat.IsDir() && string(p) != "/dev/null":
		return p, true
	}
	return p, false
}

// Init ...
func Init() {
	var (
		err           error
		scanner       *bufio.Scanner
		path, uuid, p string
	)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("<> Your Current Absolute Path is ->", currentDir)
	fmt.Println("<> Path cannot contain Forgit name.")
	fmt.Print("<> Enter Absolute path where you want the Forgit directory [ Enter For Here ]: ")
	scanner.Scan()
	path = scanner.Text()

	p, valid := isPath(path)
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

	if p == " " || p == "" {
		path = currentDir + "/"
	}

	scanner = bufio.NewScanner(os.Stdin)
	fmt.Print("<> Enter UUID from Forgit Online Terminal Page: ")
	scanner.Scan()
	uuid = scanner.Text()

	// build the config
	BuildConfig(string(path), string(uuid))
}
