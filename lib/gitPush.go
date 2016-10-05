package lib

import (
	"fmt"
	// "time"
)

//GitPush is simply a git push command
func GitPush(ptime int, repos []SettingRepo) {

	for {
		// a delay in the for loop
		Ticker(ptime)
		// Where the push code is going
		fmt.Println("push")
	}

}
