package controllers

import (
    "github.com/revel/revel"
    "myapp/app/models"
)

type Test struct {
    *revel.Controller
}

func (c Test) Confirm() revel.Result {
    return c.Render()
}
