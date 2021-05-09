package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/swagger"
	_ "github.com/li-zeyuan/frame/packed"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{})
}
