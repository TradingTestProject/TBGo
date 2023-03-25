NAME = tbgo
VERSION = 0.0.1

.PHONY: install-dev-tools
 install-dev-tools:
	go install github.com/abice/go-enum@v0.5.5
	go install github.com/golang/mock/mockgen@v1.6.0
	go install golang.org/x/tools/cmd/goimports@v0.7.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.1

.PHONY: pre-commit
pre-commit:
	goimports -local gitlab.s2b.tech -w ./internal/..
	gofmt -s -w ./internal/..
	golangci-lint run ./internal/...

build:
	@cd cmd/tbgo && go build -v -ldflags "-s -w"  -o bin/TBGo

run: build
	@cd cmd/tbgo && ./bin/TBGo

test:
	@go test -v ./...

.PHONY: gen-proto
gen-proto:
	protoc proto/*.proto --go_out=./proto/generated --go-grpc_out=./proto/generated

docker_build:
	docker build -t $(NAME):$(VERSION) .

