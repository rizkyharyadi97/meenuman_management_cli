package handler

import (
	"bufio"
	"fmt"

	"meenuman/entity"
	"meenuman/repository"
	"meenuman/table"
)

// CustomerHandler menangani penambahan, tampilan, dan penghapusan data customer.
type CustomerHandler struct {
	Repo repository.CustomerRepository
}

// Create menambahkan customer baru yang nantinya dapat dipilih saat membuat order.
func (h CustomerHandler) Create(in *bufio.Reader) error {
	var c entity.Customer

	fmt.Println("\n=== TAMBAH CUSTOMER ===")

	fmt.Print("Nama customer: ")
	c.Name = readLine(in)

	fmt.Print("Email customer: ")
	c.Email = readLine(in)

	fmt.Print("No. HP: ")
	c.Phone = readLine(in)

	return h.Repo.Create(c)
}

// List menampilkan seluruh customer.
// Delete menghapus customer berdasarkan ID.
func (h CustomerHandler) Delete(id int) error {
	return h.Repo.Delete(id)
}

// List menampilkan seluruh customer.
func (h CustomerHandler) List() error {
	customers, err := h.Repo.GetAll()
	if err != nil {
		return err
	}

	fmt.Println("\n=== DAFTAR CUSTOMER ===")

	if len(customers) == 0 {
		fmt.Println("Belum ada customer.")
		return nil
	}

	table.PrintCustomers(customers)

	return nil
}
