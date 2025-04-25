### 🧩 Daftar Layanan (Services) Nusaloka

| Service Name            | Port  | Fungsi Utama                                                                   |
| ----------------------- | ----- | ------------------------------------------------------------------------------ |
| **🚪 Gateway Service**  | 8080  | Reverse proxy semua request dari client, verifikasi JWT, route ke service lain |
| **👤 Account Service**  | 9001  | Autentikasi, registrasi, login, pengelolaan user profile, WebAuthn, Referral   |
| **🧳 Trip Service**     | 9002  | CRUD paket trip, itinerary, matching, pemesanan trip,  filter dan pencarian    |
| **💰 Finance Service**  | 9003  | Dompet digital, mutasi, pembayaran booking, integrasi ke payment gateway       |
| **🖼️ Media Service**   | 9004  | Upload media, kompresi (ffmpeg), resize gambar, video streaming (Reels)        |
| **💬 Realtime Service** | 9005  | Chat realtime, notifikasi, WebSocket connection                                |
| **📈 Observability**    | 9091+ | Prometheus, Grafana, Jaeger, Loki untuk metrics, logs, tracing                 |
| **🧠 Redis**            | 6379  | Cache, rate limiting, pub/sub untuk notifikasi                                 |
| **🗃️ PostgreSQL (DB)** | 5432  | Database utama semua services                                                  |
| **✉️ Mailpit**          | 8025  | Testing dan development untuk email (SMTP viewer)                              |

---

### 🛣️ Daftar Route per Service

#### 🚪 Gateway Service

- `POST /api/v1/auth/login`
- `POST /api/v1/auth/register`
- `GET /api/v1/account/profile`
- `GET /api/v1/trips`

#### 👤 Account Service

- `POST /auth/login`
- `POST /auth/register`
- `GET /account/profile`
- `PUT /account/profile`

#### 🧳 Trip Service

- `GET /trips`
- `GET /trips/:id`
- `POST /trips`
- `PUT /trips/:id`
- `DELETE /trips/:id`

#### 💰 Finance Service

- `GET /wallet/balance`
- `POST /wallet/topup`
- `GET /transactions`

#### 🖼️ Media Service

- `POST /media/upload`
- `GET /media/:id`
- `GET /media/stream/:id`

#### 💬 Realtime Service

- `GET /ws/chat`
- `GET /ws/notification`

#### 📈 Observability

- `GET /metrics` (per service)
- `GET /debug/pprof` (opsional)

#### ✉️ Mailpit

- `GET /`
- `GET /api/messages`

> Semua route bisa berubah sesuai kebutuhan bisnis dan arsitektur final.

