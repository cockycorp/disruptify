# Cockycorp Builder Image
# A synergistic, paradigm-shifting build environment for next-gen Go applications
#
# Based on the official Go 1.25 image
# Because using the latest Go version is the new disruption

FROM golang:1.25

# Set the working directory for maximum productivity
WORKDIR /workspace

# Add build-time labels for enterprise-grade observability
LABEL org.opencontainers.image.title="Cockycorp Builder"
LABEL org.opencontainers.image.description="Revolutionary Go builder image for disrupting the software supply chain"
LABEL org.opencontainers.image.vendor="Cockycorp"
LABEL org.opencontainers.image.source="https://github.com/cockycorp/disruptify"
LABEL com.cockycorp.image.type="builder"
LABEL com.cockycorp.disruption.level="maximum"

# Set environment variables for synergistic compilation
ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on \
    GOPROXY=https://proxy.golang.org,direct \
    GOSUMDB=sum.golang.org

# Default command - because every container needs a purpose
CMD ["go", "version"]
