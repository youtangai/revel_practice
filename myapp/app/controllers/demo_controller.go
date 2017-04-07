package controllers

import (
    "github.com/revel/revel"
    "github.com/ChimeraCoder/anaconda"
)

type Demo struct {
    *revel.Controller
}

func (c Demo) Tweet(text string) revel.Result {
    api := GetTwitterApi()
    content := text
    tweet, err := api.PostTweet(content, nil)
    if err != nil {
        panic(err)
    }
    return c.RenderJSON(tweet.Text)
}


func GetTwitterApi() *anaconda.TwitterApi {
  anaconda.SetConsumerKey("7PM3EQ1RJxiDwF7bpXoZbi4BL")
  anaconda.SetConsumerSecret("SbOhBbOgJuKDBejDlOGSWrfwcPRwGHFwkjNOwvJKDlHxdSsK7j")
  api := anaconda.NewTwitterApi("847357434649694208-y4ChfAdwrmB9CwfKd96HL63faqAlOTG", "l7YFOzoPZucXMG8FCEavbTwGazmSsR8LeMMJeNi0Po7ub")
  return api
}
