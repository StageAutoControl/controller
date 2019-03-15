PACKAGES:=$$(go list ./... | grep -v /vendor/)

.PHONY: build build build-arm build-linux build-docker test fmt all

all: build-all

install:
	@glide install --strip-vendor

update:
	@glide update --strip-vendor

fmt:
	@go fmt $(PACKAGES)

lint:
	@golint ./... | grep -vE "vendor|\.pb\.go" || printf ""

test:
	go test -v $(PACKAGES)

build-all: build-darwin build-arm build-linux build-docker

build-darwin:
	go build -o bin/controller_darwin .

build-linux:
	GOOS=linux CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/controller_linux .

build-arm:
	GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/controller_arm .

build-docker: build-linux
	docker build -t stageautocontrol/controller .
