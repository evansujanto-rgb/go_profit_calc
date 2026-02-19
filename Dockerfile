# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY . .

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o profit_calc

# Final stage
FROM scratch

COPY --from=builder /app/profit_calc /profit_calc

CMD ["/profit_calc"] 