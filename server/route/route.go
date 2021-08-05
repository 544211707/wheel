package route

import (
	"brown/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func InitRouter(app *iris.Application) {
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})
	baseUrl := "/api"
	mvc.New(app.Party(baseUrl + "/wheel")).Handle(controller.NewWheelController())
	mvc.New(app.Party(baseUrl + "/prizeList")).Handle(controller.NewPrizeListController())


}
