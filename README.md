# counter
go语言计算代码行数工具
**实现功能描述**  
    1.可区分操作系统  
	2.可设置文件扩展名类型，比如.go .py .java .cpp等  
	3.剔除文件中的空行，仅仅计算实际的代码行数  
	4.可设置剔除注释  
	5.(......)  
**运行命令举例**  
例如统计当前文件夹下所有的go语言代码命令:  
```
go run main.go --platform=windows --src=./ --suffix=go --ignore=true
```
