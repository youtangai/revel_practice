package controllers

import "github.com/revel/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) WelcomeToUnderGround() revel.Result {
	return c.Render()
}

func (c App) NoThankYou() revel.Result {
	c.Flash.Success("結構です")
	return c.Redirect(App.WelcomeToUnderGround)
}
