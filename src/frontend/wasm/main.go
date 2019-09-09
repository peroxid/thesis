package main

import (
	"crypto/sha256"
	"fmt"
	"syscall/js"
)

var hasher Hasher

func consoleLog(this js.Value, in []js.Value) interface{} {
	println(js.Global().
		Get("document").
		Call("getElementById", "test").
		Get("value").
		String())
	return this
}

func registerCallbacks() {
	js.Global().Set("consoleLog", js.FuncOf(consoleLog))
	js.Global().Set("progressiveHash", js.FuncOf(progressiveHash))
	js.Global().Set("startHash", js.FuncOf(startHash))
	js.Global().Set("getHash", js.FuncOf(getHash))
}

type Hasher struct {
	input  chan []byte
	result chan []byte
}

func (h *Hasher) sha256Hash() {
	hash := sha256.New()
	for {
		data, more := <-h.input
		if more {
			fmt.Println("received data")
			hash.Write(data)
			fmt.Println("hashed data")
		} else {
			fmt.Println("received close")
			break
		}
	}
	h.result <- hash.Sum(nil)
	println(h.result)
	return
}

func progressiveHash(this js.Value, in []js.Value) interface{} {
	array := in[0]
	buf := make([]byte, array.Get("length").Int())
	n := js.CopyBytesToGo(buf, array)
	fmt.Printf("Copied %d bytes\n", n)
	hasher.input <- buf
	return this
}

func startHash(this js.Value, in []js.Value) interface{} {
	hasher.result = make(chan []byte)
	hasher.input = make(chan []byte)
	go hasher.sha256Hash()
	return this
}

func getHash(this js.Value, in []js.Value) interface{} {
	close(hasher.input)
	hash := <-hasher.result
	hashStr := fmt.Sprintf("%x", hash)
	fmt.Printf("Hash: %s\n", hashStr)

	return js.ValueOf(hashStr)
}

func waitForever() {
	c := make(chan struct{}, 0)
	<-c
}

func main() {
	fmt.Println("WASM Go Initialized")
	registerCallbacks()
	waitForever()
}
