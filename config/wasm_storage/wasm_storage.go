package wasm_storage

import "syscall/js"

func Store(key string, val string) {
	js.Global().Get("localStorage").Call("setItem", key, val)
}

func Retrieve(key string) string {
	value := js.Global().Get("localStorage").Call("getItem", key)
	if value.IsNull() {
		return ""
	}
	return value.String()
}

func Remove(key string) {
	js.Global().Get("localStorage").Call("removeItem", key)
}

func Clear() {
	js.Global().Get("localStorage").Call("clear")
}
