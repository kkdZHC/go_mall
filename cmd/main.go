package main

import (
	"go_mall/conf"
	"go_mall/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
