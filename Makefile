GOFMT=gofmt -s -tabs=false -tabwidth=2
GOFILES=$(wildcard *.go **/*.go)

default:
	go build

format:
	${GOFMT} -w ${GOFILES}

clean:
	if [ -f chatwit ] ; then rm chatwit ; fi
	if [ -d public/assets ] ; then  rm -rf public/assets ; fi
	if [ -f .*.css ] ; then  rm .*.css ; fi
	if [ -f .*.js ] ; then  rm .*.js ; fi

setup:
	go get github.com/shaoshing/train
	go get github.com/alloy-d/goauth
	go get github.com/gorilla/sessions
	go get github.com/gorilla/mux
	go get code.google.com/p/go.net/websocket

assets-bundle:
	${GOPATH}/bin/train bundle

