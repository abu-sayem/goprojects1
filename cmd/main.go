package main

import (
	"database/sql"
	"fmt"
	"log"

	"simplebank.com/internal/controller"
	"simplebank.com/internal/handler"
	"simplebank.com/internal/repository"
	"simplebank.com/internal/server"
)



const (
    dbDriver = "postgres"
    dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)


func main() {
	 
	// connect to db
	conn, err := sql.Open(dbDriver, dbSource)
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
	server.Start(":8080")
}