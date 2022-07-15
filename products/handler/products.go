package handler_products

import (
	"net/http"
	"ppob/helper/claudinary"
	err_conv "ppob/helper/err"
	"ppob/products/handler/request"
	"ppob/products/handler/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

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

	// get category
	category, err := ph.Service.GetCategory(category_id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// product section
	err = ph.Service.InsertData(category_id, category, request.ToDomain(req))
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
