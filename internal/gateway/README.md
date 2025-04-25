```
nusaloka/
└── internal/
    └── gateway/
        ├── config/
        │   └── config.go              # Load env vars, setup base config
        ├── handler/
        │   └── proxy_handler.go       # Handler buat route ke microservices
        ├── router/
        │   └── gateway_router.go      # Setup semua route dari gateway
        ├── middleware/
        │   └── auth.go                # Verifikasi JWT, CORS, rate limit, dsb
        ├── observability/
        │   ├── metrics.go             # Prometheus metrics
        │   └── prometheus.go          # Setup handler `/metrics`
        └── service/
            └── reverse_proxy.go       # Reverse proxy ke tiap service
```