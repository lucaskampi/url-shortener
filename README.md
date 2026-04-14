# URL Shortener

A self-hosted URL shortener built with Go backend and Next.js frontend.

## Architecture

```
Frontend (Next.js) → Backend (Go) → In-Memory Store
```

## Tech Stack

- **Backend**: Go, net/http
- **Frontend**: Next.js 14, React, TypeScript, Tailwind CSS
- **Ports**: Backend on :8080, Frontend on :3000

## Quick Start

### Backend

```bash
go run ./cmd/server
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| POST | `/urls` | Create short URL |
| GET | `/:shortCode` | Redirect to long URL |
| GET | `/urls/:shortCode/stats` | Get URL statistics |
| GET | `/health` | Health check |

## Project Structure

```
url-shortener/
├── cmd/server/          # Go entry point
├── internal/            # Business logic
│   ├── handler/         # HTTP handlers
│   ├── model/           # Domain models
│   ├── repository/      # Data access
│   └── service/        # Business services
├── pkg/generator/       # Short code generator
├── frontend/            # Next.js frontend
│   ├── app/             # App Router pages
│   └── components/      # React components
└── PLAN.md              # Project plan
```