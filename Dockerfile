FROM golang:1.17-alpine as gin
WORKDIR /app
COPY . .
ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"
RUN go build main.go

FROM alpine:latest
WORKDIR /app
COPY --from=gin /app/main .
CMD [ "./main" ]
