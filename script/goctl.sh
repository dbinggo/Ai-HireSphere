# 以下命令执行前，请先保证自己的common目录下有xcode和interceptors代码，直接复制该项目下的代码文件即可
# 使用template模版生成api代码(请在api文件所在的目录下执行，并把*改成对应api文件名字)
goctl api go -api *.api -dir ./  --style=goZero --home=../../../template
# 使用template模版生成rpc代码(请在proto文件所在的目录下执行，并把*改成对应proto文件名字)
goctl rpc protoc *.proto --go_out=./ --go-grpc_out=./  --zrpc_out=./ --style=goZero --home=../../../template