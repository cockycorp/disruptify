# GitHub Actions Workflows

## Overview

This project uses two workflows for maximum disruption:

1. **Builder Image Workflow** - Builds and publishes the base builder image
2. **Application Image Workflow** - Builds and publishes the application using the builder image

## Cockycorp Builder Image

The `builder-image.yml` workflow builds and publishes the `cockycorp-builder` base image.

### Triggers

- **Push to main**: When `images/Dockerfile.builder` or the workflow file changes
- **Builder tags**: Push tags like `builder-v1.0.0` for versioned releases
- **Pull requests**: Builds image only (doesn't push)
- **Manual**: Can be triggered via workflow_dispatch

### Builder Image Features

- Based on official `golang:1.25` image
- Pre-configured with optimal Go build settings
- Multi-platform support (AMD64 and ARM64)
- Published to `ghcr.io/cockycorp/disruptify-builder`

### Using the Builder Image

```dockerfile
FROM ghcr.io/cockycorp/disruptify-builder:latest

COPY . .
RUN go build -o myapp
```

### Build Order

**Important**: The builder image should be built and published first before building the application image.

1. First run: Builder image workflow
2. Then run: Application image workflow (uses the builder image)

## Docker Build and Push

The `docker-build-push.yml` workflow automatically builds and publishes Docker images to GitHub Container Registry (ghcr.io).

### Triggers

- **Push to main**: Builds and pushes image with tags: `latest`, `main`, and `main-<sha>`
- **Version tags**: Push tags like `v1.0.0` to create semver-tagged releases
- **Pull requests**: Builds image only (doesn't push) to verify the build works

### Image Tags

The workflow creates multiple tags for flexibility:

- `latest` - Always points to the latest main branch build
- `main` - Tracks the main branch
- `main-<sha>` - Specific commit SHA for reproducibility
- `v1.0.0`, `v1.0`, `v1` - Semantic version tags (when you push git tags)

### Multi-platform Support

Images are built for both `linux/amd64` and `linux/arm64` architectures, so they work on:
- x86_64 servers and workstations
- ARM-based systems (like Apple Silicon, AWS Graviton, Raspberry Pi)

### Using the Published Images

```bash
# Pull the latest version
docker pull ghcr.io/cockycorp/disruptify:latest

# Pull a specific version
docker pull ghcr.io/cockycorp/disruptify:v1.0.0

# Pull a specific commit
docker pull ghcr.io/cockycorp/disruptify:main-abc1234
```

### Build Cache

The workflow uses GitHub Actions cache to speed up builds:
- First build: ~2-3 minutes
- Subsequent builds: ~30-60 seconds (with cache)

### Permissions

The workflow requires:
- `contents: read` - To checkout the repository
- `packages: write` - To push images to ghcr.io

These are automatically provided by the `GITHUB_TOKEN`.

### Build Attestation

The workflow generates artifact attestations for supply chain security, providing cryptographic proof of:
- What repository built the image
- When it was built
- What commit SHA was used
- The build environment details

This helps verify the provenance and integrity of the published images.
