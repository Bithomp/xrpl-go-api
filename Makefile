.DEFAULT_GOAL := test

.PHONY:fmt vet test
fmt:
		go fmt ./...

vet: fmt
		go vet ./...
		
test: vet
		go test ./...

update:
	go get -u ./...
	go get -u all
	go mod tidy
