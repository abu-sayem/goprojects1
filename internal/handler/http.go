package handler

import (
	"context"

	"simplebank.com/internal/controller"
	models "simplebank.com/pkg"
	pkg "simplebank.com/pkg/params"
)

// APIHandler
type Handler struct {
	controller *controller.Controller
}

func NewHandler(controller *controller.Controller) *Handler {
	return &Handler{
		controller: controller,
	}
}



type CreateAccountParams struct {
    Owner    string `json:"owner"`
    Balance  int64  `json:"balance"`
    Currency string `json:"currency"`
}

func (handler *Handler) CreateAccount(ctx context.Context, arg CreateAccountParams) (models.Account, error) {
	params := pkg.CreateAccountParams{
		Owner:    arg.Owner,
		Currency: arg.Currency,
		Balance:  arg.Balance,
	}
	return handler.controller.CreateAccount(ctx, params)
}



type GetAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (handler *Handler) GetAccount(ctx context.Context, arg GetAccountParams) (models.Account, error) {
	params := controller.GetAccountParams{
		ID: arg.ID,
	}
	return handler.controller.GetAccount(ctx, params)
}
