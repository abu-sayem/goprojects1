package handler

import (
	"context"

	"simplebank.com/internal/controller"
	"simplebank.com/internal/repository"
	models "simplebank.com/pkg"
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
	params := repository.CreateAccountParams{
		Owner:    arg.Owner,
		Currency: arg.Currency,
		Balance:  arg.Balance,
	}
	return handler.controller.CreateAccount(ctx, params)
}

