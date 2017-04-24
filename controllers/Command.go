package controllers

import (
    "github.com/astaxie/beego"
    "os/exec"
)

type CommandController struct {
    beego.Controller
}

func (c *CommandController) ExecTest() {
    var command string
    c.Ctx.Input.Bind(&command, "command")
    cmd, _ := exec.Command("/bin/sh", "-c", command).Output()
    c.Ctx.Output.Body(cmd)
}