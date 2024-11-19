# Test OTTO Digital BE

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
   git clone https://github.com/username/test-ottodigital-be.git
   cd test-ottodigital-be
   ```
