gen:
	bash ./scripts/clean.sh
	bash ./scripts/api.sh
	bash ./scripts/rpc.sh

api:
	rm -rf ./docs/*.json
	bash ./scripts/api.sh

rpc:
	rm -rf ./common/call/*
	bash ./scripts/rpc.sh

ddd:
	# 使用方法 make ddd SERVICE=tstt 会生成对应目录的ddd文件
	echo -e "开始生成$(SERVICE)服务的ddd文件"
	bash ./scripts/ddd.sh -s "$(SERVICE)"
	bash ./scripts/api.sh -s "$(SERVICE)"
	bash ./scripts/rpc.sh -s "$(SERVICE)"

clean:
	bash  ./scripts/clean.sh



