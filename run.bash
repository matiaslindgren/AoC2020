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
	local input=txt/input/${day}.txt
	local exe=./bin/$day

	printf "\n%s\n" $exe
	assert_is_file $input
	assert_is_file $exe

	printf "input:  %d lines from %s\n" $(wc -l $input)
	local output="$(cat $input | $exe)"
	printf "output: %s\n" "$output"

	local expect=txt/expect/${day}.txt
	if [ -f "$expect" ]; then
		local e="$(cat $expect)"
		printf "expect: %s\n" "$e"
		if [ "$output" != "$e" ]; then
			printf "%s != %s\n" "$output" "$e"
			exit 1
		fi
	fi
}


if [ $# -eq 0 ]; then
	for day in $(find -s bin -regex 'bin/[0-2][0-9]$'); do
		run_day $(basename $day)
	done
else
	run_day $1
fi
