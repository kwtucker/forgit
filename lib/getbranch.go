package lib

import (
	"os"
	"os/exec"
	"strings"
)

// GetCurrentBranch returns a branch name
func GetCurrentBranch(p string) (string, error) {
	var (
		err    error
		branch string
	)

	// Change directroy to path and run git branch.
	err = os.Chdir(p)
	gpArgs := []string{"branch"}
	gitBranch := exec.Command("git", gpArgs...)
	branchByte, err := gitBranch.Output()

	// Format output to a slice and grab the one with a star next to it.
	brArr := strings.Split(string(branchByte), "\n")
	for i := range brArr {
		if strings.Contains(brArr[i], "*") {
			sub := strings.Split(brArr[i], "*")
			branch = strings.TrimSpace(sub[1])

		}
	}
	return branch, err
}
