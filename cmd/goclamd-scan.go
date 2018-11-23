package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/dietsche/rfsnotify"
	fsnotify "gopkg.in/fsnotify.v1"
)

func execBashPipedCommand(command string) ([]byte, error) {
	return exec.Command("bash", "-c", command).CombinedOutput()
}

func clamdScan(eventName string) {
	log.Println("Modified file:", eventName)

	file := strings.Replace(eventName, "(", "\\(", -1)
	file = strings.Replace(file, ")", "\\)", -1)
	file = strings.Replace(file, " ", "\\ ", -1)

	var clamavDaemonScan = fmt.Sprintf("clamdscan %v --remove", file)
	out, _ := execBashPipedCommand(clamavDaemonScan)
	log.Println(out)
}

func main() {
	scanDir := flag.String("scan-dir", "/home", "Directory to be scanned")
	flag.Parse()
	if *scanDir == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	watcher, err := rfsnotify.NewWatcher()
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
				log.Println("Event:", event)
				if event.Op&fsnotify.Create == fsnotify.Create || event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename || event.Op&fsnotify.Chmod == fsnotify.Chmod {
					clamdScan(event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Error:", err)
			}
		}
	}()

	err = watcher.AddRecursive(*scanDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
