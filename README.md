# Test Golang
Ini adalah repository untuk menjawab beberapa soal test golang.
<br>
[Dokumentasi API ](https://documenter.getpostman.com/view/1720373/SW18wurk?version=latest)
## Instalasi
#### Clone project
```sh
$ git clone https://github.com/fachryansyah/test-golang
```
#### Memasang Module
```sh
$ go get github.com/fachryansyah/test-golang-solusi-menghitung-huruf
```
#### Menjalankan Webservice
```sh
$ go run server.go
```

## Menjalankan Service Docker
#### Membuat Image
Membuat image dari Dockerfile yang sudah di konfigurasi.
```sh
$ docker build --tag app-test-golang-fahri:1.0 .
```
#### Membuat Container
```sh
docker container create --name app-test-golang-fahri -p 1337:1337 app-test-golang-fahri:1.0
```
#### Menjalankan Container
```sh
docker container start app-test-golang-fahri
```
