FROM golang:1.19-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o product-service .
CMD ["./product-service"]
