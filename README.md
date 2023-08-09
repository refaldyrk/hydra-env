
# Hydra Environment Manager

Hydra Environment Manager adalah alat baris perintah yang memungkinkan Anda mengelola konfigurasi lingkungan secara aman menggunakan pasangan kunci-nilai yang dienkripsi. Alat ini memungkinkan Anda untuk membuat, mengubah, dan menghapus kunci sambil memastikan nilai-nilai dienkripsi, menjaga privasi data.

## Fitur

- Menghasilkan dan mencetak kunci enkripsi yang aman.
- Membuat, membaca, memperbarui, dan menghapus pasangan kunci-nilai dalam berkas lingkungan JSON.
- Melakukan enkripsi dan dekripsi nilai untuk keamanan yang lebih baik.

## Memulai

1. Install aplikasi:

   ```shell
   go install github.com/refaldyrk/hydra-env@latest
   ```

2. Jalankan aplikasi:

   ```shell
   # Generate encryption key
   hydra-env --gen-key

   # Create an environment file (if not set)
   hydra-env --env=mycustomenv.json

   # Add a new key-value pair to the environment file
   hydra-env --env=mycustomenv.json --add-key="api_key|my_secret_key"

   # Retrieve the value of a key from the environment file
   hydra -env--env=mycustomenv.json --get-key=api_key

   # List all keys in the environment file
   hydra-env --env=mycustomenv.json --list-keys

   # Load all keys in the environment file
   hydra-env --env=mycustomenv.json --load-env=example.env

   # Delete a key from the environment file
   hydra-env --env=mycustomenv.json --del-key=api_key
   ```

## Penggunaan

- Untuk menghasilkan kunci enkripsi baru:

   ```shell
   hydra-env --gen-key
   ```

- Untuk menambahkan pasangan kunci-nilai baru ke berkas lingkungan:

   ```shell
   hydra-env --env=mycustomenv.json --add-key="api_key|my_secret_key"
   ```

- Untuk mendapatkan nilai dari kunci dalam berkas lingkungan:

   ```shell
   hydra-env --env=mycustomenv.json --get-key=api_key
   ```

- Untuk menampilkan daftar semua kunci dalam berkas lingkungan:

   ```shell
   hydra-env --env=mycustomenv.json --list-keys
   ```

- Untuk menambahkan daftar semua kunci dalam berkas lingkungan:

   ```shell
   hydra-env --env=mycustomenv.json --load-env=example.env
   ```

- Untuk menghapus kunci dari berkas lingkungan:

   ```shell
   hydra-env --env=mycustomenv.json --del-key=api_key
   ```

## Kontribusi

Kontribusi sangat diterima! Jika Anda menemukan bug atau memiliki saran perbaikan, jangan ragu untuk membuka *issue* atau mengajukan *pull request*.

....
