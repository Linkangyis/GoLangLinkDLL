<h1>GoLangLinkDLL</h1>
<h5>简介：Windows使用go引入用go开发的DLL</h5>

## 开发背景
<h5>在原本的Go标准库中，是不支持之间使用go链接用go开发的DLL这在我们开发中带来了许多不便</h5>
<h5>此"框架/插件"致力于解决以下两个issues的问题</h5>
1.https://github.com/golang/go/issues/19282<br>
2.https://github.com/golang/go/issues/22192

## 食用方法
<h5>可以直接测试/Demo中的实例DLL</h5>

1. 编译DLL(需要GCC以及开启CGO) 请注意public包，包位置自行修改
```shell
go env -w GO111MODULE=auto
go build -buildmode=c-shared -o test.dll ./
```
2. 测试程序
```go
package main

import (
	"TestDLL/LoadDll"
	"TestDLL/Public"
	"fmt"
)

func main() {
	Tmp := make(map[int]string)
	Tmp[0] = "123123"
	Tmp[1] = "你好"
	Open := load.OpenDLL("./test.dll")
	fmt.Println(Open.LoadFunc("Demo")(Tmp))
	fmt.Println(Open.ReadConfig())
	fmt.Println(public.Tmp)
}

```
3 运行
```shell
go run test.go

map[0:123123 1:你好] <DLLLOAD>
你好
123123
{false 你好}
map[0:Demo]
0    //这里的0未修改，证明DLL是无法修改共享包中的全局变量 这点和Linux中的Plugin不一样
```
