FROM golang:1.23-alpine AS builder
WORKDIR /gau_assistant
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /gau_validation
COPY --from=builder /gau_assistant/main .
EXPOSE 8900
CMD ["./main"]