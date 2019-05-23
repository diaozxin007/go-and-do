package controller

import (
	"github.com/kataras/iris"
	"go-and-do/models"
)


func GetActive(ctx iris.Context) {
	activeId, _ := ctx.Params().GetInt64("id")
	active, _ := models.GetActive(activeId)
	if active != nil {
		ctx.JSON(active)
	}

}

func PostActive(ctx iris.Context)  {
	active := new(models.Active)
	ctx.ReadJSON(&active)
	i, _ := models.AddActive(active.Name, 1)
	if i != 0 {
		ctx.JSON(i)
	}
}
