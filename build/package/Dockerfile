FROM golang:1.18-bullseye

RUN go install github.com/cosmtrek/air@latest && \
      go install github.com/swaggo/swag/cmd/swag@latest && \
      go install github.com/fzipp/gocyclo/cmd/gocyclo@latest && \
      go install -v github.com/quasilyte/go-ruleguard/cmd/ruleguard@latest  && \
      curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.45.2 # Install golangci-lint

CMD air
