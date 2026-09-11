package entity

// User menyimpan akun yang dapat login ke aplikasi.
type User struct {
	ID int
	Email string
	Password string
	Role string
}

// Customer menyimpan data pelanggan.
type Customer struct {
	ID int
	Name string
	Email string
	Phone string
}

// Product menyimpan data minuman yang tersedia di menu.
type Product struct {
	ID int
	Name string
	Price float64
	Stock int
}

// PaymentMethod menyimpan metode pembayaran yang tersedia.
type PaymentMethod struct {
	ID int
	Name string
}

// Transaction menyimpan ringkasan transaksi.
type Transaction struct {
	ID int
	CustomerID int
	PaymentMethodID int
	Subtotal float64
	Tax float64
	Discount float64
	Total float64
	PaymentName string
	CustomerName string
	CreatedAt string
}

// TransactionDetail menyimpan rincian produk pada transaksi.
type TransactionDetail struct {
	ProductID int
	ProductName string
	Price float64
	Quantity int
	Subtotal float64
}
