#!/bin/bash
#暂时只支持mac->windows,mac->linux,mac->mac
cd ..
target=$1 #目标操作系统
if [[ ! -n "$target" ]]; then
  echo "请填写目标平台eg: windows || linux || mac"
  exit
fi

os=$(uname -a) #当前操作系统
b="Darwin"
if [[ $os =~ $b ]]; then
  echo "当前为mac平台"
  #编译为mac平台应用程序
  if [[ $target =~ "mac" ]]; then
    echo "正在编译mac平台可执行程序..."
    go build -o main ./main.go
  fi

  #编译为windows平台应用程序
  if [[ $target =~ "windows" ]]; then
    echo "正在编译windows平台可执行程序..."
    CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o main ./main.go
  fi

  #编译为linux平台应用程序
  if [[ $target =~ "linux" ]]; then
    echo "正在编译linux平台可执行程序..."
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go
  fi

fi

rm -rf ./release/*
mkdir ./release/main
mkdir ./release/main/scripts
mkdir ./release/main/runtime
mv ./main ./release/main

cp -r ./asset ./release/main
cp -r ./conf ./release/main
cp -r ./scripts/restart_server.sh ./release/main/scripts
cp -r ./scripts/stop_server.sh ./release/main/scripts
