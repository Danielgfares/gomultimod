EXECUTABLE=$(shell basename $(CURDIR))
# WINDOWS=$(EXECUTABLE)_windows_amd64.exe
# LINUX_AMD=$(EXECUTABLE)_linux_amd64
LINUX_ARM=$(EXECUTABLE)_linux_arm64
# DARWIN=$(EXECUTABLE)_darwin_amd64
VERSION=$(shell git describe --tags --always --long --dirty)

# .PHONY: all test clean

# all: test build

# test:
# 	go test ./...

build: linux
	@echo version: $(VERSION)

# windows: $(WINDOWS)

linux: $(LINUX_ARM) #$(LINUX_AMD) 

# darwin: $(DARWIN)

# $(WINDOWS):
# 	env GOOS=windows GOARCH=amd64 go build -v -o $(WINDOWS) -ldflags="-s -w -X main.version=$(VERSION)" ./main.go

# $(LINUX_AMD):
# 	env GOOS=linux GOARCH=amd64 go build -v -o $(LINUX_AMD) -ldflags="-s -w -X main.version=$(VERSION)" ./main.go

$(LINUX_ARM): # GOARCH=arm GOARM=7
	env GOOS=linux go build -v -o $(LINUX_ARM) -ldflags="-s -w -X main.version=$(VERSION)" ./main.go

# $(DARWIN):
# 	env GOOS=darwin GOARCH=amd64 go build -v -o $(DARWIN) -ldflags="-s -w -X main.version=$(VERSION)" ./main.go

clean:
	rm -f $(LINUX_ARM)

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
