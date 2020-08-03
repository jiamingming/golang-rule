package commonutils

import (
	"bytes"
	"unsafe"
)

/*
* 拼接字符串
*
 */
func ContactString(item ...string) string {

	var item_buffer bytes.Buffer
	for _, v := range item {
		item_buffer.WriteString(v)
	}
	return item_buffer.String()
}

/*
* string转换为[]byte
 */
func Str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

/*
* []byte转换为string
 */
func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
