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

assets-bundle:
	${GOPATH}/bin/train bundle

