FROM golang:1.21

WORKDIR /usr/src/app

COPY . .

RUN go build -v -o /usr/local/bin/app ./net_pro/main.go

CMD ["app"]