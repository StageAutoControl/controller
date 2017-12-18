FROM instrumentisto/glide:0.13.0 as builder
WORKDIR /go/src/github.com/StageAutoControl/controller/

COPY glide.yaml glide.lock ./
RUN glide install --strip-vendor

COPY . ./
RUN CGO_ENABLED=0 go build -a -ldflags '-s' -installsuffix cgo -o controller .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/StageAutoControl/controller/controller .
RUN chmod +x ./controller
ENTRYPOINT ["./controller"]
