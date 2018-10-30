package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/fsnotify/fsnotify"
)

func execBashPipedCommand(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)

	out, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	if out != nil {
		fmt.Printf(string(out))
	}
	return string(out), err
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
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)

					// var err = os.Remove($HOME/virus-scan.log)
					// if err != nil {
					// 	log.Print(err.Error())
					// }

					// out, err := os.Create("$HOME/virus-quarantine")
					// if err != nil {
					// 	log.Println(err)
					// } else {
					// 	defer out.Close()
					// }

					// var clamavScan = fmt.Sprintf("clamdscan --move=$HOME/virus-quarantine %s >> $HOME/virus-scan.log", event.Name)
					// var zenityWarning = fmt.Sprintf("zenity --warning --title \"Virus scan of %s\" --text \"$(cat $HOME/virus-scan.log)\"", event.Name)
					// var clamavScanZenityWarning = fmt.Sprintf("%s | %s", clamavScan, zenityWarning)
					// execBashPipedCommand(clamavScanZenityWarning)
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
