package main

import (
	"flag"
	"lvf/counter/global"
)

func main() {
	//os.Args
	global.Args.Platform = flag.String("platform", "Windows", "")
	global.Args.Src = flag.String("src", "", "")

}
