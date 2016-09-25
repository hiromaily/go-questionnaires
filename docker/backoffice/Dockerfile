# Dcokerfile for go-gin-wrapper

FROM node:6.6

RUN mkdir -p /usr/src/back-office
RUN mkdir /var/log/questionnaire/

#COPY docker-entrypoint.sh /go/src/github.com/hiromaily/go-questionnaires/api/docker-entrypoint.sh
COPY docker-entrypoint.sh /docker-entrypoint.sh

WORKDIR /usr/src/back-office

#RUN go get -d -v ./... && go build -v -o /go/bin/questionnaires ./cmd/server/

EXPOSE 8082
