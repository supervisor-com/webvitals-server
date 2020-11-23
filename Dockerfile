FROM golang:1.15-alpine as builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o webvitals-server cmd/webvitals-server/main.go

FROM alpine:3.12

WORKDIR /app
COPY --from=builder /build/webvitals-server /usr/bin
COPY --from=builder /build/views/ /app/views/
COPY --from=builder /build/assets/ /app/assets/

ENV GIN_MODE=release
ENTRYPOINT [ "/usr/bin/webvitals-server" ]
