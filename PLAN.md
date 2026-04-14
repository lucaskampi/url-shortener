# URL Shortener - Project Plan

A self-hosted URL shortener built on AWS, starting simple and evolving to serverless.

---

## Architecture Overview

```
User clicks short URL
  ↓
CloudFront (global CDN)
  ↓
API Gateway → Lambda → DynamoDB (redirect lookup)
  ↓
Redirect to long URL
```

### AWS Services Used

| Service | Purpose |
|---------|---------|
| Lambda | URL redirect logic |
| API Gateway | REST API for creating/managing short URLs |
| DynamoDB | Store URL mappings (fast key-value lookups) |
| S3 | Store click analytics |
| CloudFront | Global CDN for redirects |
| Route 53 | Custom domain |
| ACM | SSL certificates |

---

## Project Phases

### Phase 1: Local Development (MVP)

- [ ] **Setup Go project structure**
  - Initialize Go module
  - Create clean architecture folder structure
  - Add main.go entry point
  - Configure Makefile for build/run

- [ ] **Implement URL redirect handler**
  - In-memory URL store (for initial development)
  - GET `/:shortCode` endpoint that redirects
  - Health check endpoint

- [ ] **Implement URL creation endpoint**
  - POST `/urls` to create short URL
  - Request: `{ "longUrl": "https://..." }`
  - Response: `{ "shortCode": "abc123", "shortUrl": "https://short.url/abc123" }`
  - Generate 6-character alphanumeric codes

- [ ] **Add basic validation**
  - Validate URL format
  - Handle malformed requests
  - Return appropriate HTTP status codes

- [ ] **Write unit tests**
  - Test redirect logic
  - Test URL creation
  - Test validation

---

### Phase 2: Data Persistence

- [ ] **Setup DynamoDB**
  - Design table schema:
    - Partition key: `shortCode` (String)
    - Attributes: `longUrl`, `createdAt`, `clickCount`
  - Create local DynamoDB for development
  - Configure AWS SDK

- [ ] **Replace in-memory store with DynamoDB**
  - Implement URL repository interface
  - Add DynamoDB repository implementation
  - Update dependency injection

---

### Phase 3: Analytics

- [ ] **Track clicks**
  - Increment clickCount on each redirect
  - Add `lastClickedAt` timestamp

- [ ] **Store analytics in S3**
  - Log click events to S3 (JSON Lines format)
  - Include: timestamp, shortCode, referrer, userAgent

- [ ] **Add analytics endpoint**
  - GET `/urls/:shortCode/stats`
  - Response: click count, created date, last clicked

---

### Phase 4: Production-Ready

- [ ] **Deploy to AWS**
  - Package as Docker image
  - Deploy to ECS or Lambda
  - Setup API Gateway
  - Configure DynamoDB table

- [ ] **Add custom domain**
  - Setup Route 53
  - Configure ACM for SSL
  - Setup CloudFront

- [ ] **Rate limiting & security**
  - Add API key authentication
  - Implement rate limiting
  - Add input sanitization

---

### Phase 5: Serverless Migration (Future)

- [ ] **Convert to Lambda**
  - Package redirect handler as Lambda
  - Use API Gateway proxy integration
  - Benchmark performance

- [ ] **Optimize for Lambda**
  - Add DynamoDB DAX if needed
  - Implement connection pooling
  - Configure reserved concurrency

---

## File Structure (Phase 1)

```
url-shortener/
├── cmd/
│   └── server/
│       └── main.go           # Entry point
├── internal/
│   ├── handler/
│   │   ├── redirect.go      # Redirect handler
│   │   └── create.go        # Create URL handler
│   ├── service/
│   │   └── url.service.go    # Business logic
│   ├── repository/
│   │   └── url.repository.go # Data access interface
│   └── model/
│       └── url.go            # Domain model
├── pkg/
│   └── generator/
│       └── shortcode.go      # Short code generator
├── Makefile
├── go.mod
└── go.sum
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/:shortCode` | Redirect to long URL |
| POST | `/urls` | Create new short URL |
| GET | `/urls/:shortCode/stats` | Get URL statistics |
| GET | `/health` | Health check |

## Example Requests

```bash
# Create short URL
curl -X POST http://localhost:8080/urls \
  -H "Content-Type: application/json" \
  -d '{"longUrl": "https://example.com/very/long/url"}'

# Response: { "shortCode": "a1b2c3", "shortUrl": "http://localhost:8080/a1b2c3" }

# Redirect
curl -I http://localhost:8080/a1b2c3
# Returns 302 redirect to long URL

# Stats
curl http://localhost:8080/urls/a1b2c3/stats
# Response: { "shortCode": "a1b2c3", "clickCount": 42, "createdAt": "..." }
```

## Running Locally

```bash
# Start server
make run

# Run tests
make test

# Build
make build
```

## Next Steps

1. Initialize Go project with `go mod init`
2. Create folder structure
3. Implement URL model and short code generator
4. Build in-memory repository for rapid iteration
5. Implement create and redirect handlers
6. Add tests
7. Verify locally with `make run`