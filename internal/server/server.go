package server

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"simplebank.com/internal/handler"
)


type Server struct {
	handler *handler.Handler
	router *gin.Engine
}

func NewServer(handler *handler.Handler) *Server {
	server := &Server{
		handler: handler,
	}

	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router

	return server
}


func (server *Server) Start(address string) error {
    return server.router.Run(address)
}


type CreateAccountParams struct {
    Owner    string `json:"owner"`
    Balance  int64  `json:"balance"`
    Currency string `json:"currency"`
}

type createAccountRequest struct {
    Owner    string `json:"owner" binding:"required"`
    Currency string `json:"currency" binding:"required,oneof=USD EUR"`
	Balance  int64  `json:"balance"`
}




func errorResponse(err error) gin.H {
    return gin.H{"error": err.Error()}
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	arg := handler.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  req.Balance,
	}

	account, err := server.handler.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, account)
}

func (server *Server) getAccount(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(400, errorResponse(err))
		return
	}

	account, err := server.handler.GetAccount(ctx, handler.GetAccountParams{ID: id})
	if err != nil {
		ctx.JSON(500, errorResponse(err))
		return
	}

	ctx.JSON(200, account)
}