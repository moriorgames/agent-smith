Agent Smith
===========

A server to perform shipping containers on AWS with a web tool.
It can be considered as a cloud agent, the last step after a continuous integration pipeline has been built.

# Build, Run and Test

```
# Build the project
$ go build -o public/main public/main.go
$ go build -o public/main public/main.go && cd public && ./main
# Testing the project
$ go test ./src -v
# Golang testing in short mode
$ go test ./src -v -short
$ go test ./src -v -run TestIsAbleToSetKeyOnRedis
$ go test ./src -v -short | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
$ go test ./src -v -run TestIsAbleToSetKeyOnRedis | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
# Colors on tests: | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
# Code coverage on go
# REPLACE: _/Users/morior/Development/agent-smith TO github.com/moriorgames/agent-smith
$ go test ./src -v -short -coverprofile=coverage/output.out
$ sed -i -e 's/_\/Users\/morior\/Development/github.com\/moriorgames/g' coverage/output.out
$ go tool cover -html=coverage/output.out -o coverage/output.html 
```
