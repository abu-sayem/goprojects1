package controller

import (
	"context"

	"simplebank.com/internal/repository"
	models "simplebank.com/pkg"
)


type Controller struct {
	repo *repository.Store
}

func NewController(repo *repository.Store) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (controller *Controller) CreateAccount(ctx context.Context, arg repository.CreateAccountParams) (models.Account, error) {
	return controller.repo.CreateAccount(ctx, arg)
}


func (controller *Controller) ListAccounts(ctx context.Context, arg repository.ListAccountsParams) ([]models.Account, error) {
	return controller.repo.ListAccounts(ctx, arg)
}

func (controller *Controller) CreateEntry(ctx context.Context, arg repository.CreateEntryParams) (models.Entry, error) {
	return controller.repo.CreateEntry(ctx, arg)
}
