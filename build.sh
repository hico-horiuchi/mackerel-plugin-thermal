#!/bin/bash

VERSION=$(awk '/VERSION/{gsub(/"/,"");print $NF}' lib/version.go)
GOLANG_VERSION=1.19.3
NAME=mackerel-plugin-thermal

docker run -i --rm -v $(pwd):/usr/src/$NAME -w /usr/src/$NAME golang:$GOLANG_VERSION bash <<EOS
go install github.com/Songmu/goxz/cmd/goxz@latest
go mod tidy
goxz -pv $VERSION -z -os linux -arch amd64,arm64
EOS
