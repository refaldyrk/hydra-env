# Hydra-Env

Hydra-Env adalah alat baris perintah (command-line tool) untuk mengelola variabel lingkungan (environment variables) dan kunci (keys), serta memiliki kemampuan integrasi dengan MongoDB. Alat ini memungkinkan Anda dengan mudah membuat, mengelola, dan memuat variabel lingkungan dari berkas, menghasilkan dan menyimpan kunci dengan aman, serta berinteraksi dengan server MongoDB untuk penyimpanan data.

## Instalasi

Untuk menginstal Hydra-Env, ikuti langkah-langkah berikut:


1. Buat berkas biner eksekusi:

   ```bash
   go install github.com/refaldyrk/hydra-env@latest
   ```


## Penggunaan

Hydra-Env menyediakan serangkaian perintah untuk mengelola variabel lingkungan, kunci, dan berinteraksi dengan server MongoDB. Berikut adalah perintah-perintah yang tersedia:

### Menghasilkan dan Mencetak Kunci Baru

Menghasilkan kunci baru dan mencetaknya ke konsol.

```bash
hydra-env -gen-key
```

### Menambahkan Kunci Baru

Menambahkan kunci baru ke berkas lingkungan.

```bash
hydra-env -env=path/ke/berkas/env -add-key=namaKunci|nilaiKunci
```

### Mendapatkan Nilai Kunci

Mendapatkan nilai dari kunci tertentu dalam lingkungan.

```bash
hydra-env -env=path/ke/berkas/env -get-key=namaKunci
```

### Daftar Kunci

Menampilkan daftar semua kunci yang ada dalam berkas lingkungan.

```bash
hydra-env -env=path/ke/berkas/env -list-keys
```

### Menghapus Kunci

Menghapus kunci tertentu dari lingkungan.

```bash
hydra-env -env=path/ke/berkas/env -del-key=namaKunci
```

### Memuat Lingkungan dari Berkas

Memuat variabel lingkungan dari berkas dan menambahkannya ke lingkungan.

```bash
hydra-env -env=custom.json -load-env=path/ke/berkas/env
```

### Perintah Server

Berinteraksi dengan server MongoDB menggunakan perintah-perintah berikut:

- Pengecekan koneksi server:

```bash
hydra-env  -server=ping -env=path/ke/berkas/env
```

- Menampilkan bantuan perintah server:

```bash
hydra-env -server=help
```

## Konfigurasi

Sebelum menggunakan Hydra-Env, pastikan Anda telah mengatur variabel lingkungan `HYDRA_MONGO_SERVER` untuk menentukan URL server MongoDB.

```bash
export HYDRA_MONGO_SERVER=mongodb://localhost:27017
```

## Contoh

1. Menghasilkan dan mencetak kunci baru:

```bash
hydra-env -gen-key
```

2. Menambahkan kunci baru:

```bash
hydra-env -env=contoh.json -add-key=KUNCI_API|nilai-kunci-api-anda
```

3. Mendapatkan nilai kunci:

```bash
hydra-env -env=contoh.json -get-key=KUNCI_API
```

4. Menampilkan daftar kunci:

```bash
hydra-env -env=contoh.json -list-keys
```

5. Menghapus kunci:

```bash
hydra-env -env=contoh.json -del-key=KUNCI_API
```

6. Memuat lingkungan dari berkas:

```bash
hydra-env -load-env=berkas_env.env
```

7. Berinteraksi dengan server MongoDB - Pengecekan koneksi:

```bash
hydra-env -server=ping -env=contoh.json
```

8. Berinteraksi dengan server MongoDB - Mengimpor kunci dan data lingkungan:

```bash
hydra-env -server=import -env=contoh.json
```

9. Menampilkan bantuan perintah server:

```bash
hydra-env -server=help
```

## Kontak

Untuk pertanyaan, silakan hubungi [Email](mailto:refaldy.rizky22@gmail.com).
