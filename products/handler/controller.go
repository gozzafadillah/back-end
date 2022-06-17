package handler_products

import (
	"fmt"
	"net/http"
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
	req := request.RequestJSON{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// product section
	err := ph.Service.InsertData(request.ToDomain(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add product",
		"rescode": http.StatusOK,
	})
}

// Implementation get all product data
func (ph *ProductsHandler) GetAllProduct(ctx echo.Context) error {
	var sliceProduct []response.ResponseJSONProduct
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
	req := request.RequestJSON{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	id, _ := strconv.Atoi(ctx.Param("id"))
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
	param := ctx.Param("id")
	fmt.Println("param ", param)
	if param == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	id, _ := strconv.Atoi(param)
	product, err := ph.Service.GetProduct(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	sliceDetail := ph.Service.GetDetails(product.Code)

	details := []interface{}{}
	for _, value := range sliceDetail {
		details = append(details, response.FromDomainDetail(value))
	}
	if len(details) == 0 {
		details = []interface{}{"data empty"}
	}

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
	param := ctx.Param("category_id")
	id, _ := strconv.Atoi(param)
	if param == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	sliceProduct := []interface{}{}
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
	req := request.DataDetail{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	codeParam := ctx.Param("code")

	err := ph.Service.InsertDetail(codeParam, request.ToDomainDetail(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add product",
		"rescode": http.StatusOK,
	})
}

func (ph *ProductsHandler) GetDetailsProduct(ctx echo.Context) error {
	var sliceDetail []response.ResponseJSONDetail
	code := ctx.Param("code")
	if code == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}
	res := ph.Service.GetDetails(code)
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

func (ph *ProductsHandler) EditDetail(ctx echo.Context) error {
	param := ctx.Param("getID")
	fmt.Println("parameter", param)
	if param == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "parameter required",
			"rescode": http.StatusBadRequest,
		})
	}

	id, _ := strconv.Atoi(param)

	req := request.DataDetail{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	err := ph.Service.EditDetail(id, request.ToDomainDetail(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success edit detail product",
		"rescode": http.StatusOK,
	})
}

func (ph *ProductsHandler) DestroyDetail(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("getID"))
	err := ph.Service.DestroyDetail(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete detail product",
		"rescode": http.StatusOK,
	})
}

func (ph *ProductsHandler) InsertCategory(ctx echo.Context) error {
	req := request.RequestJSONCategory{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	err := ph.Service.InsertCategory(request.ToDomainCategory(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add category",
		"rescode": http.StatusOK,
	})
}

func (ph *ProductsHandler) GetCategories(ctx echo.Context) error {
	var sliceCat []response.ResponseJSONPCategory
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

func (ph *ProductsHandler) EditCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	req := request.RequestJSONCategory{}
	ctx.Bind(&req)
	if err := ph.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	err := ph.Service.EditCategory(id, request.ToDomainCategory(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update category",
		"rescode": http.StatusOK,
	})
}

func (ph *ProductsHandler) DestroyCategory(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := ph.Service.DestroyCategory(id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete category",
		"rescode": http.StatusOK,
	})
}
