package table

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"meenuman/entity"
)

// newTable membuat format tabel yang konsisten untuk seluruh aplikasi.
func newTable(headers []string) *tablewriter.Table {
	t := tablewriter.NewWriter(os.Stdout)

	t.Header(headers)

	return t
}

// PrintProducts menampilkan daftar menu minuman, harga, dan stok.
func PrintProducts(products []entity.Product) {
	if len(products) == 0 {
		fmt.Println("Belum ada menu.")
		return
	}

	t := newTable([]string{"ID", "NAMA MINUMAN", "HARGA", "STOK"})

	for _, p := range products {
		t.Append([]string{
			fmt.Sprintf("%d", p.ID),
			p.Name,
			fmt.Sprintf("Rp%.0f", p.Price),
			fmt.Sprintf("%d", p.Stock),
		})
	}

	t.Render()
}

// PrintCustomers menampilkan daftar customer.
func PrintCustomers(customers []entity.Customer) {
	if len(customers) == 0 {
		fmt.Println("Belum ada customer.")
		return
	}

	t := newTable([]string{"ID", "NAMA", "EMAIL", "NO. HP"})

	for _, c := range customers {
		t.Append([]string{
			fmt.Sprintf("%d", c.ID),
			c.Name,
			c.Email,
			c.Phone,
		})
	}

	t.Render()
}

// PrintOrderDetails menampilkan rincian item pada order.
func PrintOrderDetails(details []entity.TransactionDetail) {
	if len(details) == 0 {
		fmt.Println("Tidak ada detail order.")
		return
	}

	t := newTable([]string{"PRODUCT", "HARGA", "QTY", "SUBTOTAL"})

	for _, d := range details {
		t.Append([]string{
			d.ProductName,
			fmt.Sprintf("Rp%.0f", d.Price),
			fmt.Sprintf("%d", d.Quantity),
			fmt.Sprintf("Rp%.0f", d.Subtotal),
		})
	}

	t.Render()
}

// PrintTransactions menampilkan ringkasan seluruh transaksi.
func PrintTransactions(transactions []entity.Transaction) {
	if len(transactions) == 0 {
		fmt.Println("Belum ada transaksi.")
		return
	}

	t := newTable([]string{
		"ID",
		"CUSTOMER",
		"PEMBAYARAN",
		"SUBTOTAL",
		"PPN",
		"DISKON",
		"TOTAL",
		"WAKTU",
	})

	for _, trx := range transactions {
		t.Append([]string{
			fmt.Sprintf("%d", trx.ID),
			trx.CustomerName,
			trx.PaymentName,
			fmt.Sprintf("Rp%.0f", trx.Subtotal),
			fmt.Sprintf("Rp%.0f", trx.Tax),
			fmt.Sprintf("Rp%.0f", trx.Discount),
			fmt.Sprintf("Rp%.0f", trx.Total),
			trx.CreatedAt,
		})
	}

	t.Render()
}

// PrintDailyReport menampilkan transaksi dan total pendapatan harian.
func PrintDailyReport(transactions []entity.Transaction) {
	if len(transactions) == 0 {
		fmt.Println("Tidak ada transaksi pada tanggal tersebut.")
		return
	}

	t := newTable([]string{"ID", "CUSTOMER", "PEMBAYARAN", "TOTAL"})
	var income float64

	for _, trx := range transactions {
		income += trx.Total

		t.Append([]string{
			fmt.Sprintf("%d", trx.ID),
			trx.CustomerName,
			trx.PaymentName,
			fmt.Sprintf("Rp%.0f", trx.Total),
		})
	}

	t.Render()

	fmt.Println()
	fmt.Printf("Jumlah transaksi : %d\n", len(transactions))
	fmt.Printf("Total pendapatan : Rp%.0f\n", income)
}

// PrintPaymentMethods menampilkan pilihan metode pembayaran.
func PrintPaymentMethods(methods []entity.PaymentMethod) {
	if len(methods) == 0 {
		fmt.Println("Belum ada metode pembayaran.")
		return
	}

	t := newTable([]string{"ID", "METODE PEMBAYARAN"})

	for _, method := range methods {
		t.Append([]string{
			fmt.Sprintf("%d", method.ID),
			method.Name,
		})
	}

	t.Render()
}
