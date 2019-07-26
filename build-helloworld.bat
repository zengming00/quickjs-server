rem TDM-GCC
qjsc.exe -o helloworld.c  helloworld.js
gcc -DCONFIG_VERSION=\"0.0.1\" -c helloworld.c -o helloworld.o
gcc -DCONFIG_VERSION=\"0.0.1\" -o helloworld.exe helloworld.o quickjs.o libregexp.o libunicode.o cutils.o quickjs-libc.o

