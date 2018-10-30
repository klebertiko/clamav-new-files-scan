package main

import (
	"log"
	"os"
	"os/user"

	"github.com/fsnotify/fsnotify"
)

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

					usr, err := user.Current()
					if err != nil {
						log.Fatal(err)
					}
					log.Println(usr)

					// var err = os.Remove($HOME/virus-scan.log)
					// if err != nil {
					// 	log.Print(err.Error())
					// }

					out, err := os.Create("")
					if err != nil {
						log.Println(err)
					} else {
						defer out.Close()
					}

					// var clamavScan = fmt.Sprintf("clamdscan --move=$HOME/virus-quarantine %s >> $HOME/virus-scan.log", event.Name)
					// var zenity = fmt.Sprintf("zenity --warning --title \"Virus scan of %s\" --text \"$(cat $HOME/virus-scan.log)\"", event.Name)
					// var clamavScanZenity = fmt.Sprintf("%s | %s", clamavScan, zenity)
					// cmd := exec.Command("bash", "-c", clamavScanZenity)

					// out, err := cmd.CombinedOutput()
					// if err != nil {
					// 	log.Fatalf("==> ExecBashPipedCommand failed with %s\n", err)
					// }
					// if out != nil {
					// 	fmt.Printf("==> Out \n%s\n", string(out))
					// }
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
