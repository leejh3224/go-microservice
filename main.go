package main

import (
	"db"
	"server"
	"service"
)

func initBoltClient() {
	service.DbClient = &db.BoltClient{}
	service.DbClient.Init()
	service.DbClient.Seed()
}

func main() {
	initBoltClient()
	server.StartServer("8080")
}
