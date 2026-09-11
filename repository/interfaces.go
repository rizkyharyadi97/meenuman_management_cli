package repository

import "meenuman/entity"

// UserRepository menangani data akun dan login.
type UserRepository interface {
	Register(entity.User) error
	Login(string, string) (entity.User, error)
}

// CustomerRepository menangani data customer.
type CustomerRepository interface {
	Create(entity.Customer) error
	GetAll() ([]entity.Customer, error)
	GetByID(int) (entity.Customer, error)
	Delete(int) error
}

// ProductRepository menangani data menu minuman.
type ProductRepository interface {
	Create(entity.Product) error
	GetAll() ([]entity.Product, error)
	GetByID(id int) (entity.Product, error)
	Update(entity.Product) error
	Delete(id int) error
	AddStock(id int, quantity int) error
}

// TransactionRepository menangani transaksi dan laporan penjualan.
type TransactionRepository interface {
	Create(entity.Transaction, []entity.TransactionDetail) error
	GetAll() ([]entity.Transaction, error)
	GetDetails(int) ([]entity.TransactionDetail, error)
	GetByDate(string) ([]entity.Transaction, error)
	GetPaymentMethods() ([]entity.PaymentMethod, error)
}
