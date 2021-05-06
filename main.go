package main

import (
	"fmt"

	"github.com/thaoeu/pavment_management_system/conf"
	"github.com/thaoeu/pavment_management_system/router"
)

func main() {
	conf.DefaultInit()

	r := router.RoutersInit()
	fmt.Println("开始运行")
	_ = r.Run(":8081")

}
