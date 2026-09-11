package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"meenuman/entity"
)

// MySQLUserRepository adalah implementasi UserRepository menggunakan MySQL.
type MySQLUserRepository struct {
	DB *sql.DB
}

// Register menyimpan akun baru dengan role yang dipilih saat registrasi.
func (r MySQLUserRepository) Register(u entity.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(email, password, role) VALUES (?, ?, ?)",
		u.Email,
		u.Password,
		u.Role,
	)
	return err
}

// Login mencari akun berdasarkan email dan password.
func (r MySQLUserRepository) Login(email, password string) (entity.User, error) {
	var u entity.User

	err := r.DB.QueryRow(
		"SELECT id, email, password, role FROM users WHERE email = ? AND password = ?",
		email,
		password,
	).Scan(&u.ID, &u.Email, &u.Password, &u.Role)

	return u, err
}

// MySQLCustomerRepository adalah implementasi CustomerRepository menggunakan MySQL.
type MySQLCustomerRepository struct {
	DB *sql.DB
}

// Create menambahkan customer baru.
func (r MySQLCustomerRepository) Create(c entity.Customer) error {
	_, err := r.DB.Exec(
		"INSERT INTO customers(name, email, phone) VALUES (?, ?, ?)",
		c.Name,
		c.Email,
		c.Phone,
	)
	return err
}

// GetAll mengambil semua customer.
func (r MySQLCustomerRepository) GetAll() ([]entity.Customer, error) {
	rows, err := r.DB.Query(
		"SELECT id, name, email, phone FROM customers ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []entity.Customer

	for rows.Next() {
		var c entity.Customer

		if err := rows.Scan(&c.ID, &c.Name, &c.Email, &c.Phone);
		err != nil {
			return nil, err
		}

		customers = append(customers, c)
	}

	return customers, rows.Err()
}

// GetByID mengambil satu customer berdasarkan ID.
func (r MySQLCustomerRepository) GetByID(id int) (entity.Customer, error) {
	var c entity.Customer

	err := r.DB.QueryRow(
		"SELECT id, name, email, phone FROM customers WHERE id = ?",
		id,
	).Scan(&c.ID, &c.Name, &c.Email, &c.Phone)

	return c, err
}

// Delete menghapus customer berdasarkan ID.
func (r MySQLCustomerRepository) Delete(id int) error {
	result, err := r.DB.Exec("DELETE FROM customers WHERE id = ?", id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("customer tidak ditemukan")
	}

	return nil
}

// MySQLProductRepository adalah implementasi ProductRepository menggunakan MySQL.
type MySQLProductRepository struct {
	DB *sql.DB
}

// Create menambahkan menu minuman baru.
func (r MySQLProductRepository) Create(p entity.Product) error {
	_, err := r.DB.Exec(
		"INSERT INTO products (name, price, stock) VALUES (?, ?, ?)",
		p.Name,
		p.Price,
		p.Stock,
	)
	return err
}

// Update mengubah data menu minuman.
func (r MySQLProductRepository) Update(p entity.Product) error {
	_, err := r.DB.Exec(
		"UPDATE products SET name = ?, price = ?, stock = ? WHERE id = ?",
		p.Name,
		p.Price,
		p.Stock,
		p.ID,
	)
	return err
}

// Delete menghapus menu minuman berdasarkan ID
func (r MySQLProductRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM products WHERE id = ?", id)

	return err
}

// AddStock menambahkan jumlah stok pada menu yang dipilih.
func (r MySQLProductRepository) AddStock(id int, quantity int) error {
	if quantity <= 0 {
		return errors.New("jumlah stok harus lebih dari 0")
	}

	result, err := r.DB.Exec(
		"UPDATE products SET stock = stock + ? WHERE id = ?",
		quantity,
		id,
	)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return errors.New("menu tidak ditemukan")
	}

	return nil
}

// GetAll mengambil seluruh menu minuman.
func (r MySQLProductRepository) GetAll() ([]entity.Product, error) {
	rows, err := r.DB.Query(
		"SELECT id, name, price, stock FROM products ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		var p entity.Product

		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock); err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, rows.Err()
}

// GetByID mengambil satu produk berdasarkan ID.
func (r MySQLProductRepository) GetByID(id int) (entity.Product, error) {
	var p entity.Product

	err := r.DB.QueryRow(
		"SELECT id, name, price, stock FROM products WHERE id = ?",
		id,
	).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)

	return p, err
}

// MySQLTransactionRepository adalah implementasi TransactionRepository menggunakan MySQL.
type MySQLTransactionRepository struct {
	DB *sql.DB
}

// Create menyimpan transaksi sekaligus mengurangi stok secara atomik.
func (r MySQLTransactionRepository) Create(
	t entity.Transaction,
	details []entity.TransactionDetail,
) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	result, err := tx.Exec(
		`INSERT INTO transactions
		(customer_id, payment_method_id, subtotal, tax, discount, total)
		VALUES (?, ?, ?, ?, ?, ?)`,
		t.CustomerID,
		t.PaymentMethodID,
		t.Subtotal,
		t.Tax,
		t.Discount,
		t.Total,
	)
	if err != nil {
		return err
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	for _, detail := range details {
		result, err := tx.Exec(
			`UPDATE products
			SET stock = stock - ?
			WHERE id = ? AND stock >= ?`,
			detail.Quantity,
			detail.ProductID,
			detail.Quantity,
		)
		if err != nil {
			return err
		}

		affected, _ := result.RowsAffected()
		if affected == 0 {
			return fmt.Errorf("stok produk %s tidak cukup", detail.ProductName)
		}

		_, err = tx.Exec(
			`INSERT INTO transaction_details
			(transaction_id, product_id, quantity, price, subtotal)
			VALUES (?, ?, ?, ?, ?)`,
			transactionID,
			detail.ProductID,
			detail.Quantity,
			detail.Price,
			detail.Subtotal,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetAll mengambil semua transaksi untuk riwayat dan laporan.
func (r MySQLTransactionRepository) GetAll() ([]entity.Transaction, error) {
	rows, err := r.DB.Query(`
		SELECT t.id, t.customer_id, t.payment_method_id,
		       t.subtotal, t.tax, t.discount, t.total,
		       pm.name, c.name, DATE_FORMAT(t.created_at, '%Y-%m-%d %H:%i:%s')
		FROM transactions t
		JOIN payment_methods pm ON pm.id = t.payment_method_id
		JOIN customers c ON c.id = t.customer_id
		ORDER BY t.id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction

	for rows.Next() {
		var t entity.Transaction

		if err := rows.Scan(
			&t.ID,
			&t.CustomerID,
			&t.PaymentMethodID,
			&t.Subtotal,
			&t.Tax,
			&t.Discount,
			&t.Total,
			&t.PaymentName,
			&t.CustomerName,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, rows.Err()
}

// GetDetails mengambil detail produk dari sebuah transaksi.
func (r MySQLTransactionRepository) GetDetails(transactionID int) ([]entity.TransactionDetail, error) {
	rows, err := r.DB.Query(`
		SELECT p.id, p.name, td.price, td.quantity, td.subtotal
		FROM transaction_details td
		JOIN products p ON p.id = td.product_id
		WHERE td.transaction_id = ?
		ORDER BY td.id
	`, transactionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var details []entity.TransactionDetail

	for rows.Next() {
		var d entity.TransactionDetail

		if err := rows.Scan(
			&d.ProductID,
			&d.ProductName,
			&d.Price,
			&d.Quantity,
			&d.Subtotal,
		); err != nil {
			return nil, err
		}

		details = append(details, d)
	}

	return details, rows.Err()
}

// GetByDate mengambil transaksi pada tanggal tertentu dalam format YYYY-MM-DD.
func (r MySQLTransactionRepository) GetByDate(date string) ([]entity.Transaction, error) {
	rows, err := r.DB.Query(`
		SELECT t.id, t.customer_id, t.payment_method_id,
		       t.subtotal, t.tax, t.discount, t.total,
		       pm.name, c.name, DATE_FORMAT(t.created_at, '%Y-%m-%d %H:%i:%s')
		FROM transactions t
		JOIN payment_methods pm ON pm.id = t.payment_method_id
		JOIN customers c ON c.id = t.customer_id
		WHERE DATE(t.created_at) = ?
		ORDER BY t.id DESC
	`, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []entity.Transaction

	for rows.Next() {
		var t entity.Transaction

		if err := rows.Scan(
			&t.ID,
			&t.CustomerID,
			&t.PaymentMethodID,
			&t.Subtotal,
			&t.Tax,
			&t.Discount,
			&t.Total,
			&t.PaymentName,
			&t.CustomerName,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}

	return transactions, rows.Err()
}

// GetPaymentMethods mengambil seluruh metode pembayaran.
func (r MySQLTransactionRepository) GetPaymentMethods() ([]entity.PaymentMethod, error) {
	rows, err := r.DB.Query(
		"SELECT id, name FROM payment_methods ORDER BY id",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []entity.PaymentMethod

	for rows.Next() {
		var method entity.PaymentMethod

		if err := rows.Scan(&method.ID, &method.Name); err != nil {
			return nil, err
		}

		methods = append(methods, method)
	}

	return methods, rows.Err()
}
