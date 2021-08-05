package controller

import (
	"brown/model"
	"brown/service"
	"fmt"
	"log"

	"github.com/kataras/iris"
)

type PrizeListController struct {
	Ctx     iris.Context
	Service service.PrizeListService
}

func NewPrizeListController() *PrizeListController {
	return &PrizeListController{Service: service.NewPrizeListServices()}
}

func (i *PrizeListController) Post() model.Result {

	var prizeList model.PrizeContent

	err:= i.Ctx.ReadJSON(&prizeList)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	fmt.Println(prizeList,"controller")

	result := i.Service.PrizeList(prizeList)
	return result
}
