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

start-playback-visualizer: build-darwin
	./bin/controller_darwin playback song "${SONG}" \
		--data-dir "${SAC_DATA_DIR}" \
		--transport visualizer \
		--visualizer-endpoint localhost:1337

start-playback-stream: build-darwin
	./bin/controller_darwin playback song "${SONG}" \
			--data-dir "${SAC_DATA_DIR}" \
			--transport stream

start-playback-none: build-darwin
	./bin/controller_darwin playback song "${SONG}" \
			--data-dir "${SAC_DATA_DIR}"

start-playback-artnet: build-darwin
	./bin/controller_darwin playback song "${SONG}" \
			--data-dir "${SAC_DATA_DIR}" \
			--transport artnet

start-api: build-darwin
	./bin/controller_darwin api \
	--data-dir "${SAC_DATA_DIR}"

build-all: build-darwin build-arm build-linux build-docker

build-darwin:
	go build -o bin/controller_darwin .

build-linux: 
	GOOS=linux CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/controller_linux .

build-arm:
	GOOS=linux GOARCH=arm GOARM=6 CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/controller_arm .

build-docker: build-linux
	docker build -t stageautocontrol/controller .
