# Marshal & Unmarshal

## 1. membuat struck sesuai dengan struktur dari data yg di marshal
### 1.1. membuat stuck char untuk menampung data character
### 1.2. membuat struck resp untuk mengarahkan data yg mau di gunakan

## 2. membuat function pencarian dengan memanfaatkan dari inputan menggunakan package bufio dan data hasil dari marshal
### 2.1. mendeklarasi variabel penampung data dengan struck charackter yg sebeumnya di buat
### 2.2. memanfaatkan for lop untuk membaca seluruh isi dari data hasil marshal
#### 2.2.1. membuat kondisi untuk melakukan pencarian, mengecilkan seluruh huruf dari data dan inputan, kalo sama masukan data yg cocok ke dalam variabel yg di bikin sebelumnya

## 3. membuat function main
### 3.1. mendeklarasikan variabel untuk menampung data dengan scop main function
### 3.2. memanfaatkan package http dan method Get untuk mengambil data dari link
### 3.3. menghendle kondisi eror yg di kembalikan method Get
### 3.4. memanfaatkan package io untuk mengakses method ReadAll yg berfungsi membaca data hasil dari Get kalo ga pake ini hasilnya masih json
### 3.5. memanfaatkan unmarshal untuk mengconfert data dari json ke bit, unmarshal mengembalikan 2 hal eror dan data, eror ini harus di handel 
### 3.6. memanfaarkan bufio untuk interface dan menerima input
### 3.7. hasil input itu di kelola function pencarian yg sebelumnya di buat