#!/usr/bin/env sh
set -ue

for day in $(ls bin); do
	echo;
	input=input/${day}.txt;
	exe=./bin/$day;
	echo $exe;
	[ ! -f "$input" -o ! -f "$exe" ] && exit 1;
	echo i: $(wc -l $input);
	echo o: $(cat $input | $exe);
done
