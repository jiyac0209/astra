# Use a Golang base image
FROM golang:1.22 AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire source code to the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Use a lightweight base image
FROM alpine:latest

# Set environment variables
ENV PORT=8080

# Expose port 8080 to the outside world
EXPOSE $PORT

# Copy the executable from the build stage to the /app directory of the final image
COPY --from=build /app/app /app/app

# Command to run the executable
CMD ["/app/app"]
