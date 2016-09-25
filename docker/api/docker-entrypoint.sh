#!/bin/sh
###
# initialize for docker environment
###

#go get -d -v ./...
go get -u github.com/tools/godep
go build -v -o /go/bin/questionnaires ./cmd/api-server/

questionnaires -docker 1
