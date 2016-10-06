package lib

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//GitPushPull is simply a git push command
func GitPushPull(p, branch, command string, wg *sync.WaitGroup) {
	var (
		err error
		// args   []string
		// cmd    *exec.Cmd
		remote string
		ww     sync.WaitGroup
	)

	remote, err = getRemote(p)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(remote)
	ww.Add(1)
	switch command {
	case "push":
		go func() {
			defer ww.Done()
			// args = []string{"push", remote, branch}
			// cmd = exec.Command("git", args...)
			fmt.Println("git push "+remote, branch)
		}()
	case "pull":
		fmt.Println("pull it")
	}
	ww.Wait()
	wg.Done()

}

func getRemote(path string) (string, error) {
	var (
		err    error
		remote string
	)

	err = os.Chdir(path)
	if err != nil {
		log.Println(err)
	}

	remoteCmd := exec.Command("git", "remote")
	remoteStdout, err := remoteCmd.Output()

	rSplit := strings.Split(string(remoteStdout), "\n")
	lineArr := rSplit[0 : len(rSplit)-1]

	for _, v := range lineArr {
		if v == "origin" {
			remote = v
		}
	}
	if remote == "" && len(lineArr) < 2 {
		remote = lineArr[0]
	}
	return remote, err
}
