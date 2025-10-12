#get the base go image

FROM golang:1.23-alpine

# Set working directory
WORKDIR /app


# Copy go.mod and go.sum first (for caching dependencies)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .


# Build the app
RUN go build -o server main.go

# Expose port
EXPOSE 8080

# Run the app
CMD ["./server"]