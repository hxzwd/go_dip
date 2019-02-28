#!/bin/bash

SRC=`ls *.go`
CMD="go run"

for src in $SRC
do
#	echo $src
	CMD=$CMD" $src"
done

eval "$CMD"

#go run "$SRC"

