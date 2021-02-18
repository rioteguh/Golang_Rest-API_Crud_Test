# Golang_Rest-API_Crud_Test
Contoh aplikasi sederhana CRUD Rest-API menggunakan Bahasa Pemrograman Go

Cara Menggunakan:
1.  Install Go
2.  Atur GOPATH di environment agar bisa di running
3.  Download file aplikasi ini dengan cara di clone.
4.  go run main.go

Istilah dalam Golang
1. GOROOT adalah folder dimana berisi hasil instalasi file Golang. Misalnya di Linux yang terletak di /usr/lokal. Golang dapat digunakan ketika terdapat GOROOT.
2. GOBIN adalah Lokasi untuk meletakkan file Binari dari build projek File Golang.
3. GOOS adalah folder yang digunakan untuk mengkhususkan sistem operasi yang di gunakan
4. GOARCH adalah folder yang digunakan untuk mengkususkan / mensesifikasi arsitektur yang berupa proccessor. Parameter pada struktur folder ini tidak wajib di berikan.
5. GOPATH adalah folder yang digunakan untuk menaruh file / folder projek kita. Jika menggunakan XAMPP untuk menjalankan file PHP ibaratnya htdocs nya Golang.

Install Library Go Rest-API
// mysql
go get -u github.com/go-sql-driver/mysql
// mux
go get -u github.com/gorilla/mux
// felixge
go get -u github.com/gorilla/handlers
