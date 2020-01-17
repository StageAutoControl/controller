FROM golang:1.10 as builder

RUN apt-get update \
    && apt-get install -y \
        libportmidi-dev \
        portaudio19-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/StageAutoControl/controller/

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o bin/controller .

FROM ubuntu

RUN apt-get update \
    && apt-get install -y \
        libportmidi-dev \
        portaudio19-dev \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/
COPY --from=builder /go/src/github.com/StageAutoControl/controller/bin/controller ./controller
RUN chmod +x ./controller
ENTRYPOINT ["./controller"]
