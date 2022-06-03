package domain_products

type Service interface {
	// Product
	InsertData(domain Products) (Products, error)
	GetProducts() ([]Products, error)
	GetProduct(code string) (Products, error)
	Edit(code string, domain Products) (Products, error)
	Delete(code string) error
	GetCategory(name string) (Category_Product, error)
	// Category
	InsertCategory(domain Category_Product) (Category_Product, error)
}

type Repository interface {
	// Product
	Store(domain Products) (string, error)
	GetByCode(code string) (Products, error)
	GetAll() ([]Products, error)
	Update(code string, domain Products) (Products, error)
	Delete(code string) error
	// Category
	StoreCategory(domain Category_Product) (Category_Product, error)
	GetCategories() ([]Category_Product, error)
	GetCategoryByName(name string) (Category_Product, error)
}
