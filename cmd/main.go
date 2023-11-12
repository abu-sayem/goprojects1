package main

import (
	"database/sql"
	"fmt"
	"log"

	"simplebank.com/internal/controller"
	"simplebank.com/internal/handler"
	"simplebank.com/internal/repository"
	"simplebank.com/internal/server"
	"simplebank.com/internal/utils"

	_ "github.com/lib/pq"
)



func main() {

    config, err := utils.LoadConfig("../")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }
	 
	// connect to db
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer conn.Close()

	// create store and repository
	store := repository.NewStore(conn)

	// create controller
	ctrl := controller.NewController(store)

	// create handler
	handler := handler.NewHandler(ctrl)

	// create server
	server := server.NewServer(handler)

	// start server
	fmt.Println("Server is running at port 8080")
	server.Start(config.ServerAddress)
}