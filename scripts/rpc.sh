#!/bin/bash

# 定义log文件路径
log="./generate_rpc.log"
call_dir="./common/call"
application_dir="./application"


file_name="*"
# 定义call目录
# 遍历所有命令行参数
while [[ $# -gt 0 ]]; do
  case "$1" in
    -s)
      file_name="$2"
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

# 查找当前目录及其子目录下所有以 .proto 结尾的文件
find "$application_dir" -type f -name "$file_name.proto" | while read rpc_file; do
    # 获取文件所在目录
    dir=$(dirname "$rpc_file")
    # 提取文件名
    file_name=$(basename "$rpc_file")
    # 提取文件名（不包含扩展名）
    file_name_without_ext="${file_name%.*}"
    # 对每个 .proto 文件调用 goctl 命令
    goctl rpc protoc $rpc_file --go_out="$call_dir" --go-grpc_out="$call_dir"  --zrpc_out="$dir" --style=go_zero --home=./template -c >> $log 2>&1
    if [ $? -ne 0 ]; then
        echo -e "\033[31mFailed to generate code for $rpc_file \033[0m"
        cat $log
        exit 1
    fi

    # 对每个生成的客户端移动到 common/call 目录下
    find "$dir" -type d -maxdepth 1  | while read client_file; do
        # 如果不为原来目录且不为etc目录和internal目录就移动到common/call目录下
        if [[ "$client_file" != "$dir" && "$client_file" != *"etc"* && "$client_file" != *"internal"* ]]; then
            mv -f "$client_file" "$call_dir"
            rm -f "$client_file"
        fi
    done

    echo -e "\033[42mGenerated code for $rpc_file \033[0m"
    rm -f $log
done

echo "Generate rpc success"