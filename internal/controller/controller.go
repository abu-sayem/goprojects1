package controller

import (
	"context"

	"simplebank.com/internal/repository"
	models "simplebank.com/pkg"
	pkg "simplebank.com/pkg/params"
)


type Controller struct {
	repo repository.Store
}

func NewController(repo repository.Store) *Controller {
	return &Controller{
		repo: repo,
	}
}


func (controller *Controller) CreateAccount(ctx context.Context, arg pkg.CreateAccountParams) (models.Account, error) {
	return controller.repo.CreateAccount(ctx, arg)
}


func (controller *Controller) ListAccounts(ctx context.Context, arg pkg.ListAccountsParams) ([]models.Account, error) {
	return controller.repo.ListAccounts(ctx, arg)
}

func (controller *Controller) CreateEntry(ctx context.Context, arg pkg.CreateEntryParams) (models.Entry, error) {
	return controller.repo.CreateEntry(ctx, arg)
}

type GetAccountParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (controller *Controller) GetAccount(ctx context.Context, arg GetAccountParams) (models.Account, error) {
	return controller.repo.GetAccount(ctx, arg.ID)
}
