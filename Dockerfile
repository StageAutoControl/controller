FROM instrumentisto/glide:0.13.0 as builder
WORKDIR /go/src/github.com/StageAutoControl/controller/

COPY glide.yaml glide.lock ./
RUN glide install --strip-vendor

COPY . ./

RUN go test ./...

RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o bin/controller_amd64 .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/StageAutoControl/controller/bin/controller_amd64 .
RUN chmod +x ./controller
ENTRYPOINT ["./controller"]
