package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
)

type UserController struct {
    BaseController
}

func (c UserController) Index() revel.Result {
    db := connectDB()

    var allUsers []models.User
    db.Find(&allUsers)

    err := models.JsonError{0, ""}
    result := convertJsonFormat("success!", allUsers, err)

    return c.RenderJSON(result)
}

func (c UserController) New() revel.Result {
    return c.Render()
}

func (c UserController) Create(name string) revel.Result {
    db := connectDB()

    user := models.User{}
    user.Name = name

    db.NewRecord(user)

    db.Create(&user)

    return c.Redirect(UserController.Index)
}

func (c UserController) CreateUserTable() revel.Result {
  db := connectDB()
  db.CreateTable(models.User{})
  return c.RenderJSON("seccess")
}
