#!/bin/bash
go build -o main ../main.go
mv ./main ../release/
cp -r ../asset/* ../release