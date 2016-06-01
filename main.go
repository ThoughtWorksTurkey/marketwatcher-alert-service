package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
)

// HelloController is a controller
type HelloController struct {
	beego.Controller
}

// Get is a receiver method
func (helloReceiver *HelloController) Get() {
	helloReceiver.Ctx.WriteString("hello get")
}

// Post is a receiver method
func (helloReceiver *HelloController) Post() {
	helloReceiver.Ctx.WriteString("hello post")

	var player Player
	json.Unmarshal(helloReceiver.Ctx.Input.RequestBody, &player)
	fmt.Printf("%+v\n", player.Name)
	helloReceiver.Data["json"] = map[string]Player{"player": player}
	helloReceiver.ServeJSON()
}

type Player struct {
	Id   int
	Name string
}

func main() {
	beego.Router("/", &HelloController{})
	beego.Run()
}
