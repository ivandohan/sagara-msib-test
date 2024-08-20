# sagara-msib-test

Backend Project-based Test

# Clothes Inventory API
Clothes Inventory API adalah sistem backend yang dibangun dengan menggunakan mode Clean Architecture untuk memastikan 
kode yang terstruktur dengan baik, modular, dan mudah untuk dikembangkan lebih lanjut. API ini dirancang untuk menangani 
berbagai operasi terkait manajemen inventaris baju, seperti pencarian, penambahan, pengurangan stok, dan manajemen data 
lainnya.

## Fitur Utama
* Dapat dieksekusi dengan **build image docker**.
* **CRUD untuk item baju,** menyimpan, menampilkan, memperbarui, dan menghapus data baju dalam inventaris.
* **Pencarian Baju,** mencari baju berdasarkan nama, brand, warna, dan ukuran.
* **Manajemen Stok,** menambahkan dan mengurangi stok baju, menampilkan baju yang stoknya habis, serta baju dengan stok kurang dari 5.

## Docker Setup
* Image yang digunakan dalam proyek adalah *golang:1.22 dan postgres:15
* Tidak perlu setup tambahan, semua keperluan eksekusi sudah saya paketkan di Dockerfile
* Eksekusi kode berikut untuk memulai program :
```text
    // One shot execution
    docker-compose up --build
    
    // Separated execution
    docker-compose up db
    // Wait till this db image worked properly, and then
    docker-compose up app
```

## Clean Architecture
API ini dibangun dengan mengikuti prinsip **Clean Architecture** yang memisahkan kode menjadi beberapa lapisan untuk menjaga
fleksibilitas dan kemudahan dalam pemeliharaan kode. Dengan struktur ini, logika bisnis, detail implementasi, dan antarmuka 
pengguna dapat berkembang secara independen satu sama lain.

## Tracing & Logging
Untuk memudahkan pemantauan dan debugging, API ini dilengkapi dengan **tracing** dan **logging** yang diimplementasikan 
menggunakan OpenTracing dan Jaeger-Client. Setiap permintaan yang masuk dan operasi yang dilakukan oleh API akan dilacak 
secara menyeluruh untuk memastikan transparansi dan deteksi kesalahan yang cepat.

## Main Libraries
* **OpenTracing,** digunakan untuk implementasi tracing yang membantu dalam memantau kinerja dan perilaku API.
* **Jaeger-Client,** library untuk mengirim dan mengelola data tracing yang dihasilkan oleh API ke sistem observability seperti Jaeger.