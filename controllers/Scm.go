package controllers

import (
    "github.com/astaxie/beego"
    "encoding/json"
)

type ScmController struct {
    beego.Controller
}

type Test struct {
    Website string `json:"website"`
    Email   string `json:"email"`
}

func (c *ScmController) BindScm() {
    //获取输入
    var url string
    c.Ctx.Input.Bind(&url, "url")
    if url == "" {
        c.Abort("400")
    }

    //输出样例
    var mystruct Test
    mystruct.Website = "beego.me"
    mystruct.Email = "astaxie@gmail.com"
    jsonData, _ := json.Marshal(mystruct)
    c.Data["json"] = string(jsonData)
    c.ServeJSON()
}
