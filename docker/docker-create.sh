#!/bin/sh

###############################################################################
# Using docker-composer for go-questionnaires
###############################################################################
#echo $RUN_TEST

###############################################################################
# Environment
###############################################################################
CONTAINER_WEB=question-web
CONTAINER_MYSQL=question-mysql
IMAGE_NAME=go-questionnaires:1.0


###############################################################################
# Remove Container And Image
###############################################################################
DOCKER_PSID=`docker ps -af name="${CONTAINER_WEB}" -q`
if [ ${#DOCKER_PSID} -ne 0 ]; then
    docker rm -f ${CONTAINER_WEB}
fi

DOCKER_PSID=`docker ps -af name="${CONTAINER_MYSQL}" -q`
if [ ${#DOCKER_PSID} -ne 0 ]; then
    docker rm -f ${CONTAINER_MYSQL}
fi

#docker rm -f $(docker ps -aq)

DOCKER_IMGID=`docker images "${IMAGE_NAME}" -q`
if [ ${#DOCKER_IMGID} -ne 0 ]; then
    docker rmi ${IMAGE_NAME}
fi


###############################################################################
# Docker-compose / build and up
###############################################################################
docker-compose  build
docker-compose  up -d

# run server mode
# foreground
#docker exec -it ${CONTAINER_NAME} bash ./docker-entrypoint.sh

# background(trying)
docker exec -itd ${CONTAINER_NAME} bash ./docker-entrypoint.sh

###############################################################################
# Docker-compose / check
###############################################################################
sleep 3s

docker-compose ps
docker-compose logs


###############################################################################
# Exec
###############################################################################
#docker exec -it web bash


###############################################################################
# Test
###############################################################################



###############################################################################
# Docker-compose / down
###############################################################################
#docker-compose -f ${COMPOSE_FILE} down

###############################################################################
# Check connection
###############################################################################
#mysql -u root -p -h 127.0.0.1 -P 13306
#redis-cli -h 127.0.0.1 -p 16379 -a password

#Access by browser
#http://docker.hiromaily.com:9999/
