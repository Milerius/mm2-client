package main

import (
	"syscall/js"
)

var (
	promise = js.Global().Get("Promise")
	array   = js.Global().Get("Uint8Array")
)

func newPromise() (p js.Value, set, throw func(js.Value)) {
	type resultT struct {
		v  js.Value
		ok bool
	}
	type resolveT [2]js.Value
	resultCh := make(chan resultT)
	resolveCh := make(chan resolveT, 1)
	go func() {
		result := <-resultCh
		resolvers := <-resolveCh
		if result.ok {
			resolve := resolvers[0]
			resolve.Invoke(result.v)
		} else {
			reject := resolvers[1]
			reject.Invoke(result.v)
		}
	}()
	p = promise.New(js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		resolveCh <- resolveT{args[0], args[1]}
		return nil
	}))
	set = func(v js.Value) {
		resultCh <- resultT{v, true}
	}
	throw = func(v js.Value) {
		resultCh <- resultT{v, false}
	}
	return
}

func goPromise(cb func(args []js.Value) (ret js.Value, ok bool)) js.Value {
	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		p, set, reject := newPromise()
		go func() {
			if ret, ok := cb(args); ok {
				set(ret)
			} else {
				reject(ret)
			}
		}()
		return p
	})
	return f.Value
}

/*func await(awaitable js.Value) (ret js.Value, ok bool) {
	if awaitable.Type() != js.TypeObject || awaitable.Get("then").Type() != js.TypeFunction {
		return awaitable, true
	}
	done := make(chan struct{})

	onResolve := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		glg.Info("resolve")
		ret = args[0]
		ok = true
		close(done)
		return nil
	})
	defer onResolve.Release()

	onReject := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		glg.Info("reject")
		ret = args[0]
		ok = false
		close(done)
		return nil
	})
	defer onReject.Release()

	onCatch := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		glg.Info("catch")
		ret = args[0]
		ok = false
		close(done)
		return nil
	})
	defer onCatch.Release()

	awaitable.Call("then", onResolve, onReject).Call("catch", onCatch)
	<-done
	return
}*/

func bytesToValue(b []byte) js.Value {
	v := array.New(js.ValueOf(len(b)))
	js.CopyBytesToJS(v, b)
	return v
}

func valueToBytes(v js.Value) []byte {
	size := v.Length()
	b := make([]byte, size)
	if n := js.CopyBytesToGo(b, v); n != size {
		panic("bad read")
	}
	return b
}
