package handler_products

import (
	"net/http"
	domain_products "ppob/products/domain"
	"ppob/products/handler/request"
	"ppob/products/handler/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

	// Kepikiran memakai slug
	code, _ := uuid.NewRandom()
	req.Code = code.String()
	// product section
	err := ph.Service.InsertData(request.ToDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	// detail section
	dataDetail := request.DataDetail{
		Code:        code.String(),
		Price:       req.Price,
		Description: req.Description,
		Status:      true,
	}

	err = ph.Service.InsertDetail(request.ToDomainDetail(dataDetail))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add product",
		"rescode": http.StatusOK,
	})
}

func (ph *ProductsHandler) GetAllProduct(ctx echo.Context) error {
	res, err := ph.Service.GetProducts()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
			"rescode": http.StatusInternalServerError,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all product",
		"rescode": http.StatusOK,
		"result":  res,
	})
}

func (ph *ProductsHandler) GetProduct(ctx echo.Context) error {
	param := ctx.Param("id")
	if param == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data required",
			"rescode": http.StatusBadRequest,
		})
	}
	id, _ := strconv.Atoi(param)
	product, err := ph.Service.GetProduct(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data not found",
			"rescode": http.StatusBadRequest,
		})
	}
	sliceDetail, err := ph.Service.GetDetails(product.Code)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data not found",
			"rescode": http.StatusBadRequest,
		})
	}

	details := []interface{}{}
	for _, value := range sliceDetail {
		details = append(details, response.FromDomainDetail(value))
	}

	category, err := ph.Service.GetCategory(product.Category_Id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "data not found",
			"rescode": http.StatusBadRequest,
		})
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
