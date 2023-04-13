#!/usr/bin/bash

# Declare all directories, that are needed in the docker image:
declare -a dataDirectories=("dataFolder1"
                            "dataFolder2"
                            )
archive=data.tar.gz

tarCommand="tar -czf ${archive}"

for i in "${dataDirectories[@]}"
do
    tarCommand="${tarCommand} ${i}"
done
    tarCommand="${tarCommand}"

$tarCommand