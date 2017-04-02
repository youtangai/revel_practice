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

    var user models.User
    var allUsers []models.User
    db.First(&user)
    db.Find(&allUsers)
    err := models.JsonError{0, ""}
    result := arrangeJsonFormat("success!", allUsers, err)
    return c.RenderJSON(result)
}
