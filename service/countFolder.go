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
	files, err := ioutil.ReadDir(args.Src)

	if err != nil {
		log.Println("read folder src error.")
	}
	for _, file := range files {
		//curUrl := src+"/"+file.Name()
		SearchFile(args, file)
	}
	return lineNum
}

func SearchFile(args *model.Args, info os.FileInfo) {
	//过滤git文件
	if info.Name()[:1] == "." {
		return
	}
	curUrl := args.Src + "/" + info.Name()
	switch info.IsDir() {
	case true:
		files, err := ioutil.ReadDir(curUrl)
		if err != nil {
			log.Printf("read folder %s error.", curUrl)
		}
		args.Src = curUrl
		for _, file := range files {
			SearchFile(args, file)
		}
		l := len(info.Name())
		args.Src = args.Src[:len(args.Src)-l]
	default:
		//buf := bufio.NewReader(info)
		name := strings.Split(curUrl, ".")

		//判断后缀
		if bytes.Compare([]byte(name[len(name)-1]), []byte(global.Args.Suffix)) != 0 {
			break
		}
		//读文件
		byt, err := ioutil.ReadFile(curUrl)
		if err != nil {
			log.Printf("read file %s error.", curUrl)
		}
		//内容分割
		content := string(byt)
		file := strings.Split(content, "\n")
		sgn := false
		//遍历分割的每一行code
		for _, v := range file {
			v = strings.TrimSpace(v)

			//处理是否计算注释内容
			if len(v) >= 2 {
				switch args.Ignore {
				case true:
					if v[:2] == "//" {
						continue
					}
					if v[:2] == "/*" {
						sgn = true
					}
					if v[:2] == "*/" || v[len(v)-2:] == "*/" { //行首或行尾的注释结束符
						sgn = false
						continue
					}
				default:
					break
				}
			}
			//非空行计数加一
			if len(v) != 0 && !sgn {
				lineNum++
			}
		}
	}
}
