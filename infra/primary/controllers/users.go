package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-htmx/infra/secondary/persistence/dao"
	"text/template"
)

type UsersController struct {
}

type ControllerRunnable interface {
	RunController(r *gin.Engine)
}

type GinController struct {
	controllers []ControllerRunnable
}

func GetControllerInstance(c []ControllerRunnable) *GinController {
	return &GinController{c}
}
func (c *GinController) Start() {
	router := gin.Default()
	for _, o := range c.controllers {
		o.RunController(router)
	}
	if err := router.Run(":8080"); err != nil {
		return
	}
	fmt.Println("Server is running on port 8080")
}

func (c *UsersController) RunController(r *gin.Engine) {

	userDAO, err := dao.NewUserDAO("andres")
	if err != nil {
		panic(err)
	}

	r.GET("/", func(c *gin.Context) {
		users, err := userDAO.FindAllUsers()
		if err != nil {
			fmt.Printf("Error while fetching users: %v", err)
		}

		tmpl := template.Must(template.ParseFiles("infra/templates/users.html"))
		if err = tmpl.Execute(c.Writer, gin.H{"users": users}); err != nil {
			panic(err)
		}
	})

	r.POST("/users", func(c *gin.Context) {
		tmpl := template.Must(template.ParseFiles("infra/templates/user_row.html"))

		user, _ := userDAO.CreateUser(dao.User{
			Name:  c.PostForm("name"),
			Email: c.PostForm("email"),
		})

		if err := tmpl.Execute(c.Writer, user); err != nil {
			panic(err)
		}
	})

	r.DELETE("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := userDAO.DeleteUser(name)
		if err != nil {
			fmt.Printf("Error while deleting user: %v", err)
		}
		fmt.Println("Delete user with name:", name)
	})
}
