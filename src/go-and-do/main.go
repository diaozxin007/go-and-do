package main
import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/kataras/iris"
	"go-and-do/controller"
	"go-and-do/database"
	"go-and-do/models"
	"log"
)

var addr = iris.Addr("localhost:8080")

func main() {

	database.Init()
	//
	//err := database.Engine.CreateTables(new(models.Active))
	//if err != nil{
	//	log.Fatalf("error %v\n",err)
	//}

	err2 := database.Engine.Sync2(new(models.Active))
	if err2 != nil{
		log.Fatalf("error %v\n",err2)
	}

	app := iris.New()
	app.Get("/active", controller.GetActive)
	app.Post("/active",controller.PostActive)
	app.Run(addr)
}

func index(ctx iris.Context) {
	ctx.JSON(iris.Map{"message": "Hello World"})
}
