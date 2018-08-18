# Build command:
# docker build -t moriorgames/agent-smith .
# Run command:
# docker run -td --name agent_smith -p 9090:9090 moriorgames/agent-smith --mount "type=bind,source=/var/run/docker.sock,target=/var/run/docker.sock"
FROM        moriorgames/agent-smith-base:v2
MAINTAINER  MoriorGames "moriorgames@gmail.com"

# Install current package files
RUN         go get -u github.com/moriorgames/agent-smith/src

# Create Application directory
RUN         mkdir -p /app
COPY        . /app
WORKDIR     /app

# Compile application in a single binary
RUN         go build -o public/main public/main.go

# Expose ports and run
EXPOSE      9090
WORKDIR     /app/public
ENTRYPOINT  ["./main"]
