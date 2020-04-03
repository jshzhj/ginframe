#!/bin/bash
for data in $(cat ../runtime/pid.txt)
do
   kill -15 $data
   echo "平滑关闭服务成功!!!"
done