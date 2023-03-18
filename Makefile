build:
	@cd cmd/tbgo && go build -o bin/TBGo

run: build
	@cd cmd/tbgo && ./bin/TBGo

test:
	@go test -v ./...


proto:
	protoc proto/*.proto --go_out=./proto/generated --go-grpc_out=./proto/generated


.PHONY: proto