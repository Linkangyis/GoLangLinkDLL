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
