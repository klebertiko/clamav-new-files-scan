#!/bin/bash
 
DIR=$HOME
 
# Get rid of old log file, if any
rm $HOME/virus-scan.log 2> /dev/null
mkdir $HOME/virus-quarantine
 
IFS=$(echo -en "\n\b")
 
# Optionally, you can use shopt to avoid creating two processes due to the pipe
shopt -s lastpipe

# Added '--recursive' so that a directory copied into $DIR also triggers clamscan/clamdscan, although downloads
# from the Web would just be files, not directories.
inotifywait --quiet --monitor --event create,modify,close_write,moved_to --recursive --format '%w%f' $DIR | while read FILE
do
     # Have to check file length is nonzero otherwise commands may be repeated
     if [ -s $FILE ]; then
          date > $HOME/virus-scan.log
          clamscan $FILE >> $HOME/virus-scan.log --move=$HOME/virus-quarantine
          zenity --warning --title "Virus scan of $FILE" --text "$(cat $HOME/virus-scan.log)"
     fi
done