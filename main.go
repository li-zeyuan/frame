package main

import (
	"github.com/gogf/gf/frame/g"
	_ "github.com/li-zeyuan/frame/boot"
	_ "github.com/li-zeyuan/frame/router"
)

func main() {
	g.Server().Run()
}
