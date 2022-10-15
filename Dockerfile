FROM golang:latest

WORKDIR /usr/src/app

RUN go version
ENV GOPATH=/

COPY . .
RUN go mod download
RUN go build -o main .
CMD ["./main"]