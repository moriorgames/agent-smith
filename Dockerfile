# Build command:
# docker build -t moriorgames/agent-smith .
# Run command:
# docker run -td --name agent_smith -p 9090:9090 moriorgames/agent-smith
FROM        golang:1.10.3-alpine3.8
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Create Application directory
RUN         mkdir -p /app
COPY        . /app
WORKDIR     /app

# Install needed packages for GO
RUN         apk update \
            && apk upgrade \
            && apk add --no-cache bash git \
            && go get -u github.com/gorilla/mux \
            && go get -u github.com/oxtoacart/bpool \
            && go get -u github.com/moriorgames/agent-smith/src \
            && go get -u github.com/go-redis/redis \
            && go get -u github.com/stretchr/testify/assert \
            && apk del git

# Compile application in a single binary
RUN         go build -o public/main public/main.go

# Expose ports and run
EXPOSE      9090
WORKDIR     /app/public
ENTRYPOINT  ["./main"]
