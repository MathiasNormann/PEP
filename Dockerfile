FROM golang:latest

WORKDIR /go/src/app

# Init golang
RUN go mod init main
RUN go get golang.org/x/text/encoding/charmap

# Copy files
COPY PEP.go .
COPY PEP_listen.csv ..

# Build program
RUN go build PEP.go

# Run test
ENTRYPOINT ["/bin/bash", "-l", "-c", "./PEP"]
