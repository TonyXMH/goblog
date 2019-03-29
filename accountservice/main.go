package main

import (
	"fmt"
	"github.com/TonyXMH/goblog/accountservice/service"
	"github.com/TonyXMH/goblog/dbclient"
)

var (
	appName = "accountservice"
)

func main() {
	fmt.Printf("appName %+v", appName)
	initBoltClient()
	service.StartWebServer("6768")
}

func initBoltClient()  {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDB()
	service.DBClient.Seed()
}