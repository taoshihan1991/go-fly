FROM golang:alpine
WORKDIR /app
COPY . /app
RUN go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct
VOLUME ["/app/config"]
RUN go build go-fly.go
EXPOSE 8081
CMD ["/app/go-fly","server"]