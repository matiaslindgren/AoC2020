#!/usr/bin/env bash
set -ue


function assert_is_file {
	if [ ! -f "$1" ]; then
		echo "run.sh: '$1' does not exist or is not a file" > /dev/stderr
		exit 1
	fi
}


function run_day {
	local day=$1
	local input=input/${day}.txt
	local exe=./bin/$day

	printf "\n%s\n" $exe
	assert_is_file $input
	assert_is_file $exe

	printf "input:  %d lines from %s\n" $(wc -l $input)
	printf "output: %s\n" "$(cat $input | $exe)"
}


if [ $# -eq 0 ]; then
	for day in $(find -s bin -regex 'bin/[0-2][0-9]$'); do
		run_day $(basename $day)
	done
else
	run_day $1
fi
