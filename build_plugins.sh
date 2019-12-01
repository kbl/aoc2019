#!/bin/bash

export GOPATH=`pwd`

for day_path in `find . -type d | grep day`
do
    day=`echo $day_path | cut -f 4 -d '/'`
    echo Building plugin for $day
    go build -buildmode=plugin -o plugins/$day.so $day_path
done
