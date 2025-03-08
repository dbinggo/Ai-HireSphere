#!/bin/bash

deploy_dir="./deploy"

# 遍历deploy下的所有文件
find "$deploy_dir" -name "Dockerfile_api_*" | while read docker_file; do
# 获取文件名
    file_name=$(basename "$docker_file")
    image_name=${file_name#Dockerfile_api_}
    docker buildx build -f "$docker_file" --platform linux/amd64 -t "registry.cn-guangzhou.aliyuncs.com/dbinggo-docker/$image_name-api:1.0.0"  --push .

done

# 遍历deploy下的所有文件
find "$deploy_dir" -name "Dockerfile_rpc_*" | while read docker_file; do
# 获取文件名
    file_name=$(basename "$docker_file")
    image_name=${file_name#Dockerfile_rpc_}
    docker buildx build -f "$docker_file" --platform linux/amd64 -t "registry.cn-guangzhou.aliyuncs.com/dbinggo-docker/$image_name-rpc:1.0.0"  --push .

done