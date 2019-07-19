rem TDM-GCC
qjsc.exe -c -o repl.c -m quickjs-2019-07-09/repl.js
gcc -DCONFIG_VERSION=\"0.0.1\" -c repl.c -o repl.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/qjs.c -o qjs.o
gcc -DCONFIG_VERSION=\"0.0.1\" -o qjs.exe repl.o quickjs.o libregexp.o libunicode.o cutils.o quickjs-libc.o qjs.o

