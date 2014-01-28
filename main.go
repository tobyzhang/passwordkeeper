package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tobyzxj/passwordkeeper/routers"
)

func main() {
	beego.Run()
}
