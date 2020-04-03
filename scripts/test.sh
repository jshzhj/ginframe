#!/bin/bash

os=$(uname -a) #获取当前操作系统
target=$1      #目标平台
b="Darwin"
c="centos"
d="ubuntu"

if [[ $os =~ $b ]]; then
  echo "当前平台为mac"

elif [[ $os =~ $c ]]; then
  echo "当前平台为linux"

elif [[ $os =~ $d ]]; then

  echo "当前平台为linux"
else
  echo $os
fi

case $target in
  mac)
    echo "正在编译为mac平台下可执行程序..."
    CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
    ;;
  linux)
    echo "正在编译为linux平台下可执行程序..."

    ;;
  windows)
    echo "正在编译为windows平台下可执行程序..."

    ;;
  esac
