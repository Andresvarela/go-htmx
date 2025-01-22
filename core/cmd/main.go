package main

import "go-htmx/infra/primary/controllers"

func main() {
	Init()
}

func Init() {
	var ginControllers = controllers.GetControllerInstance(
		[]controllers.ControllerRunnable{&controllers.UsersController{}},
	)
	ginControllers.Start()
}
