FROM golang:1.17-alpine as builder

WORKDIR /app
COPY go.mod /app
COPY go.sum /app
RUN go mod download
COPY . .
RUN go build -o /app app/main.go

FROM alpine
WORKDIR /app
COPY --from=builder /app/main /app/
CMD /app/main