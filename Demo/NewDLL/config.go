package main

import (
	public "TestDLL/Public"
)

var ConfigMap = make(map[string]func(map[int]string) public.TypeDLL)

func INITS() {
	/* v 请在这里配置函数*/
	ConfigMap["Demo"] = Demo
	/* ^ 请在这里配置函数*/
}
