#!/bin/sh
###
# initialize for docker environment
###

go get -d -v ./...
go build -v -o /go/bin/questionnaires ./cmd/server/

questionnaires -f ./configs/settings.toml
