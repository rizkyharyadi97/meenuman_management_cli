package handler

import (
	"bufio"
	"fmt"
	"strconv"

	"meenuman/entity"
	"meenuman/repository"
)

// ProductHandler menangani menu minuman.
type ProductHandler struct {
	Repo repository.ProductRepository
}

// List menampilkan seluruh minuman, harga, dan stok.
func (h ProductHandler) List() error {
	products, err := h.Repo.GetAll()
	if err != nil {
		return err
	}

	fmt.Println("\n=== STOK MEENUMAN ===")

	for _, p := range products {
		fmt.Printf("%d. %-25s Rp%.0f | Stok: %d\n", p.ID, p.Name, p.Price, p.Stock)
	}

	return nil
}

// Create menambahkan minuman baru ke menu.
func (h ProductHandler) Create(in *bufio.Reader) error {
	var p entity.Product

	fmt.Println("\n=== TAMBAH MENU ===")

	fmt.Print("Nama minuman: ")
	p.Name = readLine(in)

	fmt.Print("Harga: ")
	p.Price = readFloat(in)

	fmt.Print("Stok awal: ")
	p.Stock = readInt(in)

	if p.Price <= 0 || p.Stock < 0 {
		return fmt.Errorf("harga harus lebih dari 0 dan stok tidak boleh negatif")
	}

	return h.Repo.Create(p)
}

// AddStock menambahkan stok pada menu yang dipilih Kepala Toko.
func (h ProductHandler) AddStock(in *bufio.Reader) error {
	fmt.Println("\n=== TAMBAH STOK MENU ===")

	products, err := h.Repo.GetAll()
	if err != nil {
		return err
	}

	if len(products) == 0 {
		fmt.Println("Belum ada menu.")
		return nil
	}

	for _, p := range products {
		fmt.Printf("%d. %-25s Rp%.0f | Stok: %d\n", p.ID, p.Name, p.Price, p.Stock)
	}

	fmt.Println("0. Kembali")
	fmt.Print("Pilih ID menu: ")

	productID := readInt(in)
	if productID == 0 {
		return nil
	}

	fmt.Print("Jumlah stok yang ditambahkan: ")
	quantity := readInt(in)

	if quantity <= 0 {
		return fmt.Errorf("jumlah stok harus lebih dari 0")
	}

	if err := h.Repo.AddStock(productID, quantity); err != nil {
		return err
	}

	fmt.Println("Stok berhasil ditambahkan.")
	return nil
}

// DeleteProduct menghapus menu berdasarkan ID.
func (h ProductHandler) DeleteProduct(in *bufio.Reader) error {
	fmt.Println("\n=== HAPUS MENU ===")

	// menampilkan daftar menu yang tersedia
	products, err := h.Repo.GetAll()
	if err != nil {
		return err
	}

	for _, p := range products {
		fmt.Printf("%d. %-25s Rp%.0f | Stok: %d\n", p.ID, p.Name, p.Price, p.Stock)
	}

	fmt.Println("0. Kembali")
	fmt.Print("Pilih ID menu yang ingin dihapus: ")

	productID := readInt(in)
	if productID == 0 {
		return nil
	}

	err = h.Repo.Delete(productID)
	if err != nil {
		return err
	}

	fmt.Println("Menu berhasil dihapus.")
	return nil
}

// readFloat membaca input angka desimal.
func readFloat(in *bufio.Reader) float64 {
	value, _ := strconv.ParseFloat(readLine(in), 64)
	return value
}
