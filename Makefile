PACKAGES:=$$(go list ./... | grep -v /vendor/)

.phony: build test fmt

build:
	@mkdir -p bin
	go build -o bin/controller .

install:
	@glide install --strip-vendor

update:
	@glide update --strip-vendor

fmt:
	@go fmt $(PACKAGES)

lint:
	@golint ./... | grep -vE "vendor|\.pb\.go" || printf ""

test:
	@go test ${PACKAGES}

proto:
	protoc -I "cntl/transport" --go_out="cntl/transport" cntl/transport/dmx.proto

start-playback-visualizer: build
	./bin/controller playback \
		--data-dir "$${SAC_DATA_DIR}" \
		--transport visualizer \
		--visualizer-endpoint localhost:1337 \
		"$${1}"

start-playback-buffer: build
	./bin/controller playback \
        	--data-dir "$${SAC_DATA_DIR}" \
        	"$${1}"

start-api: build
	./bin/controller api \
	--data-dir "$${SAC_DATA_DIR}"
