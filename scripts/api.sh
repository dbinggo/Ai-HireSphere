#!/bin/bash

# 定义log文件路径
log="./generate_api.log"
docs_dir="./docs"
application_dir="./application"
file_name="*"
# 遍历所有命令行参数
while [[ $# -gt 0 ]]; do
  case "$1" in
    -s)
      file_name="$2"
      shift 2  # 移动两个位置，跳过 -f 和它的值
      ;;
    -h | --help )
      echo "Usage: $0 [-f <file_name>]"
      echo "  -f <file_name>    指定要处理的api文件名，默认为当前目录下全部api文件"
      echo "  -h, --help        显示帮助信息"
      exit 0
      ;;
    *)  # 处理其他未知参数
      echo "未知参数: $1" >&2
      exit 1
      ;;
  esac
done

# 查找当前目录及其子目录下所有以 .api 结尾的文件
find "$application_dir" -type f -name "$file_name.api" | while read api_file; do
    # 获取文件所在目录
    dir=$(dirname "$api_file")
    # 提取文件名
    file_name=$(basename "$api_file")
    # 提取文件名（不包含扩展名）
    file_name_without_ext="${file_name%.*}"

    # 对每个 .api 文件调用 goctl 命令
    goctl api go -api "$api_file" -dir "$dir" --style=go_zero --home=./template >> $log 2>&1
    if [ $? -ne 0 ]; then
        echo -e "\033[31mFailed to generate code for $api_file \033[0m"
        cat $log
        exit 1
    fi

    goctl api plugin -plugin goctl-swagger="swagger -filename $file_name_without_ext.json" -api "$api_file" -dir "$docs_dir" >> $log 2>&1
    if [ $? -ne 0 ]; then
        echo -e "\033[31mFailed to generate code for $api_file \033[0m"
        cat $log
        exit 1
    fi
    echo -e "\033[42mGenerated code for $api_file \033[0m"
    rm -f $log
done


