package controller

import (
	"brown/model"
	"brown/service"
	"fmt"
	"log"

	"github.com/kataras/iris"
)

type WheelController struct {
	Ctx     iris.Context
	Service service.WheelService
}

func NewWheelController() *WheelController {
	return &WheelController{Service: service.NewWheelServices()}
}

func (p *WheelController) Post() model.Result {

	var wheelPrize model.WheelPrize

	err := p.Ctx.ReadJSON(&wheelPrize)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	fmt.Println(wheelPrize,"controller")

	result := p.Service.Wheel(wheelPrize)

	return result
}
