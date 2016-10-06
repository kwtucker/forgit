package lib

import (
	// "fmt"
	// "log"
	// "os"
	"os/exec"
	// "strings"
	// "sync"
)

func GitAdd(file string) {

	// Git status command
	gsArgs := []string{"add", file}
	gitStatus := exec.Command("git", gsArgs...)
	gitStatus.Run()
}
