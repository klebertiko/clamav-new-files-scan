package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func execBashPipedCommand(command string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", command)
	return cmd.CombinedOutput()
}

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("modified file:", event.Name)

					user, _ := execBashPipedCommand("who | cut -d' ' -f1")
					homeDir, _ := execBashPipedCommand(fmt.Sprintf("getent passwd %s | cut -d: -f6", user[:len(user)-1]))

					os.Remove(fmt.Sprintf("%s/virus-scan.log", homeDir[:len(homeDir)-1]))
					os.Mkdir(fmt.Sprintf("%s/virus-quarantine", homeDir[:len(homeDir)-1]), os.ModePerm)

					file := strings.Replace(event.Name, "(", "\\(", -1)
					file = strings.Replace(file, ")", "\\)", -1)
					file = strings.Replace(file, " ", "\\ ", -1)

					var clamavScan = fmt.Sprintf("clamscan %v >> %s/virus-scan.log --move=%s/virus-quarantine", file, homeDir[:len(homeDir)-1], homeDir[:len(homeDir)-1])
					var zenityWarning = fmt.Sprintf("zenity --warning --title \"Virus scan of %v\" --text \"$(cat %s/virus-scan.log)\"", event.Name, homeDir[:len(homeDir)-1])
					execBashPipedCommand(clamavScan)
					execBashPipedCommand(zenityWarning)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add("/home/kleber/Downloads/")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
