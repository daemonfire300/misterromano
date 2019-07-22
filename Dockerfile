FROM golang:1.12
EXPOSE 8080

ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .

RUN groupadd -g 1453 romansvc && \
    useradd -r -u 1453 -g romansvc romansvc
RUN chown romansvc:romansvc -R . && mkdir /home/romansvc && chown romansvc:romansvc -R /home/romansvc
USER romansvc

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build -v -o app .

CMD ["./app"]