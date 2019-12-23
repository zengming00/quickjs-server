
#include "func.h"
#include "main.h"

JSValue getJS_UNDEFINED()
{
    return JS_UNDEFINED;
}

void addMyPrint(JSContext *ctx)
{
    JSValue global_obj, console, foo;

    foo = JS_NewObject(ctx);
    JS_SetPropertyStr(ctx, foo, "myPrint",
                      JS_NewCFunction(ctx, myPrint, "myPrint", 0));

    global_obj = JS_GetGlobalObject(ctx);
    JS_SetPropertyStr(ctx, global_obj, "foo", foo);
    JS_FreeValue(ctx, global_obj);
}