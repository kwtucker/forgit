package lib

import (
	"os/exec"
	"sync"
)

// GitAdd command.
func GitAdd(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Git add command
	gsArgs := []string{"add", file}
	gitStatus := exec.Command("git", gsArgs...)
	gitStatus.Run()

}
