package main

import (
	"flag"
	"fmt"
	"lvf/counter/global"
	"lvf/counter/service"
	"time"
)

func main() {
	//os.Args
	global.Args.Platform = flag.String("platform", "Windows", "")
	global.Args.Src = flag.String("src", "./", "")
	global.Args.Suffix = flag.String("suffix", "go", "")
	flag.Parse()
	fmt.Println(*global.Args.Platform, *global.Args.Src)
	lineNum := service.CountFolder(global.Args)
	fmt.Printf("您在当前项目已经写了%d行%s代码", lineNum, *global.Args.Suffix)
	time.Sleep(time.Second * 10)
}
