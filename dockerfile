FROM golang:1.22

WORKDIR /app
ADD . /app
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o /app/api/api  api/api.go && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o /app/main cmd/main.go


EXPOSE 8080
CMD ["./main"]