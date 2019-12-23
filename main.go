package main

/*
#cgo CFLAGS: -D_GNU_SOURCE

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "func.h"
#include "main.h"


*/
import "C"

import (
	"io/ioutil"
	"os"
	"unsafe"
)

func main() {

	rt := C.JS_NewRuntime()
	defer C.JS_FreeRuntime(rt)
	ctx := C.JS_NewContext(rt)
	defer C.JS_FreeContext(ctx)

	C.js_std_add_helpers(ctx, C.int(0), (**C.char)(nil))
	C.addMyPrint(ctx)

	osString := C.CString("os")
	defer C.free(unsafe.Pointer(osString))
	C.js_init_module_os(ctx, osString)

	file := os.Args[1]
	filename := C.CString(file)
	defer C.free(unsafe.Pointer(filename))

	bts, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	buf := C.CString(string(bts))
	defer C.free(unsafe.Pointer(buf))

	val := C.JS_Eval(ctx, buf, C.strlen(buf), filename, C.int(C.JS_EVAL_TYPE_GLOBAL|C.JS_EVAL_TYPE_MODULE))
	defer C.JS_FreeValue(ctx, val)
	if C.JS_IsException(val) != 0 {
		C.js_std_dump_error(ctx)
	}
	C.js_std_loop(ctx)
}
