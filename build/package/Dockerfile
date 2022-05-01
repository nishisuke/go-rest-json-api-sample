FROM golang:1.18-bullseye

RUN go install github.com/cosmtrek/air@latest && \
      go install github.com/swaggo/swag/cmd/swag@latest && \
      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2 # Install golangci-lint

CMD air
