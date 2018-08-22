# Build command:
# docker build -t moriorgames/agent-smith-base -f docker/Dockerfile  .
FROM        golang:1.10.3-alpine3.8
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# Install needed packages for GO
RUN         apk update \
            && apk upgrade \
            && apk add --no-cache bash git curl

# build directories
RUN         mkdir /go/src/app
ADD         . /go/src/app
WORKDIR     /go/src/app

# Go dependency package manager
RUN         curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN         dep ensure -vendor-only
