package main

import (
	"fmt"
	"os"
)

type Buku struct {
	Nama   string
	Jumlah int
}

type Pengguna struct {
	Username string
	NPM      string
	Gender   string
	FavFood  string
	FavDrink string
}

type Peminjaman struct {
	NamaBuku string
	Jumlah   int
}

var pengguna = Pengguna{
	Username: "Arafat",
	NPM:      "2406356864",
	Gender:   "Laki-laki",
	FavFood:  "Ayam Geprek",
	FavDrink: "Susu Segar",
}

var buku = []Buku{
	{"Pemrograman", 10},
	{"Film", 5},
	{"Printing", 20},
}

var historiPeminjaman []Peminjaman

func main() {
	fmt.Println("Selamat Datang di Perpustakaan Vokasi")
	Login()

	for {
		var pilihan int
		fmt.Println("\nMenu Program")
		fmt.Println("1. Lihat Informasi Pengguna Program")
		fmt.Println("2. Lihat Daftar Buku")
		fmt.Println("3. Tambah Daftar Buku")
		fmt.Println("4. Tambah Peminjaman Buku")
		fmt.Println("5. Histori Peminjaman Buku")
		fmt.Println("6. Keluar dari Program")
		fmt.Print("Input Menu yang anda inginkan: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			LihatInformasiPengguna()
		case 2:
			LihatDaftarBuku()
		case 3:
			TambahDaftarBuku()
		case 4:
			TambahPeminjamanBuku()
		case 5:
			HistoriPeminjamanBuku()
		case 6:
			fmt.Println("Terima kasih telah menggunakan program ini!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func Login() {
	var username, password string
	for {
		fmt.Print("Silahkan Input Username: ")
		fmt.Scanln(&username)
		fmt.Print("Silahkan Input Password: ")
		fmt.Scanln(&password)

		if username == pengguna.Username && password == pengguna.NPM {
			fmt.Println("Login Sukses !!!!")
			break
		} else {
			fmt.Println("Login Gagal. Coba lagi!")
			os.Exit(0)
		}
	}
}

func LihatInformasiPengguna() {
	fmt.Println("\nInformasi Pengguna Program")
	fmt.Printf("Username: %s\nNPM: %s\nJenis Kelamin: %s\nMakanan Favorit: %s\nMinuman Favorit: %s\n",
		pengguna.Username, pengguna.NPM, pengguna.Gender, pengguna.FavFood, pengguna.FavDrink)
}

func LihatDaftarBuku() {
	fmt.Println("\nDaftar Buku")
	for i, b := range buku {
		fmt.Printf("%d. Nama Buku: %s, Jumlah: %d\n", i+1, b.Nama, b.Jumlah)
	}
}

func TambahDaftarBuku() {
	var namaBuku string
	var jumlah int
	fmt.Print("Masukkan Nama Buku: ")
	fmt.Scanln(&namaBuku)
	fmt.Print("Masukkan Jumlah Buku: ")
	fmt.Scanln(&jumlah)
	buku = append(buku, Buku{Nama: namaBuku, Jumlah: jumlah})
	fmt.Println("Buku berhasil ditambahkan!")
}

func TambahPeminjamanBuku() {
	var pilihan, jumlahPinjaman int
	LihatDaftarBuku()
	fmt.Print("Masukan Nomor Pinjaman Buku (1,2,3): ")
	fmt.Scanln(&pilihan)

	if pilihan < 1 || pilihan > len(buku) {
		fmt.Println("Pilihan tidak valid!")
		return
	}

	fmt.Print("Masukkan Jumlah Pinjaman: ")
	fmt.Scanln(&jumlahPinjaman)

	if jumlahPinjaman <= 0 || jumlahPinjaman > buku[pilihan-1].Jumlah {
		fmt.Println("Pinjaman tidak valid. Jumlah melebihi stok atau angka minus!")
		return
	}

	buku[pilihan-1].Jumlah -= jumlahPinjaman
	historiPeminjaman = append(historiPeminjaman, Peminjaman{NamaBuku: buku[pilihan-1].Nama, Jumlah: jumlahPinjaman})
	fmt.Println("Buku berhasil dipinjam!")
}

func HistoriPeminjamanBuku() {
	if len(historiPeminjaman) == 0 {
		fmt.Println("Belum ada peminjaman.")
	} else {
		fmt.Println("\nHistori Peminjaman Buku")
		for i, h := range historiPeminjaman {
			fmt.Printf("%d. Nama Buku: %s, Jumlah: %d\n", i+1, h.NamaBuku, h.Jumlah)
		}
	}
}
