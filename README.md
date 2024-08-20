# sagara-msib-test

Backend Project-based Test

# Clothes Inventory API
Clothes Inventory API adalah sistem backend yang dibangun dengan menggunakan mode Clean Architecture untuk memastikan 
kode yang terstruktur dengan baik, modular, dan mudah untuk dikembangkan lebih lanjut. API ini dirancang untuk menangani 
berbagai operasi terkait manajemen inventaris baju, seperti pencarian, penambahan, pengurangan stok, dan manajemen data 
lainnya.

## Fitur Utama
* **CRUD untuk item baju,** Menyimpan, menampilkan, memperbarui, dan menghapus data baju dalam inventaris.
* **Pencarian Baju,** Mencari baju berdasarkan warna dan ukuran.
* **Manajemen Stok,** Menambahkan dan mengurangi stok baju, menampilkan baju yang stoknya habis, serta baju dengan stok kurang dari 5.


## Clean Architecture
API ini dibangun dengan mengikuti prinsip **Clean Architecture** yang memisahkan kode menjadi beberapa lapisan untuk menjaga
fleksibilitas dan kemudahan dalam pemeliharaan kode. Dengan struktur ini, logika bisnis, detail implementasi, dan antarmuka 
pengguna dapat berkembang secara independen satu sama lain.

## Tracing & Logging
Untuk memudahkan pemantauan dan debugging, API ini dilengkapi dengan **tracing** dan **logging** yang diimplementasikan menggunakan OpenTracing dan Jaeger-Client. Setiap permintaan yang masuk dan operasi yang dilakukan oleh API akan dilacak secara menyeluruh untuk memastikan transparansi dan deteksi kesalahan yang cepat.

## Main Libraries
* **OpenTracing,** Digunakan untuk implementasi tracing yang membantu dalam memantau kinerja dan perilaku API.
* **Jaeger-Client,** Library untuk mengirim dan mengelola data tracing yang dihasilkan oleh API ke sistem observability seperti Jaeger.