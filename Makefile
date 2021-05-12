go_bin:=$(shell pwd)/bin

# if you want to implement api updates then use bash `make`
# and extend/update/fix sdk code after

all: clean tools gen dep lint

tools:
	GOBIN=${go_bin} go install -mod=mod github.com/go-swagger/go-swagger/cmd/swagger
	GOBIN=${go_bin} go install -mod=mod github.com/golangci/golangci-lint/cmd/golangci-lint@v1.39.0

clean:
	rm -rf ./bin
	rm -rf ./swagger
	mkdir ./bin
	mkdir ./swagger

dep:
	go mod tidy
	go mod vendor

gen:
	./bin/swagger generate client -f https://storage.gcorelabs.com/api/docs/swagger.json -t swagger

updatedep:
	go list -m -u all

lint:
	./bin/golangci-lint run ./sdk*.go

test:
	go test --count=1 -v -race ./...

help:
	echo "if you want to implement api updates then use bash `make`. and extend/update/fix sdk code after."