package domain_products

type Service interface {
	// Product
	InsertData(domain Products) error
	GetProducts() ([]Products, error)
	GetProduct(id int) (Products, error)
	GetProductByCategory(id int) []Products
	Edit(id int, domain Products) error
	Destroy(id int) error

	// Detail Product
	InsertDetail(code string, domain Detail_Product) error
	GetDetails(code string) []Detail_Product
	EditDetail(id int, domain Detail_Product) error
	DestroyDetail(id int) error

	// Category
	InsertCategory(domain Category_Product) error
	GetCategories() ([]Category_Product, error)
	GetCategory(id int) (Category_Product, error)
	EditCategory(id int, domain Category_Product) error
	DestroyCategory(id int) error
}

type Repository interface {
	// Product
	Store(domain Products) error
	GetByID(id int) (Products, error)
	GetByCategory(id int) []Products
	GetAll() ([]Products, error)
	Update(id int, domain Products) error
	Delete(id int) error

	// Detail Product
	StoreDetail(code string, domain Detail_Product) error
	GetDetailsByCode(code string) ([]Detail_Product, error)
	DeleteDetail(id int) error
	DeleteDetails(code string) error
	UpdateDetail(id int, domain Detail_Product) error
	UpdateDetails(codeLama string, codeBaru string) error

	// Category
	StoreCategory(domain Category_Product) error
	GetCategories() ([]Category_Product, error)
	GetCategoryById(id int) (Category_Product, error)
	UpdateCategory(id int, domain Category_Product) error
	DeleteCategory(id int) error
}
