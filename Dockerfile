FROM golang:1.22.2-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/app cmd/app/main.go cmd/app/wire_gen.go

FROM scratch
COPY --from=builder /app/bin/app /app
COPY --from=builder /app/cmd/app/.env .
ENTRYPOINT ["./app"]
