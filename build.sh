#!/bin/bash
set -e
export GOPATH=`pwd`
for p in \
	'github.com/golang/glog' \
	; do
	echo $p; 
	if [ ! -d src/$p ]; then
		go get $p
	fi
done
if [ ! -d bin ]; then
	mkdir bin
fi
go build -o bin/dtemplate src/cmd/dtemplate.go
if [ ! -e /usr/bin/dtemplate ]; then 
	sudo ln -s `pwd`/bin/dtemplate /usr/bin/dtemplate
fi

