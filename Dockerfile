FROM golang:latest

WORKDIR /go/src/app

COPY ./app.go /go/src/app
COPY go.mod /go/src/app
COPY go.sum /go/src/app

RUN go build -o app .

ENTRYPOINT [ "./app" ]