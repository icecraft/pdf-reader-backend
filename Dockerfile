FROM docker.shlab.tech/public/golang:1.17-alpine as builder
WORKDIR /app
COPY . .
RUN GOPROXY=https://goproxy.cn CGO_ENABLED=0 go build -v -o main cmd/server/main.go

FROM docker.shlab.tech/public/ubuntu:20.04-zh-tools
COPY --from=builder /app/main /server
CMD ["/server"]
