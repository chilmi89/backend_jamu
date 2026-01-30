# API Documentation - Jamu Nusantara

Dokumentasi ini berisi detail teknis untuk setiap endpoint API yang tersedia di Jamu Nusantara Backend.

**Base URL**: `http://localhost:8080`

---

## 1. Authentication

### Login
Digunakan untuk mendapatkan token JWT agar bisa mengakses fitur Admin.

**Admin Default (dari seeder):**
- Email: `admin@gmail.com`
- Password: `admin123`

- **Endpoint**: `POST /api/login`
- **Headers**: 
  - `Content-Type: application/json`
- **Request Body**:
  ```json
  {
    "email": "admin@gmail.com",
    "password": "admin123"
  }
  ```
- **Response (Success 200)**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "Admin Nusantara",
      "email": "admin@gmail.com",
      "role": "admin"
    }
  }
  ```

---

## 2. Products

### List Products
Mengambil daftar semua jamu yang tersedia.
- **Endpoint**: `GET /api/products`
- **Auth**: Public

### Add Product
Menambahkan produk baru.
- **Endpoint**: `POST /api/products`
- **Auth**: **Admin JWT**
- **Request Body**:
  ```json
  {
    "name": "Beras Kencur",
    "description": "Menambah nafsu makan.",
    "price": 12000,
    "image": "beras_kencur.jpg"
  }
  ```

### Update Product
Memperbarui data produk.
- **Endpoint**: `PUT /api/products?id=<id>`
- **Auth**: **Admin JWT**

### Delete Product
Menghapus produk.
- **Endpoint**: `DELETE /api/products?id=<id>`
- **Auth**: **Admin JWT**

---

## 3. News

### List News
Mengambil semua berita.
- **Endpoint**: `GET /api/news`
- **Auth**: Public

### Add News
Menambahkan berita baru.
- **Endpoint**: `POST /api/news`
- **Auth**: **Admin JWT**
- **Request Body**:
  ```json
  {
    "title": "Manfaat Temulawak",
    "content": "Temulawak sangat baik untuk hati..."
  }
  ```

### Update News
Memperbarui berita.
- **Endpoint**: `PUT /api/news?id=<id>`
- **Auth**: **Admin JWT**
- **Request Body**:
  ```json
  {
    "title": "Judul Baru",
    "content": "Konten baru..."
  }
  ```

### Delete News
Menghapus berita.
- **Endpoint**: `DELETE /api/news?id=<id>`
- **Auth**: **Admin JWT**

---

## Panduan Pengetesan (cURL)

**1. Login (Dapatkan Token):**
```bash
curl -X POST http://localhost:8080/api/login \
     -H "Content-Type: application/json" \
     -d '{"email":"admin@gmail.com", "password":"admin123"}'
```

**2. List Products:**
```bash
curl -X GET http://localhost:8080/api/products
```

**3. Tambah Produk (Ganti TOKEN):**
```bash
curl -X POST http://localhost:8080/api/products \
     -H "Authorization: Bearer <TOKEN_ANDA>" \
     -H "Content-Type: application/json" \
     -d '{"name":"Jamu Test", "price":5000}'
```
