FROM golang:1.7.3
WORKDIR /go/src/github.comn/ziggy192/tasa-vietnam-api/
COPY . . 
RUN go get ./... && cd app && go build
CMD ["./tasa-vietnam-api"]
