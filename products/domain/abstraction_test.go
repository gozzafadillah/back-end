package domain_products_test

import (
	"errors"
	"os"
	domain_products "ppob/products/domain"
	productMockRepo "ppob/products/domain/mocks"
	service_products "ppob/products/service"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productService domain_products.Service
	productDomain  domain_products.Products
	DetailDomain   domain_products.Detail_Product
	CategoryDomain domain_products.Category_Product
	productRepo    productMockRepo.Repository
)

func TestMain(m *testing.M) {
	productService = service_products.NewProductsService(&productRepo)
	productDomain = domain_products.Products{
		ID:           1,
		Product_Slug: "paket-xl",
		Name:         "Paket XL",
		Image:        "xl.jpg",
		Category_Id:  1,
		Status:       true,
	}
	DetailDomain = domain_products.Detail_Product{
		ID:           1,
		Product_Slug: "paket-xl",
		Name:         "Paket XL 20rb",
		Detail_Slug:  "paket-xl-20rb",
		Price:        20000,
		Status:       true,
	}
	CategoryDomain = domain_products.Category_Product{
		ID:            1,
		Name:          "Paket Data",
		Category_Slug: "paket-data",
		Image:         "font-paket-data",
		Status:        true,
	}
	os.Exit(m.Run())
}

func TestGetProduct(t *testing.T) {
	t.Run("success get product by Id", func(t *testing.T) {
		productRepo.On("GetByID", mock.Anything).Return(productDomain, nil).Once()
		res, err := productService.GetProduct(productDomain.ID)
		assert.NoError(t, err)
		assert.Equal(t, productDomain.ID, res.ID)
	})
	t.Run("failed get product by Id", func(t *testing.T) {
		productRepo.On("GetByID", mock.Anything).Return(domain_products.Products{}, errors.New("failed get product")).Once()
		res, err := productService.GetProduct(productDomain.ID)
		assert.Error(t, err)
		assert.Equal(t, domain_products.Products{}, res)
	})
}

func TestGetProductByCategory(t *testing.T) {
	t.Run("success get product by category id", func(t *testing.T) {
		productRepo.On("GetByCategory", mock.Anything).Return([]domain_products.Products{productDomain}).Once()
		res := productService.GetProductByCategory(productDomain.ID)
		assert.Equal(t, productDomain.ID, res[0].ID)
	})
}

func TestGetProducts(t *testing.T) {
	t.Run("success get products", func(t *testing.T) {
		productRepo.On("GetAll").Return([]domain_products.Products{productDomain}, nil).Once()
		res, err := productService.GetProducts()
		assert.NoError(t, err)
		assert.Equal(t, productDomain.ID, res[0].ID)
	})
	t.Run("failed get products", func(t *testing.T) {
		productRepo.On("GetAll").Return([]domain_products.Products{}, errors.New("database empty")).Once()
		res, err := productService.GetProducts()
		assert.Error(t, err)
		assert.Equal(t, []domain_products.Products{}, res)
	})
}

func TestGetProductTransaction(t *testing.T) {
	t.Run("success get product for transaction", func(t *testing.T) {
		productRepo.On("GetProductTransaction", mock.Anything).Return(productDomain, nil).Once()
		res, err := productService.GetProductTransaction(productDomain.Product_Slug)
		assert.NoError(t, err)
		assert.Equal(t, productDomain, res)
	})
	t.Run("failed get product for transaction", func(t *testing.T) {
		productRepo.On("GetProductTransaction", mock.Anything).Return(domain_products.Products{}, errors.New("failed get product for transaction")).Once()
		res, err := productService.GetProductTransaction(productDomain.Product_Slug)
		assert.Error(t, err)
		assert.Equal(t, domain_products.Products{}, res)
	})
}

func TestInsertData(t *testing.T) {
	t.Run("success insert data product", func(t *testing.T) {
		productRepo.On("Store", mock.Anything).Return(nil).Once()
		err := productService.InsertData(1, CategoryDomain, productDomain)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed insert data product", func(t *testing.T) {
		productRepo.On("Store", mock.Anything).Return(errors.New("failed store data")).Once()
		err := productService.InsertData(1, CategoryDomain, productDomain)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("internal server error"))
	})
}

func TestDestroy(t *testing.T) {
	t.Run("success destroy product", func(t *testing.T) {
		productRepo.On("GetByID", mock.Anything).Return(productDomain, nil).Once()
		productRepo.On("Delete", productDomain.ID).Return(nil).Once()
		productRepo.On("DeleteDetails", productDomain.Product_Slug).Return(nil).Once()

		err := productService.Destroy(productDomain.ID)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)

	})
	t.Run("failed destroy product", func(t *testing.T) {
		productRepo.On("GetByID", mock.Anything).Return(productDomain, errors.New("bad request")).Once()
		productRepo.On("Delete", productDomain.ID).Return(nil).Once()
		productRepo.On("DeleteDetails", productDomain.Product_Slug).Return(nil).Once()

		err := productService.Destroy(productDomain.ID)
		assert.Error(t, err)
		assert.Equal(t, errors.New("bad request"), err)
	})
}

func TestEdit(t *testing.T) {
	t.Run("success edit", func(t *testing.T) {
		productRepo.On("GetByID", mock.Anything).Return(productDomain, nil).Once()
		productRepo.On("Update", mock.Anything, mock.Anything).Return(nil).Once()
		productRepo.On("UpdateDetails", mock.Anything, mock.Anything).Return(nil).Once()

		err := productService.Edit(productDomain.ID, productDomain)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed edit", func(t *testing.T) {
		productRepo.On("GetByID", mock.Anything).Return(domain_products.Products{}, errors.New("failed get by id")).Once()
		productRepo.On("Update", mock.Anything, mock.Anything).Return(errors.New("failed update data product")).Once()
		productRepo.On("UpdateDetails", mock.Anything, mock.Anything).Return(errors.New("failed update detail product")).Once()

		err := productService.Edit(productDomain.ID, productDomain)
		assert.Error(t, err)
		assert.Equal(t, err, errors.New("bad request"))
	})
}

func TestGetDetails(t *testing.T) {
	t.Run("success get details product", func(t *testing.T) {
		productRepo.On("GetDetailsByCode", mock.Anything).Return([]domain_products.Detail_Product{DetailDomain}, nil).Once()
		res := productService.GetDetails(productDomain.Product_Slug)
		assert.Equal(t, []domain_products.Detail_Product{DetailDomain}, res)
	})
	t.Run("failed get details product", func(t *testing.T) {
		productRepo.On("GetDetailsByCode", mock.Anything).Return([]domain_products.Detail_Product{}, errors.New("error")).Once()
		res := productService.GetDetails(productDomain.Product_Slug)
		assert.Equal(t, []domain_products.Detail_Product{}, res)
	})
}
func TestGetDetail(t *testing.T) {
	t.Run("success get detail product", func(t *testing.T) {
		productRepo.On("GetDetail", mock.Anything).Return(DetailDomain, nil).Once()
		res, err := productService.GetDetail(DetailDomain.Detail_Slug)
		assert.NoError(t, err)
		assert.Equal(t, DetailDomain, res)
	})
	t.Run("failed get detail product", func(t *testing.T) {
		productRepo.On("GetDetail", mock.Anything).Return(domain_products.Detail_Product{}, errors.New("error")).Once()
		res, err := productService.GetDetail(DetailDomain.Detail_Slug)
		assert.Error(t, err)
		assert.Equal(t, domain_products.Detail_Product{}, res)
	})
}

func TestInsertDetail(t *testing.T) {
	t.Run("success insert detail product", func(t *testing.T) {
		productRepo.On("StoreDetail", mock.Anything, mock.Anything).Return(nil).Once()
		err := productService.InsertDetail(productDomain.Product_Slug, DetailDomain)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed insert detail product", func(t *testing.T) {
		productRepo.On("StoreDetail", mock.Anything, mock.Anything).Return(errors.New("internal status error")).Once()
		err := productService.InsertDetail(productDomain.Product_Slug, DetailDomain)
		assert.Error(t, err)
		assert.Equal(t, errors.New("internal server error"), err)
	})
}

func TestEditDetail(t *testing.T) {
	t.Run("success edit detail", func(t *testing.T) {
		productRepo.On("UpdateDetail", mock.Anything, mock.Anything).Return(nil).Once()
		err := productService.EditDetail(DetailDomain.ID, DetailDomain)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed edit detail", func(t *testing.T) {
		productRepo.On("UpdateDetail", mock.Anything, mock.Anything).Return(errors.New("update failed")).Once()
		err := productService.EditDetail(DetailDomain.ID, DetailDomain)
		assert.Error(t, err)
		assert.Equal(t, errors.New("update failed"), err)
	})
}

func TestDestroyDetail(t *testing.T) {
	t.Run("success destroy detail", func(t *testing.T) {
		productRepo.On("DeleteDetail", mock.Anything).Return(nil).Once()
		err := productService.DestroyDetail(DetailDomain.ID)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed destroy detail", func(t *testing.T) {
		productRepo.On("DeleteDetail", mock.Anything).Return(errors.New("delete failed")).Once()
		err := productService.DestroyDetail(DetailDomain.ID)
		assert.Error(t, err)
		assert.Equal(t, errors.New("delete failed"), err)
	})
}

func TestGetCategory(t *testing.T) {
	t.Run("success get category by id", func(t *testing.T) {
		productRepo.On("GetCategoryById", mock.Anything).Return(CategoryDomain, nil).Once()
		res, err := productService.GetCategory(CategoryDomain.ID)
		assert.NoError(t, err)
		assert.Equal(t, CategoryDomain, res)
	})
	t.Run("failed get category by id", func(t *testing.T) {
		productRepo.On("GetCategoryById", mock.Anything).Return(domain_products.Category_Product{}, errors.New("bad request")).Once()
		res, err := productService.GetCategory(CategoryDomain.ID)
		assert.Error(t, err)
		assert.Equal(t, domain_products.Category_Product{}, res)
	})
}

func TestGetCategories(t *testing.T) {
	t.Run("success get categories", func(t *testing.T) {
		productRepo.On("GetCategories").Return([]domain_products.Category_Product{CategoryDomain}, nil).Once()
		data, err := productService.GetCategories()
		assert.NoError(t, err)
		assert.Equal(t, CategoryDomain.ID, data[0].ID)
	})
	t.Run("failed get categories", func(t *testing.T) {
		productRepo.On("GetCategories").Return([]domain_products.Category_Product{}, errors.New("internal status error")).Once()
		data, err := productService.GetCategories()
		assert.Error(t, err)
		assert.Equal(t, []domain_products.Category_Product{}, data)
	})
}

func TestInsertCategory(t *testing.T) {
	t.Run("success insert product category", func(t *testing.T) {
		productRepo.On("StoreCategory", mock.Anything).Return(nil).Once()
		err := productService.InsertCategory(CategoryDomain)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed insert product category", func(t *testing.T) {
		productRepo.On("StoreCategory", mock.Anything).Return(errors.New("internal server error")).Once()
		err := productService.InsertCategory(CategoryDomain)
		assert.Error(t, err)
		assert.Equal(t, errors.New("internal server error"), err)
	})
}

func TestEditCategory(t *testing.T) {
	t.Run("success edit category", func(t *testing.T) {
		productRepo.On("UpdateCategory", mock.Anything, mock.Anything).Return(nil).Once()
		err := productService.EditCategory(CategoryDomain.ID, CategoryDomain)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed edit category", func(t *testing.T) {
		productRepo.On("UpdateCategory", mock.Anything, mock.Anything).Return(errors.New("update failed")).Once()
		err := productService.EditCategory(CategoryDomain.ID, CategoryDomain)
		assert.Error(t, err)
		assert.Equal(t, errors.New("update failed"), err)
	})
}

func TestDestroyCategory(t *testing.T) {
	t.Run("success destroy category", func(t *testing.T) {
		productRepo.On("DeleteCategory", mock.Anything).Return(nil).Once()
		err := productService.DestroyCategory(CategoryDomain.ID)
		assert.NoError(t, err)
		assert.Equal(t, nil, err)
	})
	t.Run("failed destroy category", func(t *testing.T) {
		productRepo.On("DeleteCategory", mock.Anything).Return(errors.New("delete failed")).Once()
		err := productService.DestroyCategory(CategoryDomain.ID)
		assert.Error(t, err)
		assert.Equal(t, errors.New("delete failed"), err)
	})
}

func TestCount(t *testing.T) {
	t.Run("count product", func(t *testing.T) {
		productRepo.On("Count").Return(1).Once()
		data := productService.CountProducts()
		assert.Equal(t, 1, data)
	})
}

// go test ./products/domain/abstraction_test.go -coverpkg=./products/service/...
// ok      command-line-arguments  0.409s  coverage: 93.3% of statements in ./products/service/...
