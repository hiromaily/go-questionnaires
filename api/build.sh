#!/bin/sh

###########################################################
# Variable
###########################################################
BUILD_MODE=2  #1:build 2:run
CURL_MODE=0   #0:off 1:curl
GODEP_MODE=1


###########################################################
# Check
###########################################################
echo '============== go fmt =============='
go fmt `go list ./... | grep -v '/vendor/'`
echo '============== go vet =============='
go vet `go list ./... | grep -v '/vendor/'`
EXIT_STATUS=$?
if [ $EXIT_STATUS -gt 0 ]; then
    exit $EXIT_STATUS
fi

echo '============== golint =============='
golint ./... | grep -v '^vendor\/' || true

echo '============== misspell =============='
#misspell .
misspell `find . -name "*.go" | grep -v '/vendor/'`

echo '============== ineffassign =============='
ineffassign .


###########################################################
#Build
###########################################################
if [ $GODEP_MODE -eq 1 ]; then
    rm -rf Godeps
    rm -rf vendor
fi

if [ $BUILD_MODE -eq 1 ]; then
    go build -i -v -o ${GOPATH}/bin/apiserver ./cmd/api-server/
elif [ $BUILD_MODE -eq 2 ]; then
    go run ./cmd/api-server/main.go -docker 0
fi


if [ $GODEP_MODE -eq 1 ]; then
    godep save ./...
fi

###########################################################
#Curl
###########################################################
if [ $CURL_MODE -eq 1 ]; then
    #[GET]
    #http://localhost:8083/api/ques
    curl localhost:8083/api/ques

    #[POST]
    #http://localhost:8083/api/ques
    curl -v -H "Accept: application/json" -H "Content-type: application/json" \
    -X POST -d '{"title":"title4", "questions":["q1","q2","q3"]}' \
    http://localhost:8083/api/ques


    #[DELETE]
    #http://localhost:8083/api/ques/5
    curl http://localhost:8083/api/ques/5 -X DELETE

    #[GET]
    #http://localhost:8083/api/answer/1
    curl localhost:8083/api/answer/1
fi

