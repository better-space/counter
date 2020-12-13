package service

import (
	"bytes"
	"io/ioutil"
	"log"
	"lvf/counter/global"
	"lvf/counter/model"
	"os"
	"strings"
)

/*实现功能描述
	1.可区分操作系统
	2.可设置文件扩展名类型，比如.go .py .java .cpp等
	3.剔除文件中的空行，仅仅计算实际的代码行数
	4.···(可设置剔除注释)
 */

var lineNum int

func CountFolder(args *model.Args) int {
	files,err := ioutil.ReadDir(*args.Src)

	if err != nil {
		log.Println("read folder src error.")
	}
	for _,file := range files {
		//curUrl := src+"/"+file.Name()
		SearchFile(*args.Src, file)
	}
	return lineNum
}

func SearchFile(baseUrl string, info os.FileInfo)  {
	curUrl := baseUrl +"/"+info.Name()
	switch info.IsDir() {
	case true:
		files,err := ioutil.ReadDir(curUrl)
		if err != nil {
			log.Printf("read folder %s error.", curUrl)
		}
		for _,file := range files {
			SearchFile(curUrl, file)
		}
	default:
		//buf := bufio.NewReader(info)
		name := strings.Split(curUrl,".")
		if bytes.Compare([]byte(name[len(name)-1]),[]byte(*global.Args.Suffix)) != 0 {
			break
		}
		byt,err := ioutil.ReadFile(curUrl)
		if err != nil {
			log.Printf("read file %s error.", curUrl)
		}
		content := string(byt)
		file := strings.Split(content, "\n")
		for _,v := range file {
			v = strings.TrimSpace(v)
			if len(v) != 0 {
				lineNum++
			}
		}
	}
}
