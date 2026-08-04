# 🚀 AI Orchestrator

### Production-Grade AI Orchestration Backend built with Go

*A modular, workflow-driven backend that transforms unstructured resume data into intelligent, structured insights using Large Language Models, Redis caching, PostgreSQL persistence, and production-ready REST APIs.*

---

![Go](https://img.shields.io/badge/Go-1.25-00ADD8?style=for-the-badge&logo=go)
![Gin](https://img.shields.io/badge/Gin-Web_Framework-00ADD8?style=for-the-badge)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-16-336791?style=for-the-badge&logo=postgresql)
![Redis](https://img.shields.io/badge/Redis-Cache-DC382D?style=for-the-badge&logo=redis)
![Docker](https://img.shields.io/badge/Docker-Containerized-2496ED?style=for-the-badge&logo=docker)
![Swagger](https://img.shields.io/badge/Swagger-API_Documentation-85EA2D?style=for-the-badge&logo=swagger)
![GitHub Actions](https://img.shields.io/badge/GitHub_Actions-CI-2088FF?style=for-the-badge&logo=githubactions)
![CodeQL](https://img.shields.io/badge/CodeQL-Security-2F80ED?style=for-the-badge)
![License](https://img.shields.io/badge/License-MIT-success?style=for-the-badge)

</div>

---

## 📑 Table of Contents

- [Overview](#-overview)
- [Why AI Orchestrator?](#-why-ai-orchestrator)
- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Repository Highlights](#-repository-highlights)
- [System Architecture](#️-system-architecture)
- [AI Workflow Pipeline](#-ai-workflow-pipeline)
- [Project Structure](#-project-structure)
- [Package Responsibilities](#-package-responsibilities)
- [Installation](#-installation)
  - [Running with Docker](#-running-with-docker)
  - [Running Locally](#-running-locally)
- [Configuration](#️-configuration)
- [API Reference](#-api-reference)
- [Core Subsystems](#-core-subsystems)
  - [Workflow Engine](#workflow-engine)
  - [Prompt Engineering](#prompt-engineering)
  - [Redis Caching](#redis-caching)
  - [PostgreSQL Persistence](#postgresql-persistence)
  - [Document Parsing](#document-parsing)
  - [Retry Mechanism](#retry-mechanism)
  - [Middleware](#middleware)
  - [Structured Logging](#structured-logging)
  - [Monitoring & Metrics](#monitoring--metrics)
- [Design Principles](#️-design-principles)
- [Testing & Quality Assurance](#-testing--quality-assurance)
- [Continuous Integration](#-continuous-integration)
- [Security](#-security)
- [Deployment](#-deployment)
- [Roadmap](#️-roadmap)
- [Future Improvements](#-future-improvements)
- [Contributing](#-contributing)
- [License](#-license)
- [Author](#-author)
- [Acknowledgements](#-acknowledgements)

---

## 📖 Overview

**AI Orchestrator** is a production-focused backend system that leverages Large Language Models (LLMs) to automate resume intelligence workflows through a modular orchestration engine.

Instead of exposing raw LLM calls directly, the project introduces reusable **workflow pipelines** responsible for prompt engineering, structured output generation, retry handling, caching, validation, persistence, and API orchestration.

The platform is designed around clean backend architecture principles, making it scalable, testable, and easy to extend with additional AI workflows.

Current supported capabilities:

- Resume Analysis
- ATS Score Generation
- Skill Extraction
- Job Matching
- Resume Improvement
- AI-powered Resume Chat
- Personalized Cover Letter Generation

Every workflow produces **structured JSON responses** suitable for downstream applications, rather than free-form text.

---

## 🎯 Why AI Orchestrator?

Modern LLM applications often become difficult to maintain because business logic, prompts, validation, caching, and API handling are tightly coupled.

AI Orchestrator addresses this by introducing dedicated workflow components, each responsible for a single AI task. Instead of treating the LLM *as* the application, the LLM becomes **one component inside a structured backend pipeline**.

This architecture provides:

- Modular workflow execution
- Reusable AI pipelines
- Consistent JSON outputs
- Retry support
- Redis-backed caching
- Database persistence
- Clean separation of concerns
- Production-ready REST APIs

---

## ✨ Features

### 🤖 AI Workflows

- Resume Analysis
- ATS Score Calculation
- Skill Extraction
- Resume Improvement
- Resume Chat
- Job Matching
- Cover Letter Generation

### ⚡ Backend

- Modular Workflow Engine
- Generic Pipeline Executor
- Prompt-based AI Orchestration
- Automatic Retry Mechanism
- Redis Caching
- PostgreSQL Persistence
- Structured JSON Responses
- REST APIs using Gin
- Swagger Documentation

### 📂 Resume Processing

- PDF Resume Parsing
- DOCX Resume Parsing
- Session-based Resume Storage
- Conversation History
- Resume Upload APIs

### 🚀 Production

- Docker Support
- Docker Compose
- GitHub Actions CI
- CodeQL Security Scanning
- Dependabot
- Unit Testing
- Prometheus Metrics
- Structured Logging (Zap)
- Request ID Middleware
- Rate Limiting Middleware

### 🧪 Quality Assurance

- High Unit Test Coverage
- Automated CI Pipeline
- Static Analysis using `go vet`
- Formatting Validation using `gofmt`
- Docker Build Validation

---

## 🛠 Tech Stack

| Category | Technologies |
|-----------|--------------|
| Language | Go 1.25 |
| Web Framework | Gin |
| AI Model | Google Gemini |
| Database | PostgreSQL |
| Cache | Redis |
| Documentation | Swagger |
| Containerization | Docker |
| CI/CD | GitHub Actions |
| Security | CodeQL |
| Logging | Zap |
| Metrics | Prometheus |
| Testing | Go Testing Package |

---

## 🌟 Repository Highlights

- 🏗 Clean Layered Architecture
- 🔄 Generic AI Pipeline Executor
- 🧠 Workflow-based LLM Orchestration
- ⚡ Redis-backed Intelligent Caching
- 🐳 Fully Dockerized
- 📖 Swagger API Documentation
- 📊 Prometheus Metrics
- 🔐 Request Tracking & Rate Limiting
- 🧪 Comprehensive Unit Testing
- 🚀 Production-ready CI/CD Pipeline

---

## 🏗️ System Architecture

The project follows a layered backend architecture with clear separation of concerns. Each layer has a single responsibility, making the application modular, maintainable, and easy to extend.

```text
                        ┌──────────────────────────────┐
                        │          Client              │
                        │ (Web / Mobile / Postman)     │
                        └──────────────┬───────────────┘
                                       │
                                       ▼
                        ┌──────────────────────────────┐
                        │        Gin HTTP Server       │
                        │  Routes + Middleware Layer   │
                        └──────────────┬───────────────┘
                                       │
                                       ▼
                        ┌──────────────────────────────┐
                        │        API Handlers          │
                        │ Request Validation & Binding │
                        └──────────────┬───────────────┘
                                       │
                                       ▼
                        ┌──────────────────────────────┐
                        │       Orchestrator Layer     │
                        │ Generic Pipeline Execution   │
                        └──────────────┬───────────────┘
                                       │
                  ┌────────────────────┴────────────────────┐
                  ▼                                         ▼
        ┌───────────────────┐                  ┌────────────────────┐
        │   Workflow Layer  │                  │     Redis Cache    │
        │ Prompt Generation │                  │ Session & Response │
        └─────────┬─────────┘                  └────────────────────┘
                  │
                  ▼
        ┌────────────────────┐
        │    Gemini LLM      │
        │ AI Content Engine  │
        └─────────┬──────────┘
                  │
                  ▼
        ┌────────────────────┐
        │ JSON Parsing Layer │
        │ Structured Outputs │
        └─────────┬──────────┘
                  │
                  ▼
        ┌────────────────────┐
        │ PostgreSQL Storage │
        │ Analysis History   │
        └────────────────────┘
```

The architecture isolates business logic from transport, caching, persistence, and AI providers, making every workflow independently testable and reusable.

---

## 🧠 AI Workflow Pipeline

Every AI endpoint follows the same execution pipeline, shared through a generic pipeline executor to minimize duplicated business logic.

```text
Incoming Request
        │
        ▼
Validate Request
        │
Generate Cache Key
        │
        ▼
Redis Lookup
        │
 ┌──────┴────────┐
 │               │
 │ Cache Hit     │
 │               │
 ▼               ▼
Return      Build Prompt
Response          │
                  ▼
          Execute Workflow
                  │
                  ▼
             Gemini LLM
                  │
                  ▼
          JSON Validation
                  │
                  ▼
          Retry (if needed)
                  │
                  ▼
           Store in Cache
                  │
                  ▼
          Return Response
```

### Full Request Lifecycle

```text
HTTP Request
      │
      ▼
Gin Router
      │
      ▼
Middleware
(Request ID + Logging + Rate Limiting)
      │
      ▼
Request Validation
      │
      ▼
Handler
      │
      ▼
Resume Parser
      │
      ▼
Orchestrator
      │
      ▼
Workflow Engine
      │
      ▼
Gemini AI
      │
      ▼
JSON Parsing
      │
      ▼
Redis Cache
      │
      ▼
PostgreSQL
      │
      ▼
HTTP JSON Response
```

---

## 📂 Project Structure

```text
AI-Orchestrator
│
├── cmd/
│   └── server/
│       ├── main.go
│       ├── router.go
│       └── bootstrap.go
│
├── docs/
│   └── Swagger Documentation
│
├── internal/
│   │
│   ├── api/
│   │   ├── handler/
│   │   ├── middleware/
│   │   └── routes/
│   │
│   ├── orchestrator/
│   ├── workflow/
│   ├── parser/
│   ├── cache/
│   ├── repository/
│   ├── database/
│   ├── llm/
│   ├── retry/
│   ├── metrics/
│   ├── logger/
│   └── model/
│
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

---

## 📦 Package Responsibilities

| Package | Responsibility |
|----------|----------------|
| `handler` | HTTP request handling and validation |
| `routes` | API route registration |
| `middleware` | Logging, Request ID, Rate Limiting |
| `orchestrator` | Generic workflow execution |
| `workflow` | Prompt generation and AI workflow logic |
| `llm` | Gemini client wrapper |
| `cache` | Redis caching and session management |
| `repository` | PostgreSQL persistence |
| `database` | Database initialization |
| `parser` | PDF & DOCX resume parsing |
| `retry` | Automatic retry mechanism |
| `metrics` | Prometheus metrics |
| `logger` | Structured Zap logging |
| `model` | Request/Response DTOs |

---

## 📥 Installation

AI Orchestrator can be run either locally using Go or inside Docker containers.

### Prerequisites

| Tool | Version |
|------|----------|
| Go | 1.25+ |
| Docker | Latest |
| Docker Compose | Latest |
| PostgreSQL | 16+ |
| Redis | 7+ |
| Git | Latest |

### Clone the Repository

```bash
git clone https://github.com/Aditya7880900936/ai-orchestrator.git
cd ai-orchestrator
```

### Install Dependencies

```bash
go mod download
go mod tidy
```

### Environment Configuration

Create a `.env` file in the project root:

```env
GEMINI_API_KEY=YOUR_GEMINI_API_KEY
DATABASE_URL=postgres://postgres:postgres@localhost:5433/ai_orchestrator?sslmode=disable
REDIS_ADDR=localhost:6379
```

---

### 🐳 Running with Docker

The fastest way to start the entire stack:

```bash
docker compose up --build
```

This automatically starts:

- AI Orchestrator
- PostgreSQL
- Redis

| Resource | URL |
|-----------|-----|
| Swagger UI | `http://localhost:8080/swagger/index.html` |
| Health Check | `http://localhost:8080/health` |
| Prometheus Metrics | `http://localhost:8080/metrics` |

To stop all services:

```bash
docker compose down
```

---

### 💻 Running Locally

Start Redis:

```bash
docker compose up redis -d
```

Start PostgreSQL:

```bash
docker compose up postgres -d
```

Run the server:

```bash
go run ./cmd/server
```

The server starts on `http://localhost:8080`.

---

## ⚙️ Configuration

The project uses environment variables for runtime configuration.

| Variable | Description |
|-----------|-------------|
| `GEMINI_API_KEY` | Google Gemini API Key |
| `DATABASE_URL` | PostgreSQL connection string |
| `REDIS_ADDR` | Redis address |

Example:

```env
GEMINI_API_KEY=xxxxxxxxxxxxxxxx
DATABASE_URL=postgres://postgres:postgres@localhost:5433/ai_orchestrator?sslmode=disable
REDIS_ADDR=localhost:6379
```

---

## 📡 API Reference

All APIs accept and return **JSON** unless explicitly mentioned.

**Base URL:** `http://localhost:8080`

**Interactive Documentation:** `http://localhost:8080/swagger/index.html`

Swagger documentation is automatically generated and includes request schemas, response schemas, error responses, and interactive testing for every endpoint.

### Endpoint Summary

| Method | Endpoint | Description |
|----------|----------|-------------|
| `POST` | `/analyze` | Resume Analysis |
| `POST` | `/resume/upload` | Upload Resume |
| `POST` | `/resume/analyze` | Resume Analysis from Session |
| `POST` | `/skills/extract` | Extract Skills |
| `POST` | `/ats/score` | ATS Score |
| `POST` | `/job/match` | Match Resume with Job Description |
| `POST` | `/resume/improve` | Improve Resume |
| `POST` | `/cover-letter/generate` | Generate Cover Letter |
| `POST` | `/resume/chat` | AI Resume Chat |
| `GET` | `/health` | Health Check |
| `GET` | `/metrics` | Prometheus Metrics |

### Standard Response Format

Successful response:

```json
{
  "success": true,
  "data": {}
}
```

Error response:

```json
{
  "success": false,
  "error": "error message"
}
```

---

### ❤️ Health Check

```http
GET /health
```

**Response**

```json
{
  "status": "ok"
}
```

---

### 📄 Upload Resume

Uploads a PDF or DOCX resume and stores it for later AI workflows.

```http
POST /resume/upload
Content-Type: multipart/form-data
```

**Parameters**

| Name | Type | Required |
|------|------|----------|
| `file` | PDF / DOCX | ✅ |

**Response**

```json
{
  "message": "Resume uploaded successfully"
}
```

**Possible Errors**

- Unsupported file type
- Empty file
- Parsing failure

---

### 🧠 Resume Analysis

Extracts structured insights from a resume.

```http
POST /resume/analyze
```

**Request**

```json
{
  "resume": "Your resume text..."
}
```

**Response**

```json
{
  "summary": "...",
  "skills": [
    "Go",
    "Redis",
    "Docker"
  ],
  "experience_years": 2,
  "strengths": [],
  "missing_skills": []
}
```

---

### 🎯 ATS Score

Evaluates resume quality from an Applicant Tracking System perspective.

```http
POST /ats/score
```

**Request**

```json
{
  "resume": "..."
}
```

**Response**

```json
{
  "overall_score": 89,
  "section_scores": {
    "skills": 95,
    "experience": 90,
    "education": 85
  },
  "strengths": [],
  "weaknesses": [],
  "missing_keywords": [],
  "suggestions": []
}
```

---

### 💼 Job Matching

Matches a resume against a job description.

```http
POST /job/match
```

**Request**

```json
{
  "resume": "...",
  "job_description": "..."
}
```

**Response**

```json
{
  "match_percentage": 86,
  "matched_skills": [],
  "missing_skills": [],
  "recommendations": []
}
```

---

### 🚀 Resume Improvement

Improves resume language while preserving factual correctness.

```http
POST /resume/improve
```

**Request**

```json
{
  "resume": "..."
}
```

**Response**

```json
{
  "improved_summary": "...",
  "improved_experience": [],
  "improved_projects": [],
  "missing_sections": [],
  "overall_suggestions": []
}
```

---

### 🛠 Skill Extraction

Extracts structured skills from a resume.

```http
POST /skills/extract
```

**Request**

```json
{
  "resume": "..."
}
```

**Response**

```json
{
  "technical_skills": [],
  "frameworks": [],
  "databases": [],
  "cloud": [],
  "tools": [],
  "soft_skills": []
}
```

---

### 💬 Resume Chat

Allows conversational interaction with a resume.

```http
POST /resume/chat
```

**Request**

```json
{
  "resume": "...",
  "conversation": [],
  "question": "What backend technologies does the candidate know?"
}
```

**Response**

```json
{
  "answer": "..."
}
```

---

### ✉️ Cover Letter Generation

Generates a personalized, ATS-friendly cover letter.

```http
POST /cover-letter/generate
```

**Request**

```json
{
  "candidate_information": "..."
}
```

**Response**

```json
{
  "cover_letter": "..."
}
```

---

### 📊 Prometheus Metrics

Exposes Prometheus-compatible application metrics.

```http
GET /metrics
```

Used for monitoring, dashboarding, performance analysis, and alerting.

---

## 🔧 Core Subsystems

### Workflow Engine

The Workflow Engine is the core abstraction of AI Orchestrator. Rather than tightly coupling prompt engineering with API handlers, every AI capability is implemented as an independent workflow.

Each workflow is responsible only for:

- Constructing the prompt
- Defining the expected JSON schema
- Invoking the LLM client

The orchestration layer handles:

- Request execution
- Retry logic
- Response parsing
- Error handling
- Caching
- Persistence

This design enables new AI capabilities to be added with minimal changes to the existing codebase.

| Workflow | Purpose |
|----------|---------|
| Resume Analysis | Generate structured resume insights |
| ATS Scoring | Evaluate ATS compatibility |
| Skill Extraction | Extract technical & soft skills |
| Job Matching | Compare resume against job description |
| Resume Improvement | Improve resume while preserving factual correctness |
| Resume Chat | Conversational querying over resume content |
| Cover Letter | Generate personalized cover letters |

**Workflow execution:**

```text
Incoming Request
        │
        ▼
Select Workflow
        │
        ▼
Generate Prompt
        │
        ▼
Call Gemini
        │
        ▼
Receive JSON
        │
        ▼
Validate Response
        │
        ▼
Return Structured Output
```

---

### Prompt Engineering

Every AI workflow follows a structured prompt engineering strategy designed to:

- Produce valid JSON only
- Prevent hallucinations
- Preserve factual correctness
- Maintain deterministic output formats
- Reduce downstream parsing complexity

Example constraints applied to prompts:

- Return **only** valid JSON
- Do **not** use Markdown
- Do **not** invent experience
- Preserve factual accuracy
- Follow the predefined JSON schema

This approach significantly improves response consistency across workflows.

---

### Redis Caching

AI inference is computationally expensive. To reduce latency and avoid repeated LLM calls, AI Orchestrator uses Redis as a caching layer. Cached responses are reused whenever possible, improving both performance and cost efficiency.

```text
Request
   │
   ▼
Generate Cache Key
   │
   ▼
Redis Lookup
   │
┌──┴────────────┐
│               │
│ Cache Hit     │
│               │
▼               ▼
Return      Execute Workflow
Result            │
                  ▼
             Gemini API
                  │
                  ▼
            Store Response
                  │
                  ▼
             Return Result
```

**Benefits:** lower latency, reduced API cost, faster repeated requests, improved user experience.

---

### PostgreSQL Persistence

The application stores analysis results in PostgreSQL, enabling analysis history, future analytics, auditability, and workflow tracking.

The repository layer abstracts database operations from business logic:

```text
Handler
   │
   ▼
Repository Interface
   │
   ▼
PostgreSQL
```

Using the Repository Pattern simplifies testing and allows database implementations to evolve independently from application logic.

---

### Document Parsing

AI Orchestrator supports parsing resumes from multiple document formats — currently **PDF** and **DOCX**. The parser layer automatically detects the document type and delegates processing to the appropriate implementation.

```text
Resume Upload
      │
      ▼
Detect Extension
      │
 ┌────┴────┐
 ▼         ▼
PDF      DOCX
Parser    Parser
 └────┬────┘
      ▼
Extract Text
      ▼
Workflow Engine
```

This abstraction allows additional parsers (TXT, ODT, HTML, etc.) to be introduced without changing business logic.

---

### Retry Mechanism

External AI services may occasionally fail due to network instability, rate limiting, or temporary provider errors. AI Orchestrator implements a retry mechanism before returning an error to the client.

```text
Execute Request
      │
      ▼
Success?
      │
 ┌────┴────┐
 │         │
Yes       No
 │         │
 ▼         ▼
Return   Retry
            │
            ▼
       Max Retries?
            │
     ┌──────┴──────┐
     │             │
    Yes            No
     │             │
     ▼             ▼
 Return Error   Retry Again
```

This improves resilience against transient failures while avoiding unnecessary request failures.

---

### Middleware

Every request passes through multiple middleware layers before reaching business logic.

```text
Incoming Request
        │
        ▼
Request ID
        │
        ▼
Logging
        │
        ▼
Rate Limiter
        │
        ▼
Request Validation
        │
        ▼
API Handler
```

- **Request ID** — assigns a unique identifier to every incoming request for traceability.
- **Logging** — captures structured request logs using Zap.
- **Rate Limiting** — protects APIs from abuse using IP-based rate limiting.

This ensures unique request IDs, structured logging, API rate limiting, request validation, and centralized error handling.

---

### Structured Logging

The application uses Uber's **Zap** logger for high-performance structured logging. Each request captures HTTP method, request path, client IP, response status, and latency.

```json
{
  "method": "POST",
  "path": "/resume/analyze",
  "status": 200,
  "latency": "42ms"
}
```

Structured logs simplify debugging, monitoring, and production observability.

---

### Monitoring & Metrics

Prometheus metrics are exposed at `GET /metrics` for monitoring application health and usage. Current metrics include analyze request count, HTTP metrics (extensible), and custom application metrics. These can be integrated with Prometheus and Grafana for production monitoring.

---

## 🏛️ Design Principles

AI Orchestrator follows several software engineering principles:

- Separation of Concerns
- Layered Architecture
- Dependency Injection
- Repository Pattern
- Modular Workflow Design
- Reusable Components
- Testability
- Production-first Development

These principles improve maintainability, scalability, and long-term extensibility.

---

## 🧪 Testing & Quality Assurance

AI Orchestrator is designed with testability as a first-class concern. The project includes comprehensive unit tests across core backend components to ensure correctness, maintainability, and confidence during future development.

### Test Coverage

| Package | Purpose |
|----------|---------|
| Handler | API request handling |
| Middleware | Request ID, Logging, Rate Limiting |
| Routes | Route registration |
| Orchestrator | Workflow orchestration |
| Workflow | AI workflow execution |
| Cache | Redis cache logic |
| Database | Database initialization |
| Repository | PostgreSQL operations |
| Retry | Retry mechanism |
| Parser | Resume parsing |
| Logger | Zap logger initialization |
| Metrics | Prometheus metrics |
| LLM | Gemini client initialization |
| Server | Bootstrap & Router |

### Running Tests

```bash
# Run all tests
go test ./...

# Verbose output
go test -v ./...

# With coverage
go test ./... -cover

# Coverage for a specific package
go test ./internal/orchestrator -cover
```

### Static Analysis & Build Verification

```bash
go vet ./...
go fmt ./...
go build ./...
```

Every Pull Request is automatically validated through GitHub Actions before merging.

---

## 🚀 Continuous Integration

The project uses GitHub Actions for automated CI. Every push and pull request executes the validation pipeline.

```text
Checkout Repository
        │
        ▼
Setup Go
        │
        ▼
Download Dependencies
        │
        ▼
Verify Formatting
        │
        ▼
Go Vet
        │
        ▼
Build Project
        │
        ▼
Run Unit Tests
        │
        ▼
Generate Coverage
        │
        ▼
Validate Docker Compose
        │
        ▼
Build Docker Image
```

The CI workflow validates code formatting, static analysis, successful build, unit tests, Docker image build, and Docker Compose configuration — preventing broken code from being merged into `main`.

---

## 🔒 Security

**CodeQL** — GitHub CodeQL performs automated static security analysis to detect potential vulnerabilities and insecure coding patterns.

**Dependabot** — continuously monitors project dependencies and automatically opens Pull Requests whenever security patches or version updates become available.

**Secure Configuration** — sensitive values such as API keys and database credentials are managed through environment variables instead of being hardcoded:

```env
GEMINI_API_KEY=YOUR_API_KEY
DATABASE_URL=...
REDIS_ADDR=...
```

---

## 🐳 Deployment

AI Orchestrator is fully containerized using Docker. The project includes a `Dockerfile`, `docker-compose.yml`, a PostgreSQL container, and a Redis container.

```bash
# Build the image
docker build -t ai-orchestrator .

# Run the full stack
docker compose up --build

# Stop containers
docker compose down
```

This provides a reproducible local development environment with minimal setup.

---

## 🗺️ Roadmap

- [ ] JWT Authentication
- [ ] OAuth Integration
- [ ] Multi-LLM Provider Support (OpenAI, Claude, Gemini)
- [ ] Streaming AI Responses
- [ ] Background Job Processing
- [ ] Vector Database Integration
- [ ] Semantic Resume Search
- [ ] Kubernetes Deployment
- [ ] Distributed Caching
- [ ] Grafana Dashboards
- [ ] OpenTelemetry Tracing
- [ ] Role-Based Access Control
- [ ] API Versioning
- [ ] WebSocket Support
- [ ] Multi-Tenant Architecture

---

## 📈 Future Improvements

- Advanced prompt versioning
- Workflow composition engine
- Agentic AI orchestration
- Resume embedding generation
- AI-powered interview preparation
- Resume similarity search
- Event-driven architecture
- Queue-based asynchronous execution
- AI evaluation framework

---

## 🤝 Contributing

Contributions are welcome!

1. Fork the repository.
2. Create a feature branch.
3. Commit your changes.
4. Open a Pull Request.

Please ensure that code is formatted, tests pass successfully, and documentation is updated where necessary.

See `CONTRIBUTING.md` for detailed instructions.

---

## 📜 License

This project is licensed under the MIT License. See the `LICENSE` file for more details.

---

## 👨‍💻 Author

**Aditya Sanskar Srivastav**

Backend Developer | Golang | AI Systems | Open Source Contributor

- GitHub: [@Aditya7880900936](https://github.com/Aditya7880900936)
- LinkedIn: *(add your LinkedIn profile here)*

If you found this project useful, consider giving it a ⭐ on GitHub.

---

## 🙏 Acknowledgements

This project makes use of several excellent open-source technologies:

Go · Gin · PostgreSQL · Redis · Google Gemini · Docker · Prometheus · Swagger · Zap Logger · GitHub Actions

Special thanks to the maintainers and contributors of these projects for building the tools that make modern backend development possible.

---

<div align="center">

**Built with Go 🐹 — designed for production.**

</div>