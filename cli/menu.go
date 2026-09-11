package cli

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"meenuman/entity"
	"meenuman/handler"
)

// Start menjalankan menu berdasarkan role akun yang sedang login.
// Setelah role ditentukan, user tidak dapat berpindah ke role lain.
func Start(
	in *bufio.Reader,
	user entity.User,
	products handler.ProductHandler,
	customers handler.CustomerHandler,
	orders handler.OrderHandler,
) {
	switch user.Role {
	case "admin":
		adminMenu(in, products, customers, orders)
	case "kepala_toko":
		kepalaTokoMenu(in, products, customers, orders)
	default:
		fmt.Println("Role tidak dikenali.")
	}
}

// adminMenu berisi semua fitur yang hanya boleh digunakan Admin.
func adminMenu(
	in *bufio.Reader,
	products handler.ProductHandler,
	customers handler.CustomerHandler,
	orders handler.OrderHandler,
) {
	for {
		fmt.Println("\n=== MENU ADMIN ===")
		fmt.Println("1. Tambah Customer")
		fmt.Println("2. Input Pesanan Customer")
		fmt.Println("3. Lihat Stok Menu")
		fmt.Println("4. Riwayat/Detail Transaksi")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		switch readChoice(in) {
		case 1:
			if err := customers.Create(in); err != nil {
				fmt.Println("Gagal:", err)
			} else {
				fmt.Println("Customer berhasil ditambahkan.")
			}
			pause(in)

		case 2:
			if err := orders.Buy(in); err != nil {
				fmt.Println("Gagal:", err)
			}
			pause(in)

		case 3:
			if err := products.List(); err != nil {
				fmt.Println("Gagal:", err)
			}
			pause(in)

		case 4:
			if err := orders.History(); err != nil {
				fmt.Println("Gagal:", err)
			}
			pause(in)

		case 0:
			return

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

// kepalaTokoMenu berisi fitur pemantauan dan pengelolaan toko.
func kepalaTokoMenu(
	in *bufio.Reader,
	products handler.ProductHandler,
	customers handler.CustomerHandler,
	orders handler.OrderHandler,
) {
	for {
		fmt.Println("\n=== MENU KEPALA TOKO ===")
		fmt.Println("1. Pantau Stok Menu")
		fmt.Println("2. Tambah Menu")
		fmt.Println("3. Tambah Stok Menu")
		fmt.Println("4. Pendapatan/Transaksi Per Hari")
		fmt.Println("5. Hapus Customer")
		fmt.Println("6. Hapus Menu")
		fmt.Println("0. Kembali")
		fmt.Print("Pilih: ")

		switch readChoice(in) {
		case 1:
			if err := products.List(); err != nil {
				fmt.Println("Gagal:", err)
			}
			pause(in)

		case 2:
			if err := products.Create(in); err != nil {
				fmt.Println("Gagal:", err)
			} else {
				fmt.Println("Menu berhasil ditambahkan.")
			}
			pause(in)

		case 3:
			if err := products.AddStock(in); err != nil {
				fmt.Println("Gagal:", err)
			}
			pause(in)

		case 4:
			if err := orders.DailyReport(in); err != nil {
				fmt.Println("Gagal:", err)
			}
			pause(in)

		case 5:
			deleteCustomer(in, customers)
			pause(in)

		case 6:
			if err := products.DeleteProduct(in); err != nil {
				fmt.Println("Gagal menghapus menu:", err)
			}
			pause(in)

		case 0:
			return

		default:
			fmt.Println("Pilihan tidak tersedia.")
		}
	}
}

// deleteCustomer menghapus customer berdasarkan ID.
func deleteCustomer(in *bufio.Reader, customers handler.CustomerHandler) {
	if err := customers.List(); err != nil {
		fmt.Println("Gagal:", err)
		return
	}

	fmt.Println("0. Kembali")
	fmt.Print("Masukkan ID customer yang ingin dihapus: ")
	id := readChoice(in)
	if id == 0 {
		return
	}

	if err := customers.Delete(id); err != nil {
		fmt.Println("Gagal:", err)
		return
	}

	fmt.Println("Customer berhasil dihapus.")
}

// readChoice membaca pilihan menu angka.
func readChoice(in *bufio.Reader) int {
	value, _ := strconv.Atoi(strings.TrimSpace(readLine(in)))
	return value
}

// readLine membaca satu baris input.
func readLine(in *bufio.Reader) string {
	text, _ := in.ReadString('\n')
	return strings.TrimSpace(text)
}

// pause menahan layar sampai user menekan ENTER.
func pause(in *bufio.Reader) {
	fmt.Println("\nTekan ENTER untuk kembali...")
	_, _ = in.ReadString('\n')
}
