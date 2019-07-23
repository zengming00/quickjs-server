package main

/*
#include <string.h>
#include "quickjs.h"
#include "libregexp.h"
#include "libunicode.h"
#include "cutils.h"
#include "quickjs-libc.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	rt := C.JS_NewRuntime()
	defer C.JS_FreeRuntime(rt)
	ctx := C.JS_NewContext(rt)
	defer C.JS_FreeContext(ctx)

	C.js_std_add_helpers(ctx, C.int(0), (**C.char)(nil))

	buf := C.CString("console.log('helloworld for go & quickjs');")
	defer C.free(unsafe.Pointer(buf))

	buf_len := C.strlen(buf)

	filename := C.CString("<filename>")
	defer C.free(unsafe.Pointer(filename))

	eval_flags := C.JS_EVAL_TYPE_GLOBAL
	eval_flags |= C.JS_EVAL_FLAG_SHEBANG
	val := C.JS_Eval(ctx, buf, buf_len, filename, C.int(eval_flags))
	if C.JS_IsException(val) != 0 {
		C.js_std_dump_error(ctx)
	}
	C.JS_FreeValue(ctx, val)
	fmt.Println("val:", val)

}
