package lib

import (
	"github.com/deckarep/gosx-notifier"
	"log"
	"os/exec"
	"runtime"
)

//Notify ...
func Notify(msg Message) {
	notifyTitle := msg.Title
	notifyMessage := msg.Body
	switch runtime.GOOS {
	case "darwin":
		// osx notification
		note := gosxnotifier.NewNotification(notifyMessage)
		note.Title = notifyTitle
		note.Group = "Forgit"
		err := note.Push()
		if err != nil {
			log.Println(msg.Title, "No updates, notification didn't happen")
		}
	case "linux":
		exec.Command("notify-send", "-i", "./logo_icon.png", notifyTitle, notifyMessage, "-u", "critical").Run()
	}
}
