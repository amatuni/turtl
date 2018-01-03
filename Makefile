
all:
	@go get ./cli/...
	@go build -i -o ./turtl -ldflags "-X main.gitHash=`git rev-parse HEAD`" ./cli/*.go

install:
	@go get ./cli/...
	@go build -o $(GOPATH)/bin/turtl -ldflags "-X main.gitHash=`git rev-parse HEAD`" ./cli

libturtl:
	@go install -ldflags "-X main.gitHash=`git rev-parse HEAD`"

proto:
	@protoc -I . turtl.proto --gofast_out=plugins=grpc:cli

clean:
	@go clean -i github.com/andreiamatuni/
	@rm -f ./turtl
	@rm -f $(GOPATH)/bin/turtl
	