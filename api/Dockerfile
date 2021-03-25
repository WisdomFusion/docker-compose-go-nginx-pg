FROM golang:1.16-alpine AS build
WORKDIR /go/src/github.com/org/repo
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o server .

FROM alpine
EXPOSE 8000
COPY --from=build /go/src/github.com/org/repo/server /server
CMD ["/server"]
