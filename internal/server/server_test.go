package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"simplebank.com/internal/controller"
	mock_repository "simplebank.com/internal/gomock"
	"simplebank.com/internal/handler"
	"simplebank.com/internal/utils"
	models "simplebank.com/pkg"
)



func randomAccount() models.Account {
    return models.Account{
        ID:       int64(utils.RandomInt(1, 1000)),
        Owner:    utils.RandomOwner(),
        Balance:  utils.RandomMoney(),
        Currency: utils.RandomCurrency(),
    }
}


func TestGetAccountAPI(t *testing.T) {
    account := randomAccount()

    ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mock_repository.NewMockStore(ctrl)

	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)


		ctl := controller.NewController(store)

		// create handler
		handler := handler.NewHandler(ctl)
	
		// create server
		server := NewServer(handler)
		recorder := httptest.NewRecorder()

		url := fmt.Sprintf("/accounts/%d", account.ID)
		request, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			t.Error(err)
		}

		server.router.ServeHTTP(recorder, request)
		require.Equal(t, http.StatusOK, recorder.Code)
}