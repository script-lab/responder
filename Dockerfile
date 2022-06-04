FROM golang:1.18.3

WORKDIR /go/src/github.com/script-lab/responder
COPY go.mod .
COPY go.sum .

RUN go install github.com/cosmtrek/air@latest

COPY . .

CMD ["go", "run", "*.go"]