package main

import (
	"go_todolist/conf"
	"go_todolist/routes"
)

func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
