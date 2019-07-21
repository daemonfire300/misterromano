FROM golang:1.12
EXPOSE 8080

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -v -o app .

CMD ["./app"]