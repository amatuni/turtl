
all:
	@go build -ldflags "-X main.GitHash=`git rev-parse HEAD`"

install:
	@go install -ldflags "-X main.GitHash=`git rev-parse HEAD`"