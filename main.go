/**
 * Created by GoLand.
 * User: xzghua@gmail.com
 * Date: 2018-12-02
 * Time: 01:34
 */
package main

import (
	"fmt"

	"github.com/thaoeu/pavment_management_system/conf"
	"github.com/thaoeu/pavment_management_system/router"
)

func main() {
	conf.DefaultInit()
	//csrf

	r := router.RoutersInit()
	fmt.Println("开始运行")
	_ = r.Run(":8081")

}
