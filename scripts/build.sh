#!/bin/bash

deploy_dir="./deploy"
HARBOR_URL="8.134.142.155:8888"
HARBOR_REPO="hiresphere"
GIT_COMMIT=$1

# 遍历deploy下的所有文件
find "$deploy_dir" -name "Dockerfile_api_*" | while read docker_file; do
# 获取文件名
    file_name=$(basename "$docker_file")
    image_name=${file_name#Dockerfile_api_}
    docker buildx build -f "$docker_file" --platform linux/amd64 -t "$HARBOR_URL/$HARBOR_REPO/$image_name-api:$GIT_COMMIT" --push .

done

# 遍历deploy下的所有文件
find "$deploy_dir" -name "Dockerfile_rpc_*" | while read docker_file; do
# 获取文件名
    file_name=$(basename "$docker_file")
    image_name=${file_name#Dockerfile_rpc_}
    docker buildx build -f "$docker_file" --platform linux/amd64 -t "$HARBOR_URL/$HARBOR_REPO/$image_name-rpc:$GIT_COMMIT"  --push .
done