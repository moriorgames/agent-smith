# Build command:
# docker build -t moriorgames/agent-smith .
# Run command:
# docker run -td --name agent_smith -p 9090:9090 moriorgames/agent-smith
FROM        moriorgames/agent-smith-base:v1
MAINTAINER  MoriorGames "moriorgames@gmail.com"

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
