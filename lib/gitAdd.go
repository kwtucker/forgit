package lib

import (
	"os/exec"
	"sync"
)

// GitAdd command.
func GitAdd(file string, wg *sync.WaitGroup) {
	defer wg.Done()
	args := []string{"add", file}
	cmd := exec.Command("git", args...)
	cmd.Run()
}
