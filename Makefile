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
# Amazon EC2 Container Service (ECS) Create Images
###############################################################################
#TAG=latest
TAG=1.0

IMAGE_NGINX=hirokiy/question-nginx:${TAG}
IMAGE_MYSQL=hirokiy/question-mysql:${TAG}
IMAGE_API=hirokiy/question-api:${TAG}
IMAGE_FRONT=hirokiy/question-front:${TAG}
IMAGE_BACK=hirokiy/question-back:${TAG}

ecs_create_image:
    # Nginx
    docker rmi ${IMAGE_NGINX}
    docker build -t ${IMAGE_NGINX} ./docker/nginx

    # MySQL
    docker rmi ${IMAGE_MYSQL}
    docker build -t ${IMAGE_MYSQL} ./docker/mysql

    # API
    docker rmi ${IMAGE_API}
    docker build -t ${IMAGE_API} ./docker/api
    #docker run -it --name go-api --link mysqld:mysql-server \
    # -p 8083:8083 -d ${IMAGE_API}

    # Front
    docker rmi ${IMAGE_FRONT}
    docker build -t ${IMAGE_FRONT} ./docker/frontoffice

    # Back
    docker rmi ${IMAGE_BACK}
    docker build -t ${IMAGE_BACK} ./docker/backoffice


ecs_push_image:
    docker push ${IMAGE_NGINX}
    docker push ${IMAGE_MYSQL}
    docker push ${IMAGE_API}
    docker push ${IMAGE_FRONT}
    docker push ${IMAGE_BACK}


###############################################################################
# Amazon EC2 Container Service (ECS)
###############################################################################

#aws configure
# Default region name [None]: ap-northeast-1
# Default output format [None]: json
ecs_init:
	pip install awscli

ecs_install:
	sudo curl -o /usr/local/bin/ecs-cli https://s3.amazonaws.com/amazon-ecs-cli/ecs-cli-darwin-amd64-latest
	sudo chmod +x /usr/local/bin/ecs-cli
	ecs-cli help

ecs_config:
	ecs-cli configure --region ap-northeast-1 --cluster hy-cluster

#ecs_create_key_pair:
	#http://docs.aws.amazon.com/ja_jp/AmazonCloudFront/latest/DeveloperGuide/IIS4.1CreatingAnEC2KeyPair.html
	#https://ap-northeast-1.console.aws.amazon.com/ec2/v2/home?region=ap-northeast-1#KeyPairs:sort=keyName
	#~/.aws/hy-key.pem

#ecs_create_cluster:
	#ecs-cli up --capability-iam --keypair hy-key
	#ecs-cli up --capability-iam --keypair hy-key --size 5
	#-> Couldn't run containers

	#ecs-cli up --capability-iam --keypair hy-key --size 5 --instance-type t2.small
	#-> Couldn't run containers

	#ecs-cli up --capability-iam --keypair hy-key --size 5 --instance-type t2.medium
	#->It works well

	#Defaults to t2.micro
	#ecs-cli up --capability-iam --keypair hy-key --instance-type t2.small --size 5
	#ecs-cli up --capability-iam --keypair hy-key --instance-type t2.medium --size 5

	#scale
	#ecs-cli scale --capability-iam --size 5

# As Task
ecs_run_container:
	# Run container
	ecs-cli compose --file docker-compose-aws.yml up
	#ecs-cli compose --file docker-compose-aws.yml scale 1

	#->Couldn't run containers     reason=RESOURCE:MEMORY
	#-->ecs-cli scale --capability-iam --size 2

ecs_check_container:
	# Check container
	ecs-cli ps

ecs_stop:
	# Stop container
	ecs-cli compose --file docker-compose-aws.yml down

ecs_cleanup:
	#ecs-cli compose --file docker-compose-aws.yml down
	#ecs-cli compose --file docker-compose-aws.yml service rm
	ecs-cli down --force

# As Service
ecs_run_container2:
	# Run container
	ecs-cli compose --file docker-compose-aws.yml service up

ecs_check_container2:
	# Check container
	ecs-cli compose --file docker-compose-aws.yml service ps

ecs_stop:
	# Stop(Clean)
	ecs-cli compose --file docker-compose-aws.yml service rm


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

