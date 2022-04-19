FROM golang:1.17.8-alpine3.15

LABEL version="0.1"

EXPOSE 8080

ENV BASE_PATH=${GOPATH}/src
ENV APP_PATH=${BASE_PATH}/assignments/05

COPY go.mod ${BASE_PATH}
COPY go.sum ${BASE_PATH}
COPY ./assignments/05 ${APP_PATH}

WORKDIR ${BASE_PATH}

RUN apk --no-cache add ca-certificates git
RUN go mod download

WORKDIR ${APP_PATH}

CMD go run server.go \
    -server_address=":8080" \
    -etcd_server_address="${ETCD_HOST}:${ETCD_PORT}"