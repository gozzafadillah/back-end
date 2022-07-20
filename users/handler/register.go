package handler_users

import (
	"net/http"
	"ppob/helper/claudinary"
	"ppob/helper/encryption"
	err_conv "ppob/helper/err"
	otp_generator "ppob/helper/otp"
	regexPhone "ppob/helper/phone"
	"ppob/users/handler/request"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// implementation register users
func (uh *UsersHandler) Register(ctx echo.Context) error {
	req := request.RequestJSONUser{}
	ctx.Bind(&req)
	if err := uh.Validation.Struct(req); err != nil {
		stringerr := []string{}
		for _, errval := range err.(validator.ValidationErrors) {
			stringerr = append(stringerr, errval.Field()+" is not "+errval.Tag())
		}
		return ctx.JSON(http.StatusBadRequest, stringerr)
	}

	// check phone
	statusPhone := regexPhone.CheckPhone(req.Phone)
	if !statusPhone {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "phone not valid",
			"rescode": http.StatusOK,
		})
	}

	// change phone to international code
	req.Phone = regexPhone.GenerateNewPhone(req.Phone)

	// upload image
	req.File = claudinary.GetFile(ctx)
	img, _ := claudinary.ImageUploadHelper(req.File, "users")

	req.Image = img
	if req.Image == "" {
		req.Image = "https://res.cloudinary.com/dt91kxctr/image/upload/v1655825545/go-bayeue/users/download_o1yrxx.png"
	}

	// enkripsi password
	encrypt, err := encryption.HashPassword(req.Password)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}
	req.Password = encrypt

	// store request data to usecase layer
	data, err := uh.usecase.Register(request.ToDomainUser(req))
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	// make otp
	otpCode := otp_generator.OtpGenerator()

	err = uh.usecase.AddUserVerif(otpCode, req.Email, req.Name)
	if err != nil {
		return err_conv.Conversion(err, ctx)
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success register",
		"rescode": http.StatusCreated,
		"data": map[string]interface{}{
			"token": data,
		},
	})
}
