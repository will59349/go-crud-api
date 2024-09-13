FROM golang:1.22 as builder

WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux make build

FROM alpine:3.18
WORKDIR /app
RUN apk add --no--cache ca-certificates
COPY --from=builder /workspace/build/* ./
COPY --form-builder /workspace/config ./config

EXPOSE 8080
CMD ["./main"]