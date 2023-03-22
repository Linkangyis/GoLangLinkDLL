package main

import "C"
import (
	public "./GoLangLinkDLL/Public"
	"fmt"
	"unsafe"
)

//export INITGODLL
func INITGODLL(Map uintptr, LoadFuncName uintptr) uintptr {
	INITS()
	TmpMap := PtrMap(Map)
	FuncName := uintptrToString(LoadFuncName)
	if _, ok := ConfigMap[FuncName]; ok {
		TmpRes := ConfigMap[FuncName](TmpMap)
		fmt.Scanln(TmpRes)
		return NewRes(&TmpRes)
	}

	NewResT := public.TypeDLL{
		Is_Array: false,
		Text:     "<ERROR>",
	}
	fmt.Scanln(NewRes(&NewResT)) //进入GC不让其消失
	panic("ERROR NOT FUNC")
	return NewRes(&NewResT)
}

//export INITREADALL
func INITREADALL() uintptr {
	INITS()
	TmpMap := make(map[int][]byte)
	for k, _ := range ConfigMap {
		TmpMap[len(TmpMap)] = []byte(k)
	}
	fmt.Scanln(MapPtr(&TmpMap))
	return MapPtr(&TmpMap)
}

func NewRes(Structs *public.TypeDLL) uintptr {

	Structss := public.Type{
		Is_Array: Structs.Is_Array,
		Text:     []byte(Structs.Text),
	}

	return uintptr(unsafe.Pointer(&Structss))
}

func uintptrToString(ptr uintptr) string {
	return *(*string)(unsafe.Pointer(ptr))
}

func StrPtr(s *string) uintptr {
	return uintptr(unsafe.Pointer(s))
}

func PtrMap(ptr uintptr) map[int]string {
	res := make(map[int]string)
	tmp := *(*map[int][]byte)(unsafe.Pointer(ptr))
	for k, v := range tmp {
		res[k] = string(v)
	}
	return res
}

func MapPtr(m *map[int][]byte) uintptr {
	return uintptr(unsafe.Pointer(m))
}

func main() {

}
