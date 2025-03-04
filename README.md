# Backend Engineer Assessment - Golang API Service

## 📌 Deskripsi  
Proyek ini adalah REST API untuk layanan perbankan sederhana, yang memungkinkan:  
✅ Registrasi Nasabah  
✅ Menabung  
✅ Menarik Dana  
✅ Cek Saldo  

Aplikasi ini dibangun menggunakan **Golang, PostgreSQL, dan Docker** untuk deployment.  

---

## 🛠 Teknologi yang Digunakan  
- **Golang** (Echo Framework)  
- **PostgreSQL** (Database)  
- **GORM** (ORM untuk Golang)  
- **Docker & Docker Compose** (Deployment)  
- **Logrus** (Logging)  
- **Argument Parser (`flag`)** (Konfigurasi REST API host & port)  

---

## 📂 Struktur Proyek
```
backend-service/
│── cmd/                # Entry point aplikasi (main.go)
│── config/             # Konfigurasi database, logger, dan argument parser
│── internal/           # Business logic & domain services
│   ├── controllers/    # Handler untuk request API
│   ├── repositories/   # Operasi database (CRUD) dengan interface
│   ├── services/       # Business logic dengan interface
│── models/             # Struct model database
│── routes/             # Routing API
│── pkg/                # Utility/helper functions
│── go.mod              # Module Go
│── Dockerfile          # Konfigurasi Docker
│── docker-compose.yml  # Orkestrasi Docker container
│── .env                # Environment variables
```

---

## 🚀 Menjalankan Aplikasi  

### 1️⃣ Clone Repository  
```sh
git clone https://github.com/DickiAndriyadi/assesment-isi.git
cd assesment-isi
```
### 2️⃣ Setup Environment Variables

``` 
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=bank
```

## 📌 Menjalankan dengan Golang (Tanpa Docker)

Jalankan perintah berikut:

```
go run cmd/main.go -host=127.0.0.1 -port=9000
```

Output yang diharapkan:

```
Running on 127.0.0.1:9000
Server berjalan di 127.0.0.1:9000

```

---

## 🐳 Menjalankan dengan Docker

### 1️⃣ Build & Run Docker Compose

```
docker-compose up --build
```

Jika berhasil, akan muncul:

```
Connected to database successfully!
Database migrated successfully!
Server berjalan di 0.0.0.0:8080
```

### 2️⃣ Cek Container yang Berjalan

```
docker ps
```

Jika PostgreSQL dan aplikasi berjalan, outputnya seperti ini:

```
CONTAINER ID   IMAGE         PORTS                    NAMES
abc123456789   backend-app   0.0.0.0:8080->8080/tcp   backend-container
```

### 3️⃣ Menghentikan Container

```
docker-compose down
```

---

## 🔎 Pengujian API dengan Postman atau cURL

### 📌 1️⃣ Registrasi Nasabah
- Endpoint: POST /api/daftar
- Request Body:
```
{
    "nama": "Dicki Andriyadi",
    "nik": "1234567890",
    "no_hp": "081234567890"
}
```

- Response (200 OK):

```
{
    "no_rekening": "99A1B2C3D"
}
```


### 📌 2️⃣ Menabung
- Endpoint: POST /api/tabung
- Request Body:

```
{
    "no_rekening": "99A1B2C3D",
    "nominal": 500000
}
```

- Response (200 OK) :

```
{
    "saldo": 500000
}
```

### 📌 3️⃣ Penarikan Dana
- Endpoint: POST /api/tarik
- Request Body:

```
{
    "no_rekening": "99A1B2C3D",
    "nominal": 200000
}
```

- Response (200 OK) :

```
{
    "saldo": 500000
}
```

### 📌 4️⃣ Cek Saldo
- Endpoint: GET /api/saldo/99A1B2C3D
- Response (200 OK):

```
{
    "saldo": 300000
}
```

### 📌 Deployment ke Production
Jika ingin deploy ke VPS atau cloud, jalankan:

```
docker-compose up --build -d
```

---

`-d` = Detached mode (jalankan di background).

## ✅ Final Checklist
### ✅ Clean Code (Struktur modular & best practices)
### ✅ Interface untuk Controllers, Services, dan Repositories
### ✅ Argument Parser untuk REST API host & port
### ✅ Logging dengan Logrus
### ✅ Database PostgreSQL dalam Docker
### ✅ Deployment menggunakan Dockerfile & Docker Compose
### ✅ Pengujian dengan Postman berhasil

## 📌 Penutup
Aplikasi ini telah dibangun dengan pendekatan Clean Code, Interface, serta Docker untuk deployment.
Terima kasih telah mengunjungi proyek ini! Jika ada pertanyaan atau perbaikan, silakan buat issue atau pull request. 🚀

## 📌 Dibuat oleh:
Dicki Andriyadi







