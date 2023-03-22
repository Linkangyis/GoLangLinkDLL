package load

import (
	"TestDLL/Public"
	"syscall"
)

const CDLLFILE = "./LoadC.dll"

var MessageBoxW *syscall.LazyProc
var MessageBoxConfig *syscall.LazyProc

type DLLFILE string

func init() {
	var TmpMap = make(map[int][]byte)
	TmpMap[0] = []byte("Test")
	read := syscall.NewLazyDLL(CDLLFILE)
	MessageBoxW = read.NewProc("INITCDLL")
	MessageBoxConfig = read.NewProc("READCONFIG")
}

func OpenDLL(File string) DLLFILE {
	return DLLFILE(File)
}

func (GoDLLfile DLLFILE) LoadFunc(funcName string) func(map[int]string) public.TypeDLL {
	return func(Value map[int]string) public.TypeDLL {
		ValueByte := make(map[int][]byte)
		for k, v := range Value {
			ValueByte[k] = []byte(v)
		}
		TmpfuncName := &funcName
		Res, _, _ := MessageBoxW.Call(CtsrPtr(string(GoDLLfile)), CtsrPtr("INITGODLL"), StrPtr(TmpfuncName), uintptrToMapS_I(&ValueByte))
		return FuncRes(Res)
	}
}

func (GoDLLfile DLLFILE) ReadConfig() map[int]string {
	Res, _, _ := MessageBoxConfig.Call(CtsrPtr(string(GoDLLfile)), CtsrPtr("INITREADALL"))
	TmpMap := uintptrToMapI_S(Res)
	ResMap := make(map[int]string)
	for k, v := range TmpMap {
		ResMap[k] = string(v)
	}
	return ResMap
}
