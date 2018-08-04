Agent Smith
===========

A server to perform shipping containers on AWS with a web tool.
It can be considered as a cloud agent, the last step after a continuous integration pipeline has been built.

# Build, Run and Test

```
# Build the project
$ go build -o public/main public/main.go
$ go build -o public/main public/main.go && cd public && ./main
# Run the project
$ go run public/maingo
# Testing the project
$ go test ./src -v
# Golang testing in short mode
$ go test ./src -v -short
$ go test ./src -v -run TestIsAbleToSetKeyOnRedis
$ go test ./src -v -short | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
$ go test ./src -v -run TestIsAbleToSetKeyOnRedis | sed ''/PASS/s//$(printf "\033[32mPASS\033[0m")/'' | sed ''/FAIL/s//$(printf "\033[31mFAIL\033[0m")/''
```
