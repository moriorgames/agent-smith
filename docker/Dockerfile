# Build command:
# docker build -t moriorgames/agent-smith-base -f docker/Dockerfile  .
FROM        golang:1.10.3-alpine3.8
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Install needed packages for GO
RUN         apk update \
            && apk upgrade \
            && apk add --no-cache bash git \
            && go get -u github.com/docker/docker/api/types \
            && go get -u github.com/docker/docker/client \
            && go get -u github.com/moriorgames/agent-smith/src \
            && go get -u github.com/go-redis/redis \
            && go get -u github.com/gorilla/mux \
            && go get -u github.com/oxtoacart/bpool \
            && go get -u github.com/satori/go.uuid \
            && go get -u github.com/stretchr/testify/assert \
            && apk del git

