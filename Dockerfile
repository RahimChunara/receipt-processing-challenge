FROM golang:1.22-alpine AS builder

# Set the working directory 
WORKDIR /app

# copy to workplace
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# copy src to container
COPY . .

# build application
RUN go build -o receipt-processor

# base image
FROM alpine:latest

WORKDIR /root/

# copy binary to builder
COPY --from=builder /app/receipt-processor .

# port number
EXPOSE 9000

# run code
CMD ["./receipt-processor"]
