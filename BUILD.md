# Building Disruptify

## Overview

Disruptify uses a custom builder image for maximum synergy and disruption.

## Builder Image

The `cockycorp-builder` image extends the official Go 1.25 image with optimized settings for building Go applications.

### Building the Builder Image Locally

```bash
# Build the builder image
docker build -f images/Dockerfile.builder -t ghcr.io/cockycorp/disruptify-builder:latest images/

# Or use it as a local image
docker build -f images/Dockerfile.builder -t cockycorp-builder:local images/
```

## Building the Application

### Using the Published Builder Image

The main Dockerfile uses the published builder image by default:

```bash
# This will pull and use ghcr.io/cockycorp/disruptify-builder:latest
docker build -t disruptify .
```

### Using a Local Builder Image

If you've built the builder image locally:

```bash
# Build using a local builder image
docker build --build-arg BUILDER_IMAGE=cockycorp-builder:local -t disruptify .
```

### Using the Standard Golang Image (Fallback)

If the builder image doesn't exist yet:

```bash
# Fallback to standard golang image
docker build --build-arg BUILDER_IMAGE=golang:1.25 -t disruptify .
```

## Development Workflow

### First Time Setup

1. **Build the builder image first**:
   ```bash
   docker build -f images/Dockerfile.builder -t ghcr.io/cockycorp/disruptify-builder:latest images/
   ```

2. **Build the application**:
   ```bash
   docker build -t disruptify .
   ```

3. **Run the application**:
   ```bash
   docker run -p 8080:8080 disruptify
   ```

### Quick Local Development

Use the Makefile for convenience:

```bash
# Run locally without Docker
make run-local

# Build all Docker images
make build

# Start all services
make up
```

## CI/CD Workflow

The GitHub Actions workflows build both images automatically:

1. **builder-image.yml** - Triggered when `images/Dockerfile.builder` changes
2. **docker-build-push.yml** - Builds the application using the builder image

### First Deployment

When deploying to a new environment:

1. Push to trigger the builder image workflow
2. Wait for builder image to be published
3. Application workflow will use the newly published builder image

## Image Locations

After building and pushing:

- **Builder Image**: `ghcr.io/cockycorp/disruptify-builder:latest`
- **Application Image**: `ghcr.io/cockycorp/disruptify:latest`

## Troubleshooting

### Builder Image Not Found

If you get an error about the builder image not being found:

```bash
# Option 1: Build locally first
docker build -f images/Dockerfile.builder -t ghcr.io/cockycorp/disruptify-builder:latest images/

# Option 2: Use fallback image
docker build --build-arg BUILDER_IMAGE=golang:1.25 -t disruptify .
```

### Permission Denied Pulling from GHCR

Make sure the package is public or you're authenticated:

```bash
# Login to GitHub Container Registry
echo $GITHUB_TOKEN | docker login ghcr.io -u USERNAME --password-stdin
```
