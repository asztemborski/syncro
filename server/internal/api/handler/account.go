package handler

import (
	"github.com/asztemborski/syncro/internal/app"
	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	app app.App
}

func NewAccountHandler(app *app.App) *AccountHandler {
	return &AccountHandler{app: *app}
}

func (h *AccountHandler) Register(e *echo.Echo) {
	e.POST("/account", h.createAccount)
}

type createAccountRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required,min=5,max=20"`
	Password        string `json:"password" validate:"required,min=6,max=70"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

func (h *AccountHandler) createAccount(c echo.Context) error {
	req, err := retrieveAndValidateRequestBody[createAccountRequest](c)
	if err != nil {
		return err
	}

	return h.app.AccountService().CreateAccount(c.Request().Context(), app.CreateAccountPayload{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
	})
}
