## 🚀 User Service

Service ini bertanggung jawab untuk pengelolaan akun pengguna, otentikasi, pendaftaran, dan pengelolaan data pribadi pengguna dalam sistem Nusaloka. Termasuk di dalamnya adalah login, pendaftaran via WebAuthn (FIDO2), serta sistem referral.

---

### 📦 Fitur

- **Login & Registrasi Pengguna**
  - Pendaftaran pengguna dengan username dan password
  - Autentikasi menggunakan username/password atau WebAuthn (FIDO2)
  
- **Autentikasi & Keamanan**
  - Penggunaan JWT-ES256 untuk otentikasi berbasis token
  - Pengelolaan token JWT untuk otentikasi pengguna pada endpoint yang dilindungi

- **Pengelolaan Profil Pengguna**
  - Dapatkan dan perbarui informasi profil pengguna
  
- **Sistem Referral**
  - Generate referral code unik untuk pengguna
  - Redeem referral code untuk mendapatkan keuntungan atau reward
  
- **Pengelolaan WebAuthn**
  - Mendukung login menggunakan biometrik (WebAuthn)
  - Mendukung pendaftaran WebAuthn dengan perangkat autentikasi seperti FIDO2

---

### 🛠 Tech Stack

- **Gin** – Web framework untuk routing HTTP dan middleware
- **GORM** – ORM (Object-Relational Mapping) untuk komunikasi dengan PostgreSQL
- **Redis** – Caching dan rate limiting
- **PostgreSQL** – Database utama untuk penyimpanan data
- **JWT-ES256** – Token autentikasi berbasis JSON Web Token (JWT)
- **WebAuthn (FIDO2)** – Standar autentikasi dua faktor berbasis biometrik

---

### 📁 Struktur Direktori

```bash
user/
├── cmd/
│   └── main.go             # Entry point untuk service ini
├── internal/
│   ├── dto/                # Data Transfer Objects (DTO) untuk request dan response
│   ├── handler/            # HTTP handlers yang menangani request & response
│   ├── model/              # Model GORM untuk entitas di database
│   ├── router/             # Routing Gin untuk endpoint
│   └── service/            # Logic aplikasi, seperti register, login, dan validasi
├── migration/              # Skrip migrasi database (jika ada perubahan struktur DB)
└── README.md               # Dokumentasi untuk service ini
```

---

### 🔐 Setup & Konfigurasi

1. **Instalasi Dependensi**

   Pastikan untuk menginstal dependensi yang diperlukan menggunakan Go Modules:

   ```bash
   go mod tidy
   ```

2. **File `.env`**

   Pastikan file `.env` Anda mengandung konfigurasi berikut:

   ```env
   USER_SERVICE_PORT=7001
   JWT_PRIVATE=shared/keys/private.pem
   JWT_PUBLIC=shared/keys/public.pem
   DATABASE_URL=postgres://username:password@localhost:5432/dbname
   REDIS_URL=redis://localhost:6379/0
   ```

   - **JWT_PRIVATE**: Path ke file private key untuk JWT (digunakan untuk signing token)
   - **JWT_PUBLIC**: Path ke file public key untuk JWT (digunakan untuk verification token)
   - **DATABASE_URL**: URL untuk menghubungkan ke PostgreSQL
   - **REDIS_URL**: URL untuk Redis, digunakan untuk caching dan rate-limiting

3. **Migrasi Database**

   Pastikan untuk menjalankan migrasi database untuk membuat tabel yang diperlukan:

   ```bash
   go run ./migration
   ```

   Atau gunakan GORM AutoMigrate jika tidak ada migrasi manual.

4. **Menjalankan User Service**

   Setelah semua tergantung sudah terinstal, jalankan server User Service menggunakan perintah:

   ```bash
   go run cmd/main.go
   ```

   Server akan berjalan di `localhost:7001` (atau port yang ditentukan di `.env`).

---

### 🌍 Endpoint API

#### 1. **POST /register**
   - Mendaftarkan pengguna baru.
   - Body request:
     ```json
     {
       "username": "evelin",
       "password": "secret123"
     }
     ```
   - Response:
     ```json
     {
       "message": "User successfully registered"
     }
     ```

#### 2. **POST /login**
   - Autentikasi pengguna dan menghasilkan token JWT.
   - Body request:
     ```json
     {
       "username": "evelin",
       "password": "secret123"
     }
     ```
   - Response:
     ```json
     {
       "token": "your_jwt_token_here"
     }
     ```

#### 3. **GET /profile**
   - Mengambil data profil pengguna.
   - Headers: `Authorization: Bearer <token>`
   - Response:
     ```json
     {
       "username": "evelin",
       "email": "evelin@example.com"
     }
     ```

#### 4. **POST /update-password**
   - Memperbarui password pengguna.
   - Body request:
     ```json
     {
       "password": "secret123",
       "new_password": "newSecret456",
       "new_password_confirmation": "newSecret456"
     }
     ```
   - Response:
     ```json
     {
       "message": "Password updated successfully"
     }
     ```

#### 5. **POST /update-profile**
   - Memperbarui informasi profil pengguna.
   - Body request:
     ```json
     {
       "email": "newemail@example.com"
     }
     ```
   - Response:
     ```json
     {
       "message": "Profile updated successfully"
     }
     ```

#### 6. **POST /web-authn/register**
   - Memulai proses pendaftaran WebAuthn.
   - Body request:
     ```json
     {
       "username": "evelin"
     }
     ```
   - Response:
     ```json
     {
       "registration_url": "https://example.com/webauthn/register"
     }
     ```

#### 7. **POST /web-authn/login**
   - Memulai proses login WebAuthn.
   - Body request:
     ```json
     {
       "username": "evelin"
     }
     ```
   - Response:
     ```json
     {
       "login_url": "https://example.com/webauthn/login"
     }
     ```

---

### 🧪 Testing API

**Contoh CURL Request untuk Register:**

```bash
curl -X POST http://localhost:7001/register -H "Content-Type: application/json" -d '{"username":"evelin", "password":"secret123"}'
```

**Contoh CURL Request untuk Login:**

```bash
curl -X POST http://localhost:7001/login -H "Content-Type: application/json" -d '{"username":"evelin", "password":"secret123"}'
```

**Contoh CURL Request untuk Mendapatkan Profil Pengguna (JWT):**

```bash
curl -H "Authorization: Bearer <your_jwt_token>" http://localhost:7001/profile
```

---

### 🧾 License

Apache License 2.0  
© 2025 Nusaloka Team
