package main

import (
	"context"
	"managment/database/simple_connection"
	"managment/http"
	"managment/personal_records"
)

func main() {
	ctx := context.Background()
	conn, err := simple_connection.CreateConnection(ctx)

	if err != nil {
		panic(err)
	}



	managment := personal_records.NewManagment(conn, ctx)

	httpHandler := http.NewHTTPHandler(managment)

	httpServer := http.NewHTTPServer(httpHandler)

	if err := httpServer.StartServer(); err != nil {
		panic(err)
	}
}
