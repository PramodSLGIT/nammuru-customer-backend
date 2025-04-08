FROM golang:1.22.2

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Expose backend port
EXPOSE 8010

# Run the app from main folder
CMD ["go", "run", "./main/main.go"]
