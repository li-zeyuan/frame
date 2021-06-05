package main

import (
	"github.com/gogf/gf/frame/g"
	_ "frame/boot"
	_ "frame/router"
)

func main() {
	g.Server().Run()
}
