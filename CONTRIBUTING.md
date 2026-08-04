# 🤝 Contributing to AI Orchestrator

First of all, thank you for your interest in contributing to AI Orchestrator! 🎉

We welcome bug fixes, new features, documentation improvements, performance optimizations, and testing enhancements.

---

# 📋 Prerequisites

Before contributing, ensure you have the following installed:

- Go 1.25+
- Docker & Docker Compose
- PostgreSQL
- Redis
- Git

---

# 🚀 Getting Started

## 1. Fork the Repository

Click the **Fork** button on GitHub.

---

## 2. Clone your fork

```bash
git clone https://github.com/<your-username>/ai-orchestrator.git
```

```bash
cd ai-orchestrator
```

---

## 3. Install dependencies

```bash
go mod download
```

---

## 4. Configure Environment

Create a `.env` file.

Example:

```env
DATABASE_URL=postgres://postgres:postgres@localhost:5433/ai_orchestrator?sslmode=disable

REDIS_ADDR=localhost:6379

GEMINI_API_KEY=YOUR_API_KEY
```

---

## 5. Start Dependencies

```bash
docker compose up -d
```

---

## 6. Run the application

```bash
go run ./cmd/server
```

Swagger:

```
http://localhost:8080/swagger/index.html
```

---

# 🧪 Running Tests

Run every unit test before opening a Pull Request.

```bash
go test ./...
```

Run with coverage

```bash
go test ./... -cover
```

---

# 🎨 Formatting

Always format code before committing.

```bash
go fmt ./...
```

---

# 🔍 Static Analysis

Run Go Vet

```bash
go vet ./...
```

---

# 🐳 Docker Validation

Ensure Docker build succeeds.

```bash
docker compose up --build
```

---

# ✅ Pull Request Checklist

Before submitting a PR, verify:

- [ ] Project builds successfully
- [ ] All tests pass
- [ ] New tests added (if applicable)
- [ ] Code formatted with gofmt
- [ ] go vet passes
- [ ] Documentation updated
- [ ] No unnecessary dependencies added

---

# 💬 Commit Message Guidelines

Use Conventional Commits whenever possible.

Examples:

```text
feat: add resume improvement workflow

fix: resolve redis session bug

test: improve orchestrator coverage

docs: update README

refactor: simplify workflow execution
```

---

# 🏗 Project Structure

```
cmd/
internal/
docs/
.github/
Dockerfile
docker-compose.yml
```

---

# 💡 Contribution Ideas

You can contribute by:

- Improving AI prompts
- Optimizing Redis caching
- Adding new workflows
- Improving parser support
- Enhancing API documentation
- Increasing test coverage
- Improving Docker deployment
- Performance optimizations

---

# ❤️ Thank You

Every contribution—whether it's code, documentation, tests, or bug reports—is greatly appreciated.

Happy Coding 🚀