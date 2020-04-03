#!/bin/bash
for data in $(cat ../runtime/pid.txt)
do
   kill -1 $data
   echo "平滑重启服务成功!!!"
done
