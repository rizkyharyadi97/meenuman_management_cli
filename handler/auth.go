package handler

import (
	"bufio"
	"errors"
	"fmt"
	"strings"

	"meenuman/entity"
	"meenuman/repository"
)

// Register membuat akun baru dan meminta user memilih role.
func Register(repo repository.UserRepository, in *bufio.Reader) (entity.User, error) {
	fmt.Println("\n=== REGISTER ===")

	fmt.Print("Email: ")
	email := readLine(in)

	fmt.Print("Password: ")
	password := readLine(in)

	role := chooseRole(in)
	if role == "" {
		return entity.User{}, errors.New("registrasi dibatalkan")
	}

	u := entity.User{
		Email: email,
		Password: password,
		Role: role,
	}

	if err := repo.Register(u); err != nil {
		return u, err
	}

	return repo.Login(email, password)
}

// Login memeriksa email dan password user.
func Login(repo repository.UserRepository, in *bufio.Reader) (entity.User, error) {
	fmt.Println("\n=== LOGIN ===")

	fmt.Print("Email: ")
	email := readLine(in)

	fmt.Print("Password: ")
	password := readLine(in)

	return repo.Login(email, password)
}

// chooseRole hanya memberikan dua role: Admin dan Kepala Toko.
func chooseRole(in *bufio.Reader) string {
	for {
		fmt.Println("\nPilih Role:")
		fmt.Println("1. Admin")
		fmt.Println("2. Kepala Toko")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		switch readInt(in) {
		case 1:
			return "admin"
		case 2:
			return "kepala_toko"
		case 0:
			return ""
		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

// readLine membaca satu baris input.
func readLine(in *bufio.Reader) string {
	text, _ := in.ReadString('\n')
	return strings.TrimSpace(text)
}

// readInt membaca input angka.
func readInt(in *bufio.Reader) int {
	var value int
	fmt.Sscan(readLine(in), &value)
	return value
}
