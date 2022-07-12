package domain_products

type Service interface {
	// Product
	InsertData(category_id int, domain Products) error
	GetProducts() ([]Products, error)
	GetProduct(id int) (Products, error)
	GetProductTransaction(product_slug string) (Products, error)
	GetProductByCategory(id int) []Products
	Edit(id int, domain Products) error
	Destroy(id int) error

	// Detail Product
	InsertDetail(product_slug string, domain Detail_Product) error
	GetDetails(product_slug string) []Detail_Product
	GetDetail(detail_slug string) (Detail_Product, error)
	EditDetail(id int, domain Detail_Product) error
	DestroyDetail(id int) error

	// Category
	InsertCategory(domain Category_Product) error
	GetCategories() ([]Category_Product, error)
	GetCategory(id int) (Category_Product, error)
	EditCategory(id int, domain Category_Product) error
	DestroyCategory(id int) error

	// Admin-dashboard
	CountProducts() int
}

type Repository interface {
	// Product
	Store(domain Products) error
	GetByID(id int) (Products, error)
	GetByCategory(id int) []Products
	GetProductTransaction(detail_slug string) (Products, error)
	GetAll() ([]Products, error)
	Update(id int, domain Products) error
	Delete(id int) error

	// Detail Product
	StoreDetail(product_slug string, domain Detail_Product) error
	GetDetailsByCode(product_slug string) ([]Detail_Product, error)
	GetDetail(detail_slug string) (Detail_Product, error)
	DeleteDetail(id int) error
	DeleteDetails(Detail_Product string) error
	UpdateDetail(id int, domain Detail_Product) error
	UpdateDetails(codeLama string, codeBaru string) error

	// Category
	StoreCategory(domain Category_Product) error
	GetCategories() ([]Category_Product, error)
	GetCategoryById(id int) (Category_Product, error)
	UpdateCategory(id int, domain Category_Product) error
	DeleteCategory(id int) error

	// Admin-Dashboard
	Count() int
}
