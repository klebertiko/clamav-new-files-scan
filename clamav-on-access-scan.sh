#!/bin/bash
IFS=$(echo -en "\n\b")
 
# Get rid of old log file, if any
rm -f /var/log/clamav/virus-scan.log 2> /dev/null
mkdir -p /opt/clamav/virus-quarantine

# Optionally, you can use shopt to avoid creating two processes due to the pipe
shopt -s lastpipe

#inotifywait --exclude '(/var/log/clamav/virus-scan.log|dconf/user|/proc/|/sys/)' --quiet --monitor --recursive --event access,modify,close_write,open,moved_to,create --format '%w%f' /home | while read FILE
inotifywait --exclude '(/var/log/clamav/virus-scan.log|dconf/user|/proc/|/sys/)' --quiet --monitor --recursive --event access,modify,close_write,open,moved_to,create --format '%w%f' /home /usr /opt /mnt /media /tmp | while read FILE
#lsof -e /run/user/1000/gvfs | grep REG | grep -v DEL | awk '{if ($NF=="(deleted)") {x=3;y=1} else {x=2;y=0}; {print $(NF-x) "  " $(NF-y) } }' | sort -n -u | numfmt --field=1 --to=iec | awk '{print $2}' | while read FILE
do
     # Have to check file length is nonzero otherwise commands may be repeated
     if [ -s "$FILE" ]; then
          date >> /var/log/clamav/virus-scan.log
          # clamscan "$FILE" >> /var/log/clamav/virus-scan.log --move=/opt/clamav/virus-quarantine
          clamdscan "$FILE" >> /var/log/clamav/virus-scan.log --move=/opt/clamav/virus-quarantine
     fi
done