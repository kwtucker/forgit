package lib

import (
	"os/exec"
	"sync"
)

// GitCommit command.
func GitCommit(message string, wg *sync.WaitGroup, notifyme int, notifymeError int) {
	defer wg.Done()
	// Git commit command
	args := []string{"commit", "-m", message}
	cmd := exec.Command("git", args...)
	err := cmd.Run()
	if err != nil {
		if notifymeError == 1 {
			m := &Message{
				Title: "Commit",
				Body:  message,
			}
			Notify(*m)
		}
	}

	if notifyme == 1 {
		m := &Message{
			Title: "Commit",
			Body:  message,
		}
		Notify(*m)
	}
}
