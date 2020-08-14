#!/bin/bash

OutPath="output"

if [ -p /dev/stdin ]; then
        while IFS= read line; do
                IFS='/'
                read -a array <<< $line
                IFS='.'
                read -a array <<< ${array[-1]}
                name=${array[0]}
                git clone --mirror "$line" "$OutPath/raw/$name"
                git clone -l "$OutPath/raw/$name" "$OutPath/repo/$name"
        done
fi