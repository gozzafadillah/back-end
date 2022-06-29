package handler_transaction

import (
	"net/http"
	"ppob/app/middlewares"
	err_conv "ppob/helper/err"
	helper_xendit "ppob/helper/xendit"
	domain_products "ppob/products/domain"
	domain_transaction "ppob/transaction/domain"
	"ppob/transaction/handler/request"
	"ppob/transaction/handler/response"
	domain_users "ppob/users/domain"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TransactionHandler struct {
	TransactionUsecase domain_transaction.Service
	ProductUsecase     domain_products.Service
	UserUsecase        domain_users.Service
	Validation         *validator.Validate
}

func NewTransactionHandler(tc domain_transaction.Service, pc domain_products.Service, uc domain_users.Service) TransactionHandler {
	return TransactionHandler{
		TransactionUsecase: tc,
		ProductUsecase:     pc,
		UserUsecase:        uc,
		Validation:         validator.New(),
	}
}

func (th *TransactionHandler) Checkout(ctx echo.Context) error {
	req := request.Detail_Transaction{}
	ctx.Bind(&req)
	if err := th.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}
	// parameter
	detail_slug := ctx.Param("detail_slug")

	//  claim session from jwt
	claim := middlewares.GetUser(ctx)

	// get user
	user, err := th.UserUsecase.GetUserPhone(claim.Phone)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// get Detail product
	detailproduct, err := th.ProductUsecase.GetDetail(detail_slug)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// get product
	product, err := th.ProductUsecase.GetProductTransaction(detailproduct.Product_Slug)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// get Category product
	category, err := th.ProductUsecase.GetCategory(product.Category_Id)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	req.Fee = 2000
	req.Price = detailproduct.Price
	// make detail transaction / checkout
	detail, err := th.TransactionUsecase.AddDetailTransaction(detail_slug, request.TodomainDetail(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// make invoice
	invoice, err := helper_xendit.Xendit_Invoice(detail, detailproduct, user, category.Name)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// Transaction
	err = th.TransactionUsecase.AddTransaction(invoice, detail)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success add detail transaction",
		"rescode": http.StatusCreated,
		"result": map[string]interface{}{
			"checkout": response.FromDomainCheckout(detail),
			"product":  response.FromDomainProduct(product),
			"category": response.FromDomainCatProduct(category),
		},
		"xendit_invoice": invoice,
	})
}

func (th *TransactionHandler) Callback_Invoice(ctx echo.Context) error {
	req := request.Callback_Invoice{}
	ctx.Bind(&req)

	dataCallback, err := helper_xendit.GetCallback(request.ToDomainCallback(req), ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed get callback " + err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success get callback",
		"rescode": http.StatusCreated,
		"result":  dataCallback,
	})
}
