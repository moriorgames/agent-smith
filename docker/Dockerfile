# Build command:
# docker build -t moriorgames/agent-smith-base -f docker/Dockerfile  .
FROM        golang:1.10.3-alpine3.8
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Install needed packages for GO
RUN         apk update \
            && apk upgrade \
            && apk add --no-cache bash git curl

# Go dependency package manager
RUN         curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
