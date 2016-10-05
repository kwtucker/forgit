package lib

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
)

func GitStatus(committime int, path string, repos []SettingRepo) {

	for {
		var (
			gitStatus *exec.Cmd
			fileName  string
		)

		// Ticker to make for loop wait tell push time
		Ticker(committime)

		// making a type
		reposMap := make(map[string][]string)

		for r := range repos {
			// creating a new wait group and adding 1 goroutine to it.
			var ww sync.WaitGroup
			ww.Add(1)

			var err error
			err = os.Chdir(path + repos[r].Name)
			// Git status command
			gsArgs := []string{"status", "-s"}
			gitStatus = exec.Command("git", gsArgs...)

			// Grab the stdout of git status
			gsOut, err := gitStatus.StdoutPipe()
			if err != nil {
				log.Println(err)
			}

			// Create a new scanner to parse the gsOut
			scanner := bufio.NewScanner(gsOut)
			// go routine with WaitGroup hold the function before end.
			// Scan each line, get the text, split it by the white space.
			// append the text the filename
			go func(n map[string][]string) {
				defer ww.Done()
				for scanner.Scan() {
					line := scanner.Text()
					lineArr := strings.Split(line, " ")
					if len(lineArr[1]) > 2 {
						fileName = lineArr[1]
					} else {
						fileName = lineArr[2]
					}
					var modfiles []string
					modfiles = append(modfiles, fileName)
					n[repos[r].Name] = modfiles
				}
			}(reposMap) // pass in map to go routine

			gitStatus.Start()
			ww.Wait()
		}
		fmt.Println("Modified Files", reposMap)
	}

}
