package main

import (
        "fmt"
        "log"
        "os/exec"
        "strings"

        "github.com/dietsche/rfsnotify"
        fsnotify "gopkg.in/fsnotify.v1"
)

func execBashPipedCommand(command string) ([]byte, error) {
        cmd := exec.Command("bash", "-c", command)
        return cmd.CombinedOutput()
}

func installClamAV() {
        outClamd, _ := execBashPipedCommand("command -v clamd")
        outClamdscan, _ := execBashPipedCommand("command -v clamdscan")
        if outClamd == nil || outClamdscan == nil {
                outDistro, _ := execBashPipedCommand("lsb_release -si")
                
                switch distro := outDistro; distro {
                case "Ubuntu":
                        _, err := execBashPipedCommand("apt install -y clamav clamav-daemon")
                case "Fedora":
                        _, err := execBashPipedCommand("dnf install -y clamav clamav-update")
				}
				case "Arch":
						execBashPipedCommand("pacman -S clamav clamav-daemon")
        }
}

func refreshClamAVVirusDatabase() {
	execBashPipedCommand("freshclam")
}

func installScanService() {
	execBashPipedCommand("freshclam")
}

func main() {
        scanDir := flag.String("scan-dir", "", "Directory to be scanned")
        flag.Parse()
        if *scanDir == "" {
                flag.PrintDefaults()
                os.Exit(1)
        }

		installClamAV()
		refreshClamAVVirusDatabase()
        
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
                                        log.Println("Modified file:", event.Name)

                                        file := strings.Replace(event.Name, "(", "\\(", -1)
                                        file = strings.Replace(file, ")", "\\)", -1)
                                        file = strings.Replace(file, " ", "\\ ", -1)

                                        var clamavDaemonScan = fmt.Sprintf("clamdscan %v --remove=/opt/clamav/virus-quarantine", file)
                                        out, _ := execBashPipedCommand(clamavDaemonScan)
                                        log.Println(out)
                                }
                        case err, ok := <-watcher.Errors:
                                if !ok {
                                        return
                                }
                                log.Println("Error:", err)
                        }
                }
        }()

        err = watcher.AddRecursive(scanDir)
        if err != nil {
                log.Fatal(err)
        }
        <-done
}