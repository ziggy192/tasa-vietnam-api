FROM golang:1.13.4
EXPOSE 8000

WORKDIR /go/src/github.comn/ziggy192/tasa-vietnam-api/
COPY . . 

RUN go get ./...  && go build


