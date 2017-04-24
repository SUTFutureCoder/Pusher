package controllers

import (
    "github.com/astaxie/beego"
    "github.com/google/go-github/github"
)

type OauthController struct {
    beego.Controller
}

func (c *OauthController) Oauth() *github.Client{
    //触发
    client := github.NewClient(nil)
    return client
}

func (c *OauthController) OauthEcho() {
    //接受回调

}


