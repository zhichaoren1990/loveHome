package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	_ "loveHome/models"
	_ "loveHome/routers"
	"net/http"
	"strings"
)

/**
	ignoreStaticPath()设置首页
 */

func main() {
	ignoreStaticPath()
	beego.Run()
}

func ignoreStaticPath()  {
	beego.InsertFilter("/",beego.BeforeRouter,TransparenStatic)
	beego.InsertFilter("/*",beego.BeforeRouter,TransparenStatic)
}

func TransparenStatic(ctx *context.Context)  {
	orpath :=ctx.Request.URL.Path
	beego.Debug("request url:",orpath)
	if strings.Index(orpath,"api") >= 0{
		return
	}
	http.ServeFile(ctx.ResponseWriter,ctx.Request,"static/html/"+ctx.Request.URL.Path)
}


