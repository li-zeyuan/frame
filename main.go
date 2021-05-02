package main

import (
	_ "frame/boot"
	_ "frame/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
