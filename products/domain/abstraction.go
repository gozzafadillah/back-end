package domain_products

type Service interface {
	// Product
	InsertData(domain Products) (Products, error)
	GetProducts() ([]Products, error)
	GetProduct(id int) (Products, error)
	Edit(id int, domain Products) (Products, error)
	Destroy(id int) error

	// Detail Product
	InsertDetail(domain Detail_Product) (Detail_Product, error)
	GetDetails(id int) ([]Detail_Product, error)
	DestroyDetail(id int) error

	// Category
	InsertCategory(domain Category_Product) (Category_Product, error)
	GetCategories() ([]Category_Product, error)
	GetCategory(name string) (Category_Product, error)
	DestroyCategory(id int) error
}

type Repository interface {
	// Product
	Store(domain Products) (Products, error)
	GetByID(id int) (Products, error)
	GetAll() ([]Products, error)
	Update(id int, domain Products) (Products, error)
	Delete(id int) error

	// Detail Product
	StoreDetail(domain Detail_Product) (Detail_Product, error)
	GetDetaislByID(id int) ([]Detail_Product, error)
	DeleteDetail(id int) error

	// Category
	StoreCategory(domain Category_Product) (Category_Product, error)
	GetCategories() ([]Category_Product, error)
	GetCategoryByName(name string) (Category_Product, error)
	DeleteCategory(id int) error
}
