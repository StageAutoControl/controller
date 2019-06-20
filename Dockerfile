FROM scalify/glide:0.13.0 as dependencies
WORKDIR /go/src/github.com/StageAutoControl/controller/

COPY glide.yaml glide.lock ./
RUN glide install --strip-vendor

FROM golang:1.10 as builder

RUN apt-get update \
    && apt-get install -y \
        libportmidi-dev \
        portaudio19-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/StageAutoControl/controller/
COPY --from=dependencies /go/src/github.com/StageAutoControl/controller/vendor vendor

COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/controller .

FROM ubuntu

RUN apt-get update \
    && apt-get install -y \
        libportmidi-dev \
        portaudio19-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/
COPY --from=builder /go/src/github.com/StageAutoControl/controller/bin/controller_amd64 ./controller
RUN chmod +x ./controller
ENTRYPOINT ["./controller"]
