FROM golang:1.24
WORKDIR /app

COPY go.mod ./
COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd

ENV PORT=8080
CMD ["/app/server"]