package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"go-and-do/controller"
	"go-and-do/database"
)

var addr = iris.Addr("localhost:8080")

func main() {

	database.Init()
	app := iris.New()
	app.Get("/actives/{id:int64}", controller.GetActive)
	app.Post("/active",controller.PostActive)
	app.Run(addr)
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello World"})
}
