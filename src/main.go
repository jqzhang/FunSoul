package main

import (
	_ "GoFun/src/config"
	_ "GoFun/src/db"
	"GoFun/src/engine"
)

func main() {
	eng := engine.Init()
	defer eng.Run(":6600")

	// 启动前初始化

}
