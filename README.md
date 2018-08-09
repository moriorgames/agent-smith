Agent Smith
===========

A server to perform shipping containers on AWS with a web tool.
It can be considered as a cloud agent, the last step after a continuous integration pipeline has been built.

# Build, Run and Test

Build the project
```
$ go build -o public/main public/main.go
$ go build -o public/main public/main.go && cd public && ./main
```

Testing the project. You have different options to test the project.
```
$ go test ./src -v
```

Golang testing in short mode
```
$ go test ./src -v -short
```

Or testing specific function
```
$ go test ./src -v -run TestIsAbleToSetKeyOnRedis
```

Decorating the output
```
# Colors on tests: | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
```

Code coverage on go
```
# REPLACE: _/Users/morior/Development/agent-smith TO github.com/moriorgames/agent-smith
$ go test ./src -v -short -coverprofile=coverage/output.out
$ sed -i -e 's/_\/Users\/morior\/Development/github.com\/moriorgames/g' coverage/output.out
$ go tool cover -html=coverage/output.out -o coverage/output.html 
```

# Docker

Run the application in docker container

```
$ docker build -t moriorgames/agent-smith .
$ docker run -td --name agent_smith -p 9090:9090 moriorgames/agent-smith
```

Docker container

```
$ docker ps
$ docker ps -a
$ docker stop agent_smith && docker rm agent_smith
```

Docker images

```
$ docker images
$ docker rmi ID_IMAGE
```

Remove volumes starting with character

```
$ docker volume ls
$ docker volume rm $(docker volume ls -q | grep "^0")
```
