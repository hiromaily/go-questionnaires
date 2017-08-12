# Dcokerfile for api-server

FROM golang:1.8

RUN mkdir -p /go/src/github.com/hiromaily/go-questionnaires/api/cmd/api-server && \
mkdir -p /go/src/github.com/hiromaily/go-questionnaires/api/tmp/log \
mkdir -p /var/log/questionnaire/ && \
mkdir -p /var/log/go/

#COPY docker-entrypoint.sh /go/src/github.com/hiromaily/go-questionnaires/api/docker-entrypoint.sh
#COPY docker-entrypoint.sh /docker-entrypoint.sh

WORKDIR /go/src/github.com/hiromaily/go-questionnaires/api
COPY . .

RUN go get -u github.com/tools/godep && go get -u github.com/hiromaily/fresh
#RUN go get -d -v -u ./cmd/api-server/
RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/questionnaires ./cmd/api-server/main.go

EXPOSE 8083
CMD ["/go/bin/questionnaires", "-docker", "1"]
