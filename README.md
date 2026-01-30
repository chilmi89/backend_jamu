# Jamu Nusantara Simple API

Backend API sederhana menggunakan Go (Golang), Bun ORM, dan MySQL untuk aplikasi Jamu Nusantara.

## Fitur Baru
- **Bun ORM**: Query lebih aman dan clean.
- **Auto Seeder**: Akun admin otomatis dibuat saat aplikasi dijalankan.

## Persyaratan
- Go 1.21+
- MySQL Server

## Cara Menjalankan

1.  **Persiapan Database:**
    - Buat database dengan nama `jamu_nusantara`.
    - Impor schema dari file `schema.sql`.
    - (Opsional) Tambahkan user admin manual ke tabel `users` untuk mencoba fitur Admin. Gunakan tool bcrypt untuk menghash password.

2.  **Konfigurasi environment:**
    - Sesuaikan file `.env` dengan kredensial database Anda.

3.  **Install Dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Jalankan Server:**
    ```bash
    go run main.go
    ```
    Server akan berjalan di `http://localhost:8080`.

## API Endpoints

### Auth
- `POST /api/login` - Login untuk mendapatkan token JWT. Body: `{ "email": "...", "password": "..." }`

### Products
- `GET /api/products` - List semua produk.
- `POST /api/products` - Tambah produk baru (Memerlukan Admin JWT).

### News
- `GET /api/news` - List berita terbaru.
- `POST /api/news` - Tambah berita baru (Memerlukan Admin JWT).

## Struktur Proyek
- `config/`: Koneksi database.
- `handlers/`: Logika REST API.
- `middleware/`: Auth & CORS.
- `models/`: Definisi data (Go Structs).
- `utils/`: JWT & Password hashing.
