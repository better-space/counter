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
	platform := flag.String("platform", "Windows", "")
	src := flag.String("src", ".", "")
	suffix := flag.String("suffix", "go", "")
	ignore := flag.Bool("ignore", true, "")

	flag.Parse()

	global.Args.Platform = *platform
	global.Args.Src = *src
	global.Args.Suffix = *suffix
	global.Args.Ignore = *ignore

	fmt.Println(global.Args.Platform, global.Args.Src)
	lineNum := service.CountFolder(global.Args)
	fmt.Printf("您在当前项目已经写了%d行%s代码", lineNum, global.Args.Suffix)
	time.Sleep(time.Second * 10)
}