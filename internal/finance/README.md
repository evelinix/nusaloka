# 💸 Nusaloka - Finance Service

Finance Service adalah bagian dari ekosistem **Nusaloka**, bertanggung jawab untuk mengelola **dompet digital**, **transaksi keuangan**, dan **laporan keuangan** pengguna. Layanan ini bersifat independen, tapi saling terintegrasi dengan layanan lain seperti Auth, Trip, dan Realtime.

---

## 🧠 Fitur Utama

- 🔐 Otentikasi via JWT (dari Auth Service)
- 💰 Manajemen saldo dompet pengguna
- 💳 Pencatatan dan pengelolaan transaksi (credit/debit)
- 📈 Laporan mutasi dan riwayat transaksi
- 📦 Kompatibel dengan sistem microservice berbasis gRPC atau REST
- 🧾 Pencatatan sistematis untuk integrasi audit/logging keuangan

---

## 📂 Struktur Direktori

```bash
services/finance
├── cmd
│   └── main.go              # Entry point service
├── internal
│   ├── config               # Setup env & config loader
│   ├── handler              # HTTP/gRPC handler
│   ├── model                # Model database (GORM)
│   ├── repository           # Abstraksi akses DB
│   ├── service              # Business logic utama
│   └── transport            # Layer REST/gRPC
├── proto                    # Protobuf definition (jika pakai gRPC)
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Environment Variables

| Variable             | Keterangan                            | Contoh                     |
|----------------------|----------------------------------------|----------------------------|
| `PORT`               | Port service ini berjalan              | `8084`                     |
| `DATABASE_URL`       | URL koneksi PostgreSQL                 | `postgres://user:pass@localhost:5432/finance_db` |
| `JWT_PUBLIC_KEY`     | Path ke public key untuk verifikasi JWT| `./keys/public.pem`        |
| `ENV`                | Mode aplikasi (`dev`, `prod`)          | `dev`                      |

---

## 🚀 Cara Menjalankan

### 1. Jalankan dengan Docker Compose (rekomendasi)

```bash
docker-compose up finance
```

### 2. Jalankan manual via Go

```bash
cd services/finance
go run cmd/main.go
```

---

## 🧪 Testing

```bash
go test ./... -v
```

> ⚠️ Gunakan database khusus testing, jangan produksi.

---

## 📌 API Overview

### 🔐 Authorization
- Semua endpoint butuh header `Authorization: Bearer <token>` dari Auth Service.

### 📤 POST /wallets/topup
> Menambahkan saldo ke dompet user

```json
{
  "user_id": "uuid",
  "amount": 50000,
  "source": "bank_transfer"
}
```

### 📤 POST /wallets/withdraw
> Mengurangi saldo user

```json
{
  "user_id": "uuid",
  "amount": 25000,
  "target": "bank_account"
}
```

### 📥 GET /wallets/:user_id/balance
> Mendapatkan saldo terbaru user

### 📥 GET /wallets/:user_id/transactions
> Mendapatkan riwayat transaksi user (terbaru sampai lama)

---

## 🛠 Stack Teknologi

- **Golang** + [GORM](https://gorm.io)
- PostgreSQL [GORM/POSTGRES](https://gorm.io)
- JWT (ES256 via public/private key)
- Redis untuk caching balance
- gRPC (optional, jika diaktifkan)

---

## 🔒 Security Notes

- Seluruh transaksi diverifikasi via JWT
- Nominal transaksi dibatasi dan tervalidasi
- Transaksi bersifat **atomic**, menggunakan DB transaction
- Future: Logging transaksi ke external Audit Service

---

## 📌 TODO / Roadmap

- [ ] Integrasi dengan Midtrans / Payment Gateway
- [ ] Notifikasi realtime ke user (via Realtime Service)
- [ ] Export laporan transaksi dalam format CSV/Excel
- [ ] Sistem refund & dispute

---

## 🧑‍💻 Developer Notes

### Format Commit:
Gunakan [Conventional Commits](https://www.conventionalcommits.org/):

```
feat(finance): add feature balance endpoint
fix(finance): fix rounding bug on topup
```

### Struktur Branch:
- `main` = production
- `dev` = development
- `feature/xyz`, `fix/bug-name`

---

## 🤝 Kontribusi

Pull request selalu diterima! Pastikan ikuti format coding dan testing yang sudah ditentukan.

---

## 📄 Lisensi

Licensed under Apache 2.0 © 2025 Nusaloka

```
