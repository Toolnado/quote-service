FROM golang:1.16-buster

RUN go version

ENV GOPATH=/ 

COPY ./ ./

RUN go mod download 
RUN go build -o quote-service.exe ./cmd/main.go 

CMD ["./quote-service.exe"]