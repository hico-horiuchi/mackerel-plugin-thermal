#!/bin/bash

VERSION=1.0.0
GOLANG_VERSION=1.16.6
NAME=mackerel-plugin-thermal

docker run -i --rm -v $(pwd):/usr/src/$NAME -w /usr/src/$NAME golang:$GOLANG_VERSION bash <<EOS
go get -v -u github.com/Songmu/goxz/cmd/goxz
go mod tidy
goxz -d dist/v${VERSION} -z -os linux -arch amd64,arm64
EOS
