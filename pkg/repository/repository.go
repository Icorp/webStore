package repository

// Communicate with database
type Authorization interface {
}

type ProductList interface {
}

type ProductItem interface {
}

type Repository struct {
	Authorization
	ProductItem
	ProductList
}

func NewRepository() *Repository {
	return &Repository{}
}
