FROM golang:alpine as builder
WORKDIR /app
COPY . /app
ENV GOPROXY https://mirrors.aliyun.com/goproxy
RUN go build go-fly.go

FROM golang:alpine
WORKDIR /app
VOLUME ["/app/config"]
RUN mkdir static
COPY --from=builder /app/static ./static
COPY --from=builder /app/go-fly ./
EXPOSE 8081
ENTRYPOINT ["./go-fly","server"]
