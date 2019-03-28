package main

import (
	"fmt"
	"github.com/TonyXMH/goblog/accountservice/service"
)

var (
	appName = "accountservice"
)

func main() {
	fmt.Printf("appName %+v", appName)
	service.StartWebServer("6768")
}
