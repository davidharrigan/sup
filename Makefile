BINARY=sup
VERSION=0.0.1

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X github.com/davidharrigan/sup/cmd.Version=${VERSION}"

binary:
	go build ${LDFLAGS} -o bin/${BINARY} main.go

.PHONY: clean
clean:
	if [ -f bin/${BINARY} ] ; then rm bin/${BINARY} ; fi
