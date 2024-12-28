FROM golang:1.23 AS builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,http://goproxy.xiaoe-tools.com,direct GONOSUMDB=talkcheap.xiaoeknow.com

COPY go.mod go.mod
COPY go.sum go.sum

RUN  go mod download

COPY . .

RUN go clean

ARG COMMITID=$COMMITID
# 构建二进制文件命令,替换为自身程序的构建命令
RUN  go build -ldflags "-X main.name=transaction_server_go -X main.version=${COMMITID}" -a -o transaction_server_go ./cmd


FROM centos:7

WORKDIR /app

COPY --from=builder /app/ ./

#程序启动命令，启动命令需要录入到服务管理中
#ENTRYPOINT ["./transaction_server_go", "run", "-c", ".env.production", "-p", "2345"]
