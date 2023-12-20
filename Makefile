.PHONY: proto
proto:
	@oapi-codegen  --config api/config.yaml api/openapi.json


py_package:
	@pip3 install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple

stop:
	@kill -9 $(ps aux | grep medical | grep -v grep | awk '{print $2}')

build:
	@go mod tidy && go build -o medical .

run:
	@nohup ./medical &

test:
	@go mod tidy && go run main.go