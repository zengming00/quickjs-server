package main

/*
#include "all.h"
#include "main.h"
#include "func.h"
*/
import "C"
import "fmt"

//export myPrint
func myPrint(ctx *C.JSContext, thisVal C.JSValueConst, argc C.int, argv *C.JSValueConst) C.JSValue {
	fmt.Println("go my println()")
	return C.getJS_UNDEFINED()
}
