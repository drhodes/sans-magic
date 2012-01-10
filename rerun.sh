#!/bin/bash

WATCH_PATH=`pwd`

while inotifywait -e modify -r ./ 2> /dev/null; do	
	clear
	echo "------------------------------------------------------------------"
	make -s clean
	make > /tmp/makelog

	if [ "$?" -eq "0" ]
	then 
		echo "build succeeds"
		make install
	else 			
		cat /tmp/makelog
		echo "build failure"
	fi
done
