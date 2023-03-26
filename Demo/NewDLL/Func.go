package main

import (
	public "TestDLL/Public"
	"fmt"
)

func Demo(Value map[int]string) public.TypeDLL {
	fmt.Println(Value, "<DLLLOAD>")
	for _, v := range Value {
		fmt.Println(v)
	}
	public.Tmp = 2
	return public.TypeDLL{
		Is_Array: false,
		Text:     "你好",
	}
}
