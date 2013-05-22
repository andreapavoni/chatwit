GOFMT=gofmt -s -tabs=false -tabwidth=2
GOFILES=$(wildcard *.go **/*.go)

default:
	go build

format:
	${GOFMT} -w ${GOFILES}

clean:
	rm chatwit

