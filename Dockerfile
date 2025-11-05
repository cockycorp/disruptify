# Build stage
#
# Using Cockycorp's revolutionary builder image for maximum synergy
# Fallback to standard golang image if builder not available
ARG BUILDER_IMAGE=ghcr.io/cockycorp/disruptify-builder:latest
FROM ${BUILDER_IMAGE} AS builder

WORKDIR /workspace

# Copy go mod files
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy source code
COPY main.go .

# Build the application with our pre-configured environment
RUN go build -a -installsuffix cgo -o disruptify .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /workspace/disruptify .

# Expose port 8080
EXPOSE 8080

# Set default PORT environment variable
ENV PORT=8080

# Run the application
CMD ["./disruptify"]
