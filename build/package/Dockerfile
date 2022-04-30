FROM golang:1.18-bullseye

RUN go install github.com/cosmtrek/air@latest
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2

CMD air
