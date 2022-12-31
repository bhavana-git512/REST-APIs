package main
import (
	"fmt"
	"hexaonalapi/api"
	"hexagonalapi/service"
	"hexagonalapi/database"

	//:\Users\pc\hexagonalapiC:\Users\pc\hexagonalapi
)

func main() {
	ret := api.Createobject(&service.Middleware{})
	ret2:=Getdatabase()
	ret1:=Instance(&database.Model{})
	ret.post()
	//ret1.post()
	//ret2.post()
	fmt.Println(ret1,ret2)
}
