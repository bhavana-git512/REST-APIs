package service

import (
	"fmt"
	"hexagonalapi/database"
)
type Service interface {
	post()
}
type Middleware struct {
	xman database.Database
}

func Instance(x databse.Database) *Middleware {
	return &Middleware{xman: x}
}
func (p *Middleware) Post() {
	p.xman.Post()
	p.xman=&database.Model{}
	
	fmt.Println("model post")
}
