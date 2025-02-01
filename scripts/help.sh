#!/bin/bash

echo -e "\033[42m Makefile 自动生成代码说明 \033[0m"
echo -e "\033[42m make api \033[0m                将/application下所有的api文件生成api文件对应的go-zero代码和swagger代码docker代码"
echo -e "\033[42m make rpc \033[0m                将/application下所有的proto文件生成proto文件对应go-zero的代码client代码和docker代码"
echo -e "\033[42m make gen \033[0m                执行make api和make rpc"
echo -e "\033[42m make ddd SERVICE=xxxx \033[0m   自动生成ddd型的代码包括api和rpc"