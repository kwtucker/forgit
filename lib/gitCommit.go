package lib

import (
	"os/exec"
	"sync"
)

// GitCommit command.
func GitCommit(message string, wg *sync.WaitGroup) {
	defer wg.Done()
	// Git commit command
	args := []string{"commit", "-m", message}
	cmd := exec.Command("git", args...)
	cmd.Run()

}
