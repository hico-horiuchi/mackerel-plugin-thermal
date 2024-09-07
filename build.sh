#!/bin/bash

set -eu -o pipefail

go install github.com/Songmu/goxz/cmd/goxz@latest
go install github.com/tcnksm/ghr@latest

go mod tidy

goxz -d dist/${RELEASE_TAG} -z -os linux -arch amd64,arm64
ghr -u hico-horiuchi -r mackerel-plugin-thermal ${RELEASE_TAG} dist/${RELEASE_TAG}
