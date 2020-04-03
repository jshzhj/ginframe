#!/bin/bash
cd ..
go build -o main ./main.go
rm -rf ./release/*
mkdir ./release/main
mkdir ./release/main/scripts
mkdir ./release/main/runtime
mv ./main ./release/main

cp -r ./asset ./release/main
cp -r ./conf ./release/main
cp -r ./scripts/restart_server.sh ./release/main/scripts
cp -r ./scripts/stop_server.sh ./release/main/scripts


