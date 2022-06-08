package handler_products

import (
	"net/http"
	domain_products "ppob/products/domain"
	"ppob/products/handler/request"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductsHandler struct {
	Service    domain_products.Service
	Validation *validator.Validate
}

func NewProductsService(service domain_products.Service) ProductsHandler {
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

	dataCategory, err := ph.Service.GetCategory(req.Category_Name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
			"rescode": http.StatusNotFound,
		})
	}

	code, _ := uuid.NewRandom()

	newReq := request.NewRequest{
		Code:        code.String(),
		Name:        req.Name,
		Slug:        req.Slug,
		Description: req.Description,
		Category_Id: dataCategory.ID,
	}

	res, err := ph.Service.InsertData(request.ToDomain(newReq))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "bad request",
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success add product",
		"rescode": http.StatusOK,
		"data":    res,
	})
}

func (ph *ProductsHandler) GetAll(ctx echo.Context) error {
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
		"data":    res,
	})
}
