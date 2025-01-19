# 以下命令执行前，请先保证自己的common目录下有xcode和interceptors代码，直接复制该项目下的代码文件即可
# 使用template模版生成api代码(请在api文件所在的目录下执行，并把*改成对应api文件名字)
goctl api go -api *.api -dir ./  --style=go_zero --home=../../../../template
# 使用template模版生成rpc代码(请在proto文件所在的目录下执行，并把*改成对应proto文件名字)
goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=goZero --home=../../../../template

# 使用template模版生成docker代码 需要跑到指定文件目录下执行 api目录或者rpc目录
goctl docker --go user.go --exe user-rpc --home=../../../../template --version:1.23


# 生成swigger 用于导入apifox
# 插件下载
GOPROXY=https://goproxy.cn/,direct go install github.com/zeromicro/goctl-swagger@latest

# 插件生成
goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir ../../../../docs
