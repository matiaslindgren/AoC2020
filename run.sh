#!/usr/bin/env sh
set -ue


function assert_is_file {
	if [ ! -f "$1" ]; then
		echo "$0: error: '$1' does not exist or is not a file" > /dev/stderr
		exit 1
	fi
}


function run_day {
	local day=$1
	local input=./txt/input/${day}.txt
	local exe=./bin/$day

	printf "\n%s\n" $exe
	assert_is_file $input
	assert_is_file $exe

	printf "input:  %d lines from %s\n" $(wc -l $input)
	local output="$(cat $input | $exe)"
	printf "output: %s\n" "$output"

	local expect=./txt/expect/${day}.txt
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
	for day in $(find bin -regex 'bin/[0-2][0-9]$' | sort); do
		time run_day $(basename $day)
	done
else
	time run_day $1
fi
