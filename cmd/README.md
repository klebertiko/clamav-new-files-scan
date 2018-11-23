# goclamd-scan

## GoLang using clamav daemon to scan specified folder

* Scan *nix specified folder
* Remove malware

### Manual install ClamAV and ClamAV Daemon

* `$ sudo apt install clamav clamav-daemon`

### Update virus database

* `$ sudo freshclam`

### Downloading

* `git clone https://innersource.accenture.com/scm/concrete-tecs/alcs-iso.git`
* `$ cd alcs-iso/clamav/goclamd-scan/cmd`

### goclamd-scan-installer

* This binary will install clamav, clamav-daemon and clamtk, configure clamd to run as root, install goclamd-scan for specified folder with --scan-dir and configure goclamd-scan as service. The goclamd-scan binary must be in the same directory as goclamd-scan-installer
* `$ sudo ./goclamd-scan-installer --scan-dir /home`

### Building

* `$ go build -a goclamd-scan.go`

### Installing

* `$ go install -a goclamd-scan.go`

### Change clamd run user and group to root

* `$ sudo vi /etc/clamav/clamd.conf`

#### Change values bellow:

```
LocalSocketGroup root
User root
```

#### Reload services and clamav daemon:

* `$ sudo systemctl daemon-reload`
* `$ sudo systemctl restart clamav-daemon.service`

#### Check if clamav daemon is runnning:

* `$ sudo systemctl list-units | clamav-daemon.service`
* `$ sudo systemctl status clamav-daemon.service`

### Increase max watches for inotify

* `$ echo fs.inotify.max_user_watches=524288 | sudo tee -a /etc/sysctl.conf && sudo sysctl -p`

### Create a systemd service for goclamd-scan

* `$ sudo cp goclamd-scan-home.service /etc/systemd/system/`
* `$ sudo cp goclamd-scan-home /opt/clamav/`
* `$ sudo systemctl daemon-reload`
* `$ sudo systemctl enable goclamd-scan.service`
* `$ sudo systemctl start goclamd-scan.service`


### Check if goclamd-scan is running

* `$ sudo systemctl list-units | grep goclamd-scan`
* `$ sudo systemctl status goclamd-scan.service`

### Installer and scan flag

* ` --scan-dir = passed for installer to configure scan service for specified dir`
* `$ sudo ./goclamd-scan-installer --scan-dir /home`
* `$ sudo ./goclamd-scan --scan-dir /home`