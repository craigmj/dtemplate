SHELL := /bin/bash

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKG = dtemplate
DEST = $(PKG)

bin/$(DEST): $(SRC)
	GOPATH=`pwd` go build -o bin/$(DEST) src/cmd/dtemplate.go
	if [ ! -e /usr/bin/dtemplate ]; then \
		sudo ln -s `pwd`/bin/dtemplate /usr/bin/dtemplate;	\
	fi

src/xmlparse/xmlparser.go: src/xmlparse/xmlparser.rl
	ragel -o src/xmlparse/xmlparser.go -Z src/xmlparse/xmlparser.rl

clean:
	-@rm -rf bin/$(DEST)

prepare:
	export GOPATH=`pwd`; \
	go get -u github.com/kardianos/govendor; \
	go get -u github.com/spf13/cobra/cobra; \
	go get -u gopkg.in/yaml.v3; \
	pushd src/$(PKG); \
	../../bin/govendor init; \
	../../bin/govendor sync; \
	popd

.PHONE: clean prepare
