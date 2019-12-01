#!/bin/bash

last_day=`find src/aoc/ -type d | grep day | sort | tail -n 1 | cut -f 3 -d '/'`
if [ ${last_day:3:1} = "0" ]
then
    last_day=${last_day:4}
else
    last_day=${last_day:3}
fi

next_day=`printf "%02d" $(($last_day + 1))`
echo Creating day $next_day

mkdir src/aoc/day${next_day}
cp src/aoc/day00/day00.go src/aoc/day${next_day}/day${next_day}.go
cp src/aoc/day00/day00_test.go src/aoc/day${next_day}/day${next_day}_test.go
