# Use the official Golang image with version 1.21 as the base image for the build stage
FROM golang:1.21 as build-stage

# Set the working directory within the container
WORKDIR /assignment-4

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./
# Copy the entire local directory into the working directory of the container
COPY . .

# Update the package lists for the APT package manager
RUN apt-get update
# Install the PostgreSQL client in the container
RUN apt-get -y install postgresql-client
RUN chmod +x postgres_check.sh

# Download the Go module dependencies
RUN go mod download

# Build the Go application
RUN go build -o app ./cmd/api/

EXPOSE 8080
EXPOSE 5432

# Command to run the executable
CMD ["./app"]