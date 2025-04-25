<h1 align="center">
  <img src="assets/logo.png" alt="Nusaloka Logo" width=175 />
</h1>

# 🌴 Nusaloka

![Go](https://img.shields.io/badge/Go-1.24-blue.svg)
![Flutter](https://img.shields.io/badge/Flutter-3.29-blue.svg)
![License](https://img.shields.io/badge/License-Apache%202.0-green)

**Nusaloka** is a modular, microservice-based tour & travel marketplace platform where users can request custom travel packages and get matched with tour guides based on interests, budget, and location.  
Built with Go, Flutter, PostgreSQL, Redis, and gRPC — with a real-time twist 🌀

---

## ✨ Features

- 🔐 Auth system (JWT + WebAuthn + Passkey ready)
- 🧭 User-generated travel requests & itinerary matching (>= 80% match)
- 📅 Booking system with tracking, cancellation, refund
- 💰 Wallet with topup/withdraw & integrated payment
- 🎥 Short video reels (HLS via ffmpeg-go)
- 💬 Real-time Chat & Notifications (WebSocket)
- 🧑‍💻 Admin dashboard & content moderation
- 📦 Optional Delivery system for souvenirs

---

## 🧱 Tech Stack

| Layer         | Tech                                               |
|---------------|----------------------------------------------------|
| Backend       | Go (Cai), GORM, gRPC, PostgreSQL, Redis            |
| Realtime      | WebSocket (for chat & bidding notifications)       |
| Video         | ffmpeg-go, S3 compatible storage                   |
| Frontend      | Flutter (Cupertino-style, bloc, dio, go_router)    |
| Docs          | VitePress (modular docs per service/module)        |
| CI/CD         | GitHub Actions                                     |
| Messaging     | Redis Pub/Sub                                      |
| Font          | [Nerd Font](https://www.nerdfonts.com/#home)       |
---

## 🧬 Architecture Overview

- **Microservices**:
  - Public API via REST (API Gateway)
  - Internal RPC via gRPC (with protobufs via Buf)
  - Realtime via WebSocket
- **Shared Services**:
  - NotificationService, MatchingService, AdminService, etc.
- **Storage**:
  - PostgreSQL (Relational DB)
  - Redis (Cache, Session, PubSub)
  - S3 (Video Upload & Streaming)

---

## 📁 Project Structure (Simplified)

```bash
nusaloka/
│
├── cmd/                            # Entry point tiap service (main.go)
│   ├── gateway/                    # API Gateway
│   │   └── main.go
│   ├── account/                    # Account Service (gabungan user + auth)
│   │   └── main.go
│   ├── trip/                       # Trip Service
│   │   └── main.go
│   ├── finance/                    # Finance Service
│   │   └── main.go
│   ├── media/                      # Media Service
│   │   └── main.go
│   └── realtime/                   # Realtime Service
│       └── main.go
│
├── internal/                       # Core logic tiap service
│   ├── gateway/
│   │   ├── config/
│   │   │   └── config.go              # Load env vars, setup base config
│   │   ├── handler/
│   │   │   └── proxy_handler.go       # Handler buat route ke microservices
│   │   ├── router/
│   │   │   └── gateway_router.go      # Setup semua route dari gateway
│   │   ├── middleware/
│   │   │   └── auth.go                # Verifikasi JWT, CORS, rate limit, dsb
│   │   ├── observability/
│   │   │   ├── metrics.go             # Prometheus metrics
│   │   │   └── prometheus.go          # Setup handler `/metrics`
│   │   └── service/
│   │       └── reverse_proxy.go       # Reverse proxy ke tiap service
│   ├── account/
│   │   ├── config/                 # Konfigurasi (env loader, dsb)
│   │   │   └── config.go
│   │   ├── handler/                # HTTP handler / controller
│   │   │   └── handler.go
│   │   ├── middleware/             # Middleware: auth, logging, dsb
│   │   │   └── jwt.go
│   │   ├── model/                  # Model database
│   │   │   └── user.go
│   │   ├── repository/             # DB layer
│   │   │   └── user_repo.go
│   │   ├── router/                 # Router per service
│   │   │   └── account_router.go
│   │   ├── service/                # Business logic
│   │   │   └── user_service.go
│   │   └── observability/         # Monitoring tools (metrics, log, tracing)
│   │       ├── metrics.go         # Custom Prometheus metrics
│   │       └── prometheus.go      # Setup Prometheus HTTP handler
│   └── ... (trip/finance/media/realtime mirip structure-nya)
│
├── pkg/                            # Package yang bisa dipakai lintas service
│   ├── logger/                     # Zerolog setup
│   │   └── logger.go
│   ├── tracing/                    # OpenTelemetry setup
│   │   └── otel.go
│   └── middleware/                # Reusable middleware
│       └── cors.go
│
├── api/                            # Swagger/OpenAPI docs
│   └── openapi.yaml
│
├── docker/                         # Dockerfile per service
│   ├── gateway/Dockerfile
│   ├── account/Dockerfile
│   ├── trip/Dockerfile
│   ├── finance/Dockerfile
│   ├── media/Dockerfile
│   └── realtime/Dockerfile
│
├── deploy/                         # Deployment config (docker-compose, k8s)
│   ├── docker-compose.yml
│   ├── prometheus/
│   │   └── prometheus.yml
│   └── grafana/                   # Custom dashboards if any
│       └── datasources.yml
│
├── .env                            # Semua ENV variabel port, redis, db, etc.
├── go.mod
├── Makefile                        # Build/run helper
└── README.md
```

---
## 📜 License

```txt
Copyright 2025 evelinix@Nusaloka

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
---
<p align="center">Feel free to contribute, PR, or open issues. ✨  </p>
<p align="center">Made with 💖 and a lot of coffee ☕️ by Team Nusaloka.</p>
