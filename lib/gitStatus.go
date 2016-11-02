package lib

import (
	"log"
	"os"
	"os/exec"
	"strings"
)

// GitStatus gets the git status of the repos directory
func GitStatus(p string) ([]string, error) {
	var (
		err      error
		fileName string
		modfiles []string
	)

	// Change directory to the path passed in
	err = os.Chdir(p)
	if err != nil {
		log.Println(err)
	}

	// Git status command
	gsArgs := []string{"status", "-s"}
	gitStatus := exec.Command("git", gsArgs...)

	// Get stdout and trim the empty last index
	fileStatus, err := gitStatus.Output()
	fsSplit := strings.Split(string(fileStatus), "\n")
	lineArr := fsSplit[0 : len(fsSplit)-1]

	for i := range lineArr {
		splitIndex := strings.Split(lineArr[i], " ")
		if len(splitIndex[1]) >= 2 {
			fileName = splitIndex[1]
		} else {
			fileName = splitIndex[2]
		}
		modfiles = append(modfiles, fileName)
	}
	return modfiles, err
}
