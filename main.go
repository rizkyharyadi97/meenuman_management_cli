package main

import (
	"bufio"
	"fmt"
	"os"

	"meenuman/cli"
	"meenuman/config"
	"meenuman/handler"
	"meenuman/repository"
)

// main menghubungkan aplikasi dengan MySQL dan menjalankan Register/Login.
func main() {
	db, err := config.Connect()
	if err != nil {
		fmt.Println("Database gagal terhubung:", err)
		fmt.Println("Pastikan MySQL aktif dan database beverages_db sudah dibuat.")
		return
	}
	defer db.Close()

	in := bufio.NewReader(os.Stdin)

	userRepo := repository.MySQLUserRepository{DB: db}
	customerRepo := repository.MySQLCustomerRepository{DB: db}
	productRepo := repository.MySQLProductRepository{DB: db}
	transactionRepo := repository.MySQLTransactionRepository{DB: db}

	products := handler.ProductHandler{Repo: productRepo}
	customers := handler.CustomerHandler{Repo: customerRepo}
	orders := handler.OrderHandler{
		Products: productRepo,
		Customers: customerRepo,
		Transactions: transactionRepo,
	}

	for {
		fmt.Println("\n=== MEENUMAN ===")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih: ")

		switch readChoice(in) {
		case 1:
			user, err := handler.Register(userRepo, in)
			if err != nil {
				fmt.Println("Register gagal:", err)
				continue
			}

			fmt.Println("Register berhasil.")
			fmt.Println("Login sebagai:", user.Role)
			cli.Start(in, user, products, customers, orders)

		case 2:
			user, err := handler.Login(userRepo, in)
			if err != nil {
				fmt.Println("Login gagal: email/password salah")
				continue
			}

			fmt.Println("Login berhasil.")
			fmt.Println("Login sebagai:", user.Role)
			cli.Start(in, user, products, customers, orders)

		case 0:
			fmt.Println("Program selesai.")
			return

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

// readChoice membaca pilihan angka dari menu utama.
func readChoice(in *bufio.Reader) int {
	var value int
	_, _ = fmt.Fscanln(in, &value)
	return value
}
