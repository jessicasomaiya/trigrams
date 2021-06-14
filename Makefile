MAIN_DIR = cmd


run: 
	go run $(MAIN_DIR)/main.go

test: 
	go test -v ./...

local: test run