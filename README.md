# agent-smith

A server to perform container deploy on aws. It can be considered as a cloud agent, the last step after a pipeline has been built.

# Build, Run and Test

```
# Build the project
$ go build -o public/main public/main.go
$ go build -o public/main public/main.go && cd public && ./main
# Run the project
$ go run public/maingo
# Testing the project
$ go test ./src -v
```
