# Build command:
# docker build -t moriorgames/agent-lib .
# Run command:
# docker run -td --name agent_smith -p 9090:9090 moriorgames/agent-lib
FROM        golang:1.10.3-alpine3.8
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Create Application directory
RUN         mkdir -p /app
COPY        . /app
WORKDIR     /app

# Install needed packages for GO
RUN         apk update && apk upgrade && \
            apk add --no-cache bash git \
            && go get -u github.com/gorilla/mux \
            && apk del git

# Compile application in a single binary
RUN         go build -o bin/main src/*.go

# Expose ports
EXPOSE      9090

ENTRYPOINT  ["./bin/main"]
