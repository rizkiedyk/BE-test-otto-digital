````markdown
# Test OTT Digital BE

Ini adalah API backend untuk manajemen Voucher dan Brand yang dibangun menggunakan framework [Gin](https://github.com/gin-gonic/gin) dan Go.

## Fitur

1. **Brand Management**

   - Menambahkan Brand baru
   - Mengambil Brand berdasarkan ID
   - Mengambil semua Brand dengan filter dan pagination
   - Memperbarui Brand
   - Menghapus (soft delete) Brand

2. **Voucher Management**
   - Membuat Voucher baru
   - Mengambil Voucher berdasarkan ID
   - Mengambil Voucher berdasarkan ID Brand
   - Membuat Redemption Voucher
   - Mengambil Redemption Voucher berdasarkan Transaction ID

## Instalasi

### Persyaratan

- Go 1.21 atau lebih baru
- MongoDB atau database lain yang didukung (tergantung konfigurasi repository dan service)

### Langkah-langkah Instalasi

1. Clone repository ini:
   ```bash
   git clone https://github.com/rizkiedyk/BE-test-otto-digital.git
   cd BE-test-otto-digital
   ```
````

2. Instal dependensi Go:

   ```bash
   go mod tidy
   ```

3. Konfigurasi koneksi database di file konfigurasi (tergantung pada database yang Anda gunakan):

   - Edit `config/config.go` untuk menyesuaikan dengan pengaturan database Anda.

4. Jalankan aplikasi:

   ```bash
   go run main.go
   ```

5. Aplikasi akan berjalan di port default `1010`.

## Endpoints API

### Brand

- **POST /api/v1/brand**  
  Membuat Brand baru.  
  _Body request_:

  ```json
  {
    "name": "Brand Name",
    "price": "Brand price"
  }
  ```

- **GET /api/v1/brand/{brand_id}**  
  Mendapatkan detail Brand berdasarkan ID.

- **GET /api/v1/brand**  
  Mendapatkan daftar semua Brand dengan filter, sort, dan pagination.  
  _Query params_:

  - `page` (default: 1)
  - `limit` (default: 10)
  - `sort_by` (default: "created_at")
  - `sort_order` (default: "desc")
  - `filter_by_key` (optional)
  - `filter_by_value` (optional)

- **PATCH /api/v1/brand/{brand_id}**  
  Memperbarui informasi Brand berdasarkan ID.  
  _Body request_:

  ```json
  {
    "name": "Updated Brand Name",
    "price": "Updated Brand Price"
  }
  ```

- **DELETE /api/v1/brand/{brand_id}**  
  Menghapus (soft delete) Brand berdasarkan ID.

### Voucher

- **POST /api/v1/voucher**  
  Membuat Voucher baru.  
  _Body request_:

  ```json
  {
    "code": "code_in_here",
    "brand_id": "brand_id_here",
    "cost_in_point": 1000
  }
  ```

- **GET /api/v1/voucher?id={voucher_id}**  
  Mendapatkan detail Voucher berdasarkan ID.

- **GET /api/v1/voucher/brand?id={brand_id}**  
  Mendapatkan daftar Voucher berdasarkan ID Brand.

- **POST /api/v1/voucher/transaction/redemption**  
  Membuat Redemption Voucher.  
  _Body request_:

  ```json
  {
    "voucher_ids": "list of string"
  }
  ```

- **GET /api/v1/voucher/transaction/redemption?id={transaction_id}**  
  Mendapatkan Redemption Voucher berdasarkan Transaction ID.
