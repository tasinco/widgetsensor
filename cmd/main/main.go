package main

import (
	"widgetsensor/internal/service"
)

func main() {
	serviceInst := service.Service{}
	serviceInst.Run()
}
