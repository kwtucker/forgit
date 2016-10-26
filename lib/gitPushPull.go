package lib

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

//GitPushPull is simply a git push command
func GitPushPull(p, branch, command string, wg *sync.WaitGroup, notifyme int, notifymeError int) {
	var (
		err    error
		args   []string
		cmd    *exec.Cmd
		remote string
		ww     sync.WaitGroup
	)

	// grab remote
	remote, err = getRemote(p)
	if err != nil {
		log.Println(err)
	}
	// count of WaitGroup
	ww.Add(1)
	switch command {
	case "push":
		go func() {
			defer ww.Done()
			args = []string{"push", remote, branch}
			cmd = exec.Command("git", args...)
			err = cmd.Run()
			if err != nil {
				if notifymeError == 1 {
					m := &Message{
						Title: "Forgit Error Push",
						Body:  " ",
					}
					Notify(*m)
				}
			}
			if notifyme == 1 {
				m := &Message{
					Title: "Push Event",
					Body:  " ",
				}
				Notify(*m)
			}
			log.Println(": git push "+remote, branch)
		}()
	case "pull":
		go func() {
			defer ww.Done()
			args = []string{"pull", remote, branch}
			cmd = exec.Command("git", args...)
			err = cmd.Run()
			if err != nil {
				if notifymeError == 1 {
					m := &Message{
						Title: "Forgit Error Pull",
						Body:  " ",
					}
					Notify(*m)
				}
			}
		}()
	}
	ww.Wait()
	wg.Done()
}

// Grabs the remote b
func getRemote(path string) (string, error) {
	var (
		err    error
		remote string
	)

	err = os.Chdir(path)
	if err != nil {
		log.Println(err)
	}

	// get the remote command output
	remoteCmd := exec.Command("git", "remote")
	remoteStdout, err := remoteCmd.Output()

	// Split output on the new line and turn it into a slice.
	rSplit := strings.Split(string(remoteStdout), "\n")
	lineArr := rSplit[0 : len(rSplit)-1]

	// get only the origin
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
