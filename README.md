# Assigment 8

Projek ini adalah sebuah API sederhana untuk mengupdate status air dan angin. API ini menyediakan endpoint untuk mengupdate status dan mendapatkan status dari database.

## Instalasi

1. Clone repository ini:
```bash
git clone <repository-url>
```
2. Masuk ke direktori proyek:
```bash
cd Assigment_8\cmd\server
```
3. Install dependensi yang diperlukan:
```bash
go mod download
```
5. Jalankan server:
```bash
go run main.go
```


API dapat diakses di `http://localhost:<nomor-port>/api/v1/status`.

## Penggunaan

### Update Status

**Endpoint:** `PUT /api/v1/status`

Endpoint ini digunakan untuk mengupdate status air dan angin. Body permintaan harus mencakup field berikut:

```json
{
"wind": 10,
"water": 4
}
```
wind: Status angin (nilai integer).

water: Status air (nilai integer).

Respon sukses akan memiliki kode status 200 dan akan menyertakan status terupdate di dalam body respon.

Format Respon
API ini mengikuti format respon yang konsisten. 

Body respon mencakup field-field berikut:

code: Kode status HTTP.

status: Status permintaan ("Berhasil" untuk sukses atau "Gagal" untuk kegagalan).

message: Pesan singkat yang menjelaskan hasil permintaan.

message_detail: Detail tambahan tentang hasil permintaan.

data: Data yang dikembalikan oleh permintaan (jika ada).








