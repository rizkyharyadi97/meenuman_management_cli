package handler

import (
	"bufio"
	"errors"
	"fmt"

	"meenuman/entity"
	"meenuman/repository"
	"meenuman/table"
)

// OrderHandler menangani input order, perhitungan, pembayaran, dan riwayat transaksi.
type OrderHandler struct {
	Products repository.ProductRepository
	Customers repository.CustomerRepository
	Transactions repository.TransactionRepository
}

// Membuat orderan untuk customer yang dipilih oleh Admin.
func (h OrderHandler) Buy(in *bufio.Reader) error {
	customers, err := h.Customers.GetAll()
	if err != nil {
		return err
	}

	if len(customers) == 0 {
		return errors.New("belum ada customer, silakan tambahkan customer terlebih dahulu")
	}

	fmt.Println("\n=== PILIH CUSTOMER ===")
	table.PrintCustomers(customers)
	fmt.Println("0. Kembali")
	fmt.Print("Pilih customer: ")

	customerID := readInt(in)
	if customerID == 0 {
		return nil
	}

	if _, err := h.Customers.GetByID(customerID); err != nil {
		return errors.New("customer tidak ditemukan")
	}

	// cart menyimpan ID produk dan jumlah yang dipilih customer.
	cart := make(map[int]int)

	for {
		products, err := h.Products.GetAll()
		if err != nil {
			return err
		}

		fmt.Println("\n=== MENU MEENUMAN ===")
		table.PrintProducts(products)

		fmt.Println("0. Selesai memilih")
		fmt.Println("-1. Kembali")
		fmt.Print("Pilih ID produk: ")

		productID := readInt(in)
		if productID == -1 {
			return nil
		}
		if productID == 0 {
			break
		}

		product, err := h.Products.GetByID(productID)
		if err != nil {
			fmt.Println("Produk tidak ditemukan.")
			continue
		}

		fmt.Printf("Jumlah %s: ", product.Name)
		quantity := readInt(in)

		if quantity < 0 {
			fmt.Println("Jumlah harus lebih dari 0.")
			continue
		}

		if cart[productID]+quantity > product.Stock {
			fmt.Println("Stok tidak cukup.")
			continue
		}

		cart[productID] += quantity
		fmt.Println("Produk berhasil ditambahkan ke order.")
	}

	if len(cart) == 0 {
		return errors.New("order kosong")
	}

	// Susun detail order sekaligus hitung subtotal.
	details := make([]entity.TransactionDetail, 0, len(cart))
	var subtotal float64

	for productID, quantity := range cart {
		product, err := h.Products.GetByID(productID)
		if err != nil {
			return err
		}

		lineSubtotal := product.Price * float64(quantity)
		subtotal += lineSubtotal

		details = append(details, entity.TransactionDetail{
			ProductID: product.ID,
			ProductName: product.Name,
			Price: product.Price,
			Quantity: quantity,
			Subtotal: lineSubtotal,
		})
	}

	// PPN 11% jika subtotal lebih dari Rp100.000.
	var tax float64
	if subtotal > 100000 {
		tax = subtotal * 0.11
	}

	// Diskon 10% jika subtotal lebih dari Rp200.000.
	var discount float64
	if subtotal > 200000 {
		discount = subtotal * 0.10
	}

	// Total akhir = subtotal + PPN - diskon.
	total := subtotal + tax - discount

	fmt.Println("\n=== DETAIL ORDER ===")
	table.PrintOrderDetails(details)

	fmt.Println("----------------------------------------")
	fmt.Printf("Subtotal : Rp%.0f\n", subtotal)
	fmt.Printf("PPN 11%%  : Rp%.0f\n", tax)
	fmt.Printf("Diskon 10%%: Rp%.0f\n", discount)
	fmt.Printf("TOTAL     : Rp%.0f\n", total)

	fmt.Println("\n1. Lanjut pembayaran")
	fmt.Println("0. Kembali")
	fmt.Print("Pilih: ")

	if readInt(in) == 0 {
		return nil
	}

	// Tampilkan pilihan pembayaran.
	methods, err := h.Transactions.GetPaymentMethods()
	if err != nil {
		return err
	}

	fmt.Println("\n=== PEMBAYARAN ===")
	table.PrintPaymentMethods(methods)
	fmt.Println("0. Kembali")
	fmt.Print("Pilih pembayaran: ")

	paymentID := readInt(in)
	if paymentID == 0 {
		return nil
	}

	validPayment := false
	for _, method := range methods {
		if method.ID == paymentID {
			validPayment = true
			break
		}
	}
	if !validPayment {
		return errors.New("metode pembayaran tidak valid")
	}

	transaction := entity.Transaction{
		CustomerID: customerID,
		PaymentMethodID: paymentID,
		Subtotal: subtotal,
		Tax: tax,
		Discount: discount,
		Total: total,
	}

	if err := h.Transactions.Create(transaction, details); err != nil {
		return err
	}

	fmt.Println("\nTransaksi berhasil disimpan.")
	fmt.Printf("Total pembayaran: Rp%.0f\n", total)

	return nil
}

// History menampilkan seluruh riwayat transaksi beserta detailnya.
func (h OrderHandler) History() error {
	transactions, err := h.Transactions.GetAll()
	if err != nil {
		return err
	}

	if len(transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		return nil
	}

	fmt.Println("\n=== RIWAYAT TRANSAKSI ===")
	table.PrintTransactions(transactions)

	for _, t := range transactions {
		fmt.Printf("\nDetail transaksi #%d:\n", t.ID)
		details, err := h.Transactions.GetDetails(t.ID)
		if err != nil {
			return err
		}

		table.PrintOrderDetails(details)
	}

	return nil
}

// DailyReport menampilkan transaksi dan total pendapatan pada tanggal tertentu.
func (h OrderHandler) DailyReport(in *bufio.Reader) error {
	fmt.Println("\n=== LAPORAN PENDAPATAN HARIAN ===")
	fmt.Print("Tanggal (YYYY-MM-DD): ")
	date := readLine(in)

	transactions, err := h.Transactions.GetByDate(date)
	if err != nil {
		return err
	}

	if len(transactions) == 0 {
		fmt.Println("Tidak ada transaksi pada tanggal tersebut.")
		return nil
	}

	table.PrintDailyReport(transactions)

	return nil
}
