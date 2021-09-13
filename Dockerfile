FROM golang:1.17-alpine

WORKDIR /app

COPY . .

ENV GO111MODULE="on"
ENV GOPROXY="https://goproxy.cn,direct"

RUN go build main.go

CMD [ "./main" ]