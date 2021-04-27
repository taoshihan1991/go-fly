FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
ENV GOPROXY https://mirrors.aliyun.com/goproxy
VOLUME ["/app/config"]
RUN go build go-fly.go
EXPOSE 8081
CMD ["/app/go-fly","server"]