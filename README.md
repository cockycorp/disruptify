# 🦾 Disruptifier

> **"Disrupting the industry since before it knew it needed disruption."** — *Cockycorp Vision Statement*

Disruptifier is the **open-source core** of Cockycorp's revolutionary, synergy-driven, paradigm-shifting **DisruptOps™** platform.

What does it do? It serves **AI-powered, blockchain-ready inspirational quotes** from the most visionary techbro thought leaders in the industry. Because every successful startup needs a daily dose of synergistic wisdom.

---

## 🚀 Key Features

* **20+ Hilarious Techbro Quotes** curated from the finest thought leaders
* **Beautiful Web UI** with gradient backgrounds for maximum disruption
* **RESTful API** at `/api/quote` for your microservices architecture
* **100% Go-based** for blazing-fast quote delivery
* **PostgreSQL Database** because data is the new oil
* **Grafana Monitoring** for synergistic observability
* **Docker & Docker Compose** ready for hyperscale deployment
* **Multi-platform Images** supporting AMD64 and ARM64 architectures
* **GitHub Actions CI/CD** with automatic image publishing to GHCR
* **Health Check Endpoint** at `/health` for Kubernetes readiness probes

---

## 🧩 Architecture Overview

Disruptifier embraces a *next-gen, cloud-first, containerized ecosystem of disruption*.

```
┌─────────────────┐
│   User/Browser  │
└────────┬────────┘
         │ HTTP
         ▼
┌─────────────────┐
│   Disruptifier  │ :8080
│   (Go Web App)  │
└────┬────────┬───┘
     │        │
     │        └─────────┐
     ▼                  ▼
┌──────────┐      ┌──────────┐
│PostgreSQL│:5432 │ Grafana  │:3000
│ Database │      │Monitoring│
└──────────┘      └──────────┘
```

Three containers working in perfect synergy to deliver maximum disruption.

---

## 🧠 Philosophy

At Cockycorp, we believe **disruption** isn’t a goal—it’s a lifestyle.
Disruptifier doesn’t *solve* problems. It *identifies opportunities to redefine the nature of problem-solving itself.*

---

## ⚙️ Quick Start

### Prerequisites

* Docker & Docker Compose
* Go ≥ 1.21 (optional, for local development)
* Make (optional, for convenience commands)

### Using Docker Compose (Recommended)

```bash
# Build and start all services
docker-compose up -d

# Or use the Makefile
make up
```

That's it! Services will be available at:
- **Web App**: http://localhost:8080
- **Grafana**: http://localhost:3000 (user: `disruptor`, pass: `synergy123`)
- **PostgreSQL**: localhost:5432 (user: `disruptor`, pass: `synergy123`)

### Running Locally (Without Docker)

```bash
go run main.go
# Or
make run-local
```

Visit http://localhost:8080 to experience maximum disruption.

### Using Pre-built Image from GitHub Container Registry

For maximum efficiency, pull the pre-built image:

```bash
# Pull the latest image
docker pull ghcr.io/cockycorp/disruptifier:latest

# Run it standalone
docker run -p 8080:8080 ghcr.io/cockycorp/disruptifier:latest

# Or update docker-compose.yml to use the pre-built image
# Replace the 'build' section with:
# image: ghcr.io/cockycorp/disruptifier:latest
```

The image is automatically built and pushed on every commit to main via GitHub Actions.

---

## 🛠️ Available Commands

```bash
make help          # Show all available commands
make build         # Build Docker images
make up            # Start all services
make down          # Stop all services
make logs          # View logs from all services
make restart       # Restart all services
make clean         # Remove all containers and volumes
make status        # Show service status
```

---

## 📡 API Endpoints

### `GET /`
Returns a beautiful HTML page with a random techbro quote.

### `GET /api/quote`
Returns a random quote in JSON format:
```json
{
  "quote": "We're not just disrupting the market...",
  "author": "Chad Thunderbro, CEO of SynergyChain"
}
```

### `GET /health`
Health check endpoint for monitoring:
```json
{
  "status": "disrupting successfully",
  "message": "All systems are synergistically aligned"
}
```

---

## 🏗️ Building from Source

Disruptifier uses a custom `nextgen-builder` image based on the official Go 1.25
image for secure, optimized builds.

### Quick Build

```bash
# Using the published builder image (recommended)
docker build -t disruptifier .

# Or fallback to standard golang image
docker build --build-arg BUILDER_IMAGE=golang:1.25 -t disruptifier .
```

For detailed build instructions, see [BUILD.md](BUILD.md).

---

## 🧑‍💻 Contributing

We welcome contributions from forward-thinkers, mavericks, and anyone who can correctly spell “Kubernetes” on the first try.

1. Fork the repo (disrupt responsibly).
2. Create a branch named like `feature/reinvent-the-wheel`.
3. Write code that sounds important.
4. Submit a pull request with a long description about “synergies” and “alignment.”

We’ll probably merge it without reading it. Because that’s disruption.

---

## 🪪 License

Licensed under the **Apache 2.0 License**, because even our lawyers disrupt convention.

---

## 🏁 Roadmap

* [ ] Integrate blockchain for no reason
* [ ] Add AI-powered synergy detection
* [ ] Monetize open source (???)
* [ ] Rebrand Disruptifier as Disruptifier Cloud Enterprise™

---

## 🤝 Join the Movement

Stop asking what Disruptifier does.
Start asking what *you* can disrupt with Disruptifier.
