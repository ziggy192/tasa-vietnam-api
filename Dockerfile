FROM golang:1.13.4
WORKDIR /go/src/github.comn/ziggy192/tasa-vietnam-api/
COPY . . 

RUN go get ./...  && go build

CMD ["./tasa-vietnam-api","-address=54.169.62.76:3306"]
