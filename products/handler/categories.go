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

	// upload image
	if req.Image == "" {
		req.File = claudinary.GetFile(ctx)
		img, err := claudinary.ImageUploadHelper(req.File, "category")
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
				"rescode": http.StatusInternalServerError,
			})
		}
		req.Image = img
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

	// upload image
	req.File = claudinary.GetFile(ctx)
	img, _ := claudinary.ImageUploadHelper(req.File, "category")

	req.Image = img
	if req.Image == "" {
		data, err := ph.Service.GetCategory(id)
		if err != nil {
			return err_conv.Conversion(err, ctx)
		}
		req.Image = data.Image
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
