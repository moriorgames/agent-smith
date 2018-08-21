# Build command:
# docker build -t moriorgames/agent-smith .
# Run command:
# docker run -td --name agent_smith -p 9090:9090 moriorgames/agent-smith --mount "type=bind,source=/var/run/docker.sock,target=/var/run/docker.sock"
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

# Compile application in a single binary
RUN         go build -o public/main public/main.go

# Expose ports and run
EXPOSE      9090
WORKDIR     /go/src/app/public
ENTRYPOINT  ["./main"]
