package model

/*
	type Args struct {
		Platform *string
		Src *string
		Suffix *string
	}
 */
type Args struct {
	Platform string	//运行平台
	Src string			//要解析的项目目录
	Suffix string		//解析的代码文件后缀
	Ignore bool				//是否忽略注释
}
