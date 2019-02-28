#!/bin/bash

SRC=$(ls *.go)
OUT="bin_main"

go build -o $OUT $SRC

echo $1

if [[ $1 == "run" ]]
then
	./$OUT
fi
