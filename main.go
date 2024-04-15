package main

import (
	"github.com/go-suriya/shop-api-echo/config"
	"github.com/go-suriya/shop-api-echo/databases"
	"github.com/go-suriya/shop-api-echo/server"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
