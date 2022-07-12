package handler_products

import (
	"net/http"
	"ppob/helper/claudinary"
	err_conv "ppob/helper/err"
	domain_products "ppob/products/domain"
	"ppob/products/handler/request"
	"ppob/products/handler/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ProductsHandler struct {
	Service    domain_products.Service
	Validation *validator.Validate
}

func NewProductsHandler(service domain_products.Service) ProductsHandler {
	return ProductsHandler{
		Service:    service,
		Validation: validator.New(),
	}
}

// Implementation of insert product
func (ph *ProductsHandler) InsertProduct(ctx echo.Context) error {
	// get Request
	req := request.RequestJSONProduct{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// upload image
	if req.Image == "" {
		req.File = claudinary.GetFile(ctx)
		img, err := claudinary.ImageUploadHelper(req.File, "product")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
				"rescode": http.StatusInternalServerError,
			})
		}
		req.Image = img
	}

	// parameter category id
	category_id, _ := strconv.Atoi(ctx.Param("category_id"))

	// product section
	err := ph.Service.InsertData(category_id, request.ToDomain(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product",
		"rescode": http.StatusCreated,
	})
}

// Implementation get all product data
func (ph *ProductsHandler) GetAllProduct(ctx echo.Context) error {
	var sliceProduct []response.ResponseJSONProduct
	// get products
	res, err := ph.Service.GetProducts()
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	for _, value := range res {
		sliceProduct = append(sliceProduct, response.FromDomainProduct(value))
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all product",
		"rescode": http.StatusOK,
		"result":  sliceProduct,
	})
}

// implementation of delete product and detail product
func (ph *ProductsHandler) DestroyProduct(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get response from usecase destroy
	err := ph.Service.Destroy(id)

	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete product",
		"rescode": http.StatusOK,
	})
}

// implementatation update product
func (ph *ProductsHandler) EditProduct(ctx echo.Context) error {
	req := request.RequestJSONProduct{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// get parameter id from endpoint
	id, _ := strconv.Atoi(ctx.Param("id"))
	// upload image
	req.File = claudinary.GetFile(ctx)
	img, _ := claudinary.ImageUploadHelper(req.File, "product")

	req.Image = img
	if req.Image == "" {
		data, err := ph.Service.GetProduct(id)
		if err != nil {
			return err_conv.Conversion(err, ctx)
		}
		req.Image = data.Image
	}

	// edit product
	err := ph.Service.Edit(id, request.ToDomain(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update product",
		"rescode": http.StatusOK,
	})
}

// implementation of get product by id
func (ph *ProductsHandler) GetProduct(ctx echo.Context) error {
	// get parameter id
	param := ctx.Param("id")

	if param == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	id, _ := strconv.Atoi(param)

	// get product by id
	product, err := ph.Service.GetProduct(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// make temp slice data
	sliceDetail := ph.Service.GetDetails(product.Product_Slug)
	details := []interface{}{}
	for _, value := range sliceDetail {
		details = append(details, response.FromDomainDetail(value))
	}
	if len(details) == 0 {
		details = []interface{}{"data empty"}
	}

	// get category product
	category, err := ph.Service.GetCategory(product.Category_Id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all product",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"product":  response.FromDomainProduct(product),
			"detail":   details,
			"category": response.FromDomainCategory(category),
		},
	})
}

func (ph *ProductsHandler) GetProductByCategory(ctx echo.Context) error {
	// get parameter
	id, _ := strconv.Atoi(ctx.Param("category_id"))
	if ctx.Param("category_id") == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	// var slice for response
	sliceProduct := []interface{}{}

	// get products by category
	res := ph.Service.GetProductByCategory(id)
	if len(res) == 0 {
		sliceProduct = append(sliceProduct, "data empty")
	}
	for _, value := range res {
		sliceProduct = append(sliceProduct, response.FromDomainProduct(value))
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get product by category",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"products": sliceProduct,
		},
	})
}

// implementation of insert detail
func (ph *ProductsHandler) InsertDetail(ctx echo.Context) error {
	// get request
	req := request.DataDetail{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// get parameter
	product_slug := ctx.Param("product_slug")

	// get response after insert data
	err := ph.Service.InsertDetail(product_slug, request.ToDomainDetail(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add product",
		"rescode": http.StatusCreated,
	})
}

// Implementation get details by product slug
func (ph *ProductsHandler) GetDetailsProduct(ctx echo.Context) error {
	var sliceDetail []response.ResponseJSONDetail
	product_slug := ctx.Param("product_slug")
	if product_slug == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	res := ph.Service.GetDetails(product_slug)
	if len(res) == 0 {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data empty",
			"rescode": http.StatusBadRequest,
		})
	}
	for _, value := range res {
		sliceDetail = append(sliceDetail, response.FromDomainDetail(value))
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get details product",
		"rescode": http.StatusOK,
		"result": map[string]interface{}{
			"detail": sliceDetail,
		},
	})
}

// Implementation edit detail product by id
func (ph *ProductsHandler) EditDetail(ctx echo.Context) error {
	// get parameter
	param := ctx.Param("getID")
	if param == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	id, _ := strconv.Atoi(param)

	// get request detail product
	req := request.DataDetail{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// get response after edit detail product
	err := ph.Service.EditDetail(id, request.ToDomainDetail(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success edit detail product",
		"rescode": http.StatusOK,
	})
}

// Implementation destroy/delete detail product
func (ph *ProductsHandler) DestroyDetail(ctx echo.Context) error {
	// get parameter
	id, _ := strconv.Atoi(ctx.Param("getID"))

	// get response after destroy or delete detail product
	err := ph.Service.DestroyDetail(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete detail product",
		"rescode": http.StatusOK,
	})
}

// Implementation Insert Category
func (ph *ProductsHandler) InsertCategory(ctx echo.Context) error {
	// get request category
	req := request.RequestJSONCategory{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// get response after insert category
	err := ph.Service.InsertCategory(request.ToDomainCategory(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add category",
		"rescode": http.StatusCreated,
	})
}

// implementation get category by id
func (ph *ProductsHandler) GetCategoryByID(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := ph.Service.GetCategory(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get category",
		"rescode": http.StatusOK,
		"result":  response.FromDomainCategory(data),
	})
}

// implementation get categories
func (ph *ProductsHandler) GetCategories(ctx echo.Context) error {
	// variable slice response category
	var sliceCat []response.ResponseJSONPCategory

	// get data and response categories
	res, err := ph.Service.GetCategories()
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	for _, value := range res {
		sliceCat = append(sliceCat, response.FromDomainCategory(value))
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get categories",
		"rescode": http.StatusOK,
		"result":  sliceCat,
	})
}

// Implementation Edit Category
func (ph *ProductsHandler) EditCategory(ctx echo.Context) error {
	// get parameter
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get request category
	req := request.RequestJSONCategory{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// get response after edit category
	err := ph.Service.EditCategory(id, request.ToDomainCategory(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update category",
		"rescode": http.StatusOK,
	})
}

// Implementation Destroy Category
func (ph *ProductsHandler) DestroyCategory(ctx echo.Context) error {
	// get parameter
	id, _ := strconv.Atoi(ctx.Param("id"))

	// get response after delete/destroy
	err := ph.Service.DestroyCategory(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete category",
		"rescode": http.StatusOK,
	})
}
