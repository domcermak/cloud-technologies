FROM golang:1.17.8-alpine3.15

LABEL version="0.1"

ENV BASE_PATH=${GOPATH}/src
ENV APP_PATH=${BASE_PATH}/assignments/03

RUN mkdir -p ${APP_PATH}

COPY go.mod ${BASE_PATH}
COPY go.sum ${BASE_PATH}
COPY ./assignments/03/server.go ${APP_PATH}
COPY ./assignments/03/cmd ${APP_PATH}/cmd

WORKDIR ${BASE_PATH}

RUN apk --no-cache add ca-certificates git
RUN go mod download

WORKDIR ${APP_PATH}

EXPOSE 1234

# replacing 0.0.0.0 with localhost ends up in empty response error
# because the server is not listening on the correct host
# https://stackoverflow.com/questions/52524289/docker-empty-response-from-server

# Using only CMD because CMD is ommited when combined with ENTRYPOINT
# https://stackoverflow.com/questions/54447913/cmd-doesnt-run-after-entrypoint-in-dockerfile
CMD go run server.go \
    -server_address="0.0.0.0:1234" \
    -pg_host="${PG_HOST}" \
    -pg_port=${PG_PORT} \
    -pg_database="${PG_DATABASE}" \
    -pg_username="${PG_USERNAME}" \
    -pg_password="${PG_PASSWORD}" \
    -log_level="${LOG_LEVEL}"