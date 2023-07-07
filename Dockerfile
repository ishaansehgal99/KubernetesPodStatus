FROM golang:latest

WORKDIR /go/src/app

COPY ./app.go /go/src/app

RUN go build -o app .

ENTRYPOINT [ "./app" ]