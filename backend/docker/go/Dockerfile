FROM golang:1.17

WORKDIR /go/src/app

EXPOSE 8080

COPY . ./

RUN go get -d -v ./... && go install -v ./...

CMD ["go", "run", "main.go"]