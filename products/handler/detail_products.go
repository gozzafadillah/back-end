package handler_products

import (
	"net/http"
	err_conv "ppob/helper/err"
	"ppob/products/handler/request"
	"ppob/products/handler/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

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
