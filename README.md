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
```
