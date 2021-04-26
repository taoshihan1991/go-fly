FROM golang:alpine as builder
WORKDIR /app
COPY . /app
ENV GOPROXY https://mirrors.aliyun.com/goproxy
VOLUME ["/app/config"]
RUN go build go-fly.go

FROM golang:alpine
WORKDIR /app
COPY --from=builder app/go-fly ./
EXPOSE 8081
ENTRYPOINT ["/app/go-fly","server"]
