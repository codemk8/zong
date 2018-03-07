build:
	go build -o bin/zong-server github.com/codemk8/zong/cmd/zong-server

swagger: zong.yaml
	swagger generate server -f zong.yaml

