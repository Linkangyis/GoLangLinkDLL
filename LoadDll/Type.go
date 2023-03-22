package load

import "C"
import (
	"TestDLL/Public"
	"unsafe"
)

func CtsrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(C.CString(s)))
}
func FuncRes(ptr uintptr) public.TypeDLL {
	Tmp := *(*public.Type)(unsafe.Pointer(ptr))
	return public.TypeDLL{
		Is_Array: Tmp.Is_Array,
		Text:     string(Tmp.Text),
	}
}
func uintptrToString(ptr uintptr) string {
	return *(*string)(unsafe.Pointer(ptr))
}
func uintptrToMapI_S(ptr uintptr) map[int][]byte {
	return *(*map[int][]byte)(unsafe.Pointer(ptr))
}
func StrPtr(s *string) uintptr {
	return uintptr(unsafe.Pointer(s))
}
func uintptrToMapS_I(maps *map[int][]byte) uintptr {
	return uintptr(unsafe.Pointer(maps))
}

func uintptrToInterface(ptr uintptr) interface{} {
	return *(*interface{})(unsafe.Pointer(ptr))
}
