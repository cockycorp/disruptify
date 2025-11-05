.PHONY: help build up down logs clean restart

help: ## Show this help message
	@echo "🚀 Disruptify - Disrupting the Quote Industry"
	@echo ""
	@echo "Available commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'

build: ## Build all Docker images
	@echo "🔨 Building disruptive Docker images..."
	docker-compose build

up: ## Start all services
	@echo "🚀 Launching Disruptify ecosystem..."
	docker-compose up -d
	@echo ""
	@echo "✨ Services are now disrupting:"
	@echo "   🌐 Web App:    http://localhost:8080"
	@echo "   📊 Grafana:    http://localhost:3000 (user: disruptor, pass: synergy123)"
	@echo "   🗄️  PostgreSQL: localhost:5432 (user: disruptor, pass: synergy123)"

down: ## Stop all services
	@echo "⏹️  Stopping all services..."
	docker-compose down

logs: ## View logs from all services
	docker-compose logs -f

logs-app: ## View application logs only
	docker-compose logs -f app

logs-db: ## View database logs only
	docker-compose logs -f postgres

logs-grafana: ## View Grafana logs only
	docker-compose logs -f grafana

restart: ## Restart all services
	@echo "🔄 Restarting all services..."
	docker-compose restart

clean: ## Remove all containers, volumes, and images
	@echo "🧹 Cleaning up..."
	docker-compose down -v
	docker system prune -f

status: ## Show status of all services
	docker-compose ps

run-local: ## Run the app locally without Docker
	@echo "🏃 Running locally on port 8080..."
	go run main.go
