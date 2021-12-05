package main

import (
	"github.com/alen/echo-framework/db"
	"github.com/alen/echo-framework/routes"
)

func main() {
	db.Init()
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8000"))
}
