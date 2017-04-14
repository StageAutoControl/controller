PACKAGES:=$$(go list ./... | grep -v /vendor/)

.phony: build test fmt

build:
	go build -o controller .

install:
	@glide install --strip-vendor

update:
	@glide update --strip-vendor

fmt:
	@go fmt $(PACKAGES)

lint:
	@golint ./... | grep -vE "vendor" || printf ""

test:
	@go test ${PACKAGES}

proto:
	protoc -I "cntl/transport" --go_out="cntl/transport" cntl/transport/dmx.proto

start-visualizer: build
	./controller playback \
		--data-dir ~/src/github.com/apinnecke/omw-sac-data/ \
		--transport visualizer \
		--visualizer-endpoint localhost:1337 \
		some-song-uuid-2

start-buffer: build
	./controller playback \
        --data-dir ~/src/github.com/apinnecke/omw-sac-data/ \
        some-song-uuid-2
