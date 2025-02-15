#!/bin/bash

# 删除 /common/call 目录下所有文件
rm -rf ./common/call/*
# 删除 /docs/ 目录下的所有swigger文件
rm -rf ./docs/*.json
# 删除所有日志信息
rm -rf ./logs/*
# 删除 /deploy/ 目录下的所有docker文件
rm -rf ./deploy/Dockerfile_api_*
rm -rf ./deploy/Dockerfile_rpc_*
