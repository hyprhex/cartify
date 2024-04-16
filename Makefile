run: 
	@air --build.cmd "go build -o bin/app cmd/main.go" --build.bin "./bin/app"

test: 
	@go test -v ./...
