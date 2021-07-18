SHELL=/bin/bash
SRC=$(shell find  . -name "*.go")

../../bin/dtemplate: $(SRC)
	go build -o ../../bin/dtemplate cmd/dtemplate.go

