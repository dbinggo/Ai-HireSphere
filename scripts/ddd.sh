#!/bin/bash
file_path="./application"
service_name=""
while [[ $# -gt 0 ]]; do
  case "$1" in
    -s)
      service_name="$2"
      shift 2  # 移动两个位置，跳过 -f 和它的值
      ;;
    -h | --help )
      echo "Usage: $0 [-f <file_name>]"
      echo "  -f <file_name>    指定要处理的proto文件名，默认为当前目录下全部proto文件"
      echo "  -h, --help        显示帮助信息"
      exit 0
      ;;
    *)  # 处理其他未知参数
      echo "未知参数: $1" >&2
      exit 1
      ;;
  esac
done

# 校验service_name
if [ -z "$service_name" ]; then
  echo "请输入服务名"
  exit 1
fi

base_dir="$file_path/$service_name"
# 如果存在文件夹就退出
if [ -d "$base_dir" ]; then
  echo -e "\033[41m $base_dir 文件服务已存在 \033[0m"
  exit 1
fi


domain_name="$file_path/$service_name/domain"
app_name="$file_path/$service_name/app"
infrastructure_name="$file_path/$service_name/infrastructure"
interfaces_name="$file_path/$service_name/interfaces"

# 创建必备文件夹
mkdir -p "$domain_name" "$app_name" "$infrastructure_name" "$interfaces_name"
mkdir -p "$domain_name/event" "$domain_name/irepository" "$domain_name/service" "$domain_name/model"
mkdir -p "$domain_name/model/aggregates" "$domain_name/model/entity" "$domain_name/model/vo"
mkdir -p "$infrastructure_name/repository"
# 使用 goctl demo 创建文件

# api
goctl api new "$interfaces_name/$service_name" --style=go_zero --home=./template
mv "$interfaces_name/$service_name" "$interfaces_name/api"

# 使用 goctl demo 创建rpc文件
# rpc 现在goctl 的 rpc new 含有bug 已经提交issue https://github.com/zeromicro/go-zero/issues/4588
pwd=$(pwd)
cd "$interfaces_name"
goctl rpc new "$service_name" --style=go_zero --home=./template
cd "$pwd"
mv "$interfaces_name/$service_name" "$interfaces_name/rpc"

find "$interfaces_name/api" -mindepth 1 -maxdepth 1 ! -name "$service_name.api" -exec rm -rf {} +
find "$interfaces_name/rpc" -mindepth 1 -maxdepth 1 ! -name "$service_name.proto" -exec rm -rf {} +

echo "# $service_name" > $base_dir/README.md


