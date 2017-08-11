# Note: tabs by space can't not used for Makefile!

###############################################################################
# Docker
###############################################################################
dcbld:
	docker-compose build

dcbldfull:
	docker-compose build --no-cache

dcup:
	docker-compose up


###############################################################################
# Local
###############################################################################
bld:
	go build -i -v -o ${GOPATH}/bin/apiserver ./cmd/api-server/

run:
	go run ./cmd/api-server/main.go -docker 0


###############################################################################
# Godeps
###############################################################################
godep:
	#go get -u github.com/tools/godep
	go get -d -v -u ./api/cmd/api-server/
	rm -rf ./api/Godeps
	rm -rf ./api/vendor
	cd ./api;godep save ./...


###############################################################################
# Test by curl
###############################################################################
curl:
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

	#Answer
	#[GET]
	#http://localhost:8083/api/answer/1
	curl localhost:8083/api/answer/1

	#[POST]
	#http://localhost:8083/api/answer/1
	curl -v -H "Accept: application/json" -H "Content-type: application/json" \
	-X POST -d '{"email":"aaa@bbb.ccc", "answers":["a1","a2","a3"]}' \
	http://localhost:8084/api/answer/1

