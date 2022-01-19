# GolangTest

<h3> Golang Test :</h3>

   <p>Aplikasi ini dibuat untuk memenuhi backend enginer test case sagara technology menggunakan bahasa pemrograman golang dengan framework Fiber dan Gorm. untuk autentikasi menggunakan JWT (Json Web Token) yang disimpan pada cookies ketika telah login.</p>

### How To Get Started With API:

<p>untuk menjalankan aplikasi ini :</p>
<p>1. clone atau download aplikasi ini</p>
<p>2. membuat sebuah database kosong bernama golangtest, tidak perlu membuat table dikarenakan di dalam aplikasi terdapat AutoMigrate</p>
<p>3. konfigur sesuai kebutuhan pengguna pada ./database/connection.go seperti dibawah ini</p>

#### Connection.go:
    const DB_username = "root"
    const DB_password = ""
    const DB_name = "golangtest"
    const DB_host = "127.0.0.1"
    const DB_port = "3306"
    
#### Collection API (POSTMAN):
    https://documenter.getpostman.com/view/15996687/UVXnJEp2

### Implemented REST API:

#### Register Account:
    POST: localhost:8000/register

#### Signin:
    POST: localhost:8000/login

#### Get All Produk:
    POST: localhost:8000/api/allproduk
    
#### Create Produk:
    POST: localhost:8000/api/createproduk
    
#### Update Produk:
    POST: localhost:8000/api/updateproduk/:id
    
#### Delete Produk:
    POST: localhost:8000/api/deleteproduk/:id
    
#### Get Produk By Id:
    POST: localhost:8000/api/produk/:id
    
#### Upload Image:
    POST: localhost:8000/api/uploadimage
