package api

import (
	"fmt"
	"hexagonalapi/api"
)
type Api interface {
	Post()
}
type Restapi struct {
	heman service.Service
}

func Createobject(h service.Service) *Restapi {
	return &Restapi{heman: h}
}
func (p *Restapi) Post() {
	p.heman.Post()
	fmt.Println("middleware post")
}
