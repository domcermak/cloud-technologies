FROM golang:1.17.8-alpine3.15

LABEL version="0.1"

ENV BASE_PATH=${GOPATH}/src
ENV APP_PATH=${BASE_PATH}/assignments/03

RUN mkdir -p ${APP_PATH}

COPY go.mod ${BASE_PATH}
COPY go.sum ${BASE_PATH}
COPY ./assignments/03/main.go ${APP_PATH}
COPY ./assignments/03/cmd ${APP_PATH}/cmd

WORKDIR ${BASE_PATH}

RUN apk --no-cache add ca-certificates git
RUN go mod download

WORKDIR ${APP_PATH}

CMD go run server.go -server_address="${SERVER_ADDR}" \
                     -pg_host="${PG_HOST}" \
                     -pg_port=${PG_PORT} \
                     -pg_database="${PG_DATABASE}" \
                     -pg_username="${PG_USERNAME}" \
                     -pg_password="${PG_PASSWORD}"