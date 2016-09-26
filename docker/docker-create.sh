#!/bin/sh

###############################################################################
# Using docker-composer for go-questionnaires
###############################################################################
#echo $RUN_TEST

###############################################################################
# Environment
###############################################################################
#Container
CONTAINER_API_SERVER=question-api
CONTAINER_BACK_SERVER=question-backoffice
CONTAINER_FRONT_SERVER=question-frontoffice

CONTAINER_MYSQL=question-mysql
CONTAINER_NGINX=question-nginx

#Image
#IMAGE_NAME=node-question-backoffice:1.0
#IMAGE_NAME2=node-question-frontoffice:1.0
#IMAGE_NAME3=go-question-api:1.0


###############################################################################
# Remove Container And Image
###############################################################################
for con in $CONTAINER_API_SERVER $CONTAINER_BACK_SERVER $CONTAINER_FRONT_SERVER $CONTAINER_MYSQL $CONTAINER_NGINX; do
    DOCKER_PSID=`docker ps -af name="${con}" -q`
    if [ ${#DOCKER_PSID} -ne 0 ]; then
        docker rm -f ${con}
    fi
done
#docker rm -f $(docker ps -aq)

#DOCKER_IMGID=`docker images "${IMAGE_NAME}" -q`
#if [ ${#DOCKER_IMGID} -ne 0 ]; then
#    docker rmi ${IMAGE_NAME}
#fi


###############################################################################
# Docker-compose / build and up
###############################################################################
docker-compose build
docker-compose up -d


# background
docker exec -itd ${CONTAINER_API_SERVER} bash /docker-entrypoint.sh
docker exec -itd ${CONTAINER_BACK_SERVER} bash /docker-entrypoint.sh
docker exec -itd ${CONTAINER_FRONT_SERVER} bash /docker-entrypoint.sh

###############################################################################
# Docker-compose / check
###############################################################################
sleep 3s

docker-compose ps
docker-compose logs


###############################################################################
# Exec
###############################################################################
#docker exec -it question-api bash

#questionnaires -docker 1
#ps -aux


#docker exec -it question-backoffice bash
#docker exec -it question-frontoffice bash




###############################################################################
# Docker-compose / down
###############################################################################
#docker-compose down

###############################################################################
# Check connection
###############################################################################
#mysql -u root -p -h 127.0.0.1 -P 4306

#Access by browser
#http://localhost:8080/api/ques
#http://localhost:8081/
#http://localhost:8082/admin/
#http://localhost:8083/api/ques
