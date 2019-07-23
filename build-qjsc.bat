rem TDM-GCC
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-src/quickjs.c -o quickjs.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-src/libregexp.c -o libregexp.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-src/libunicode.c -o libunicode.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-src/cutils.c -o cutils.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-src/quickjs-libc.c -o quickjs-libc.o

gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-src/qjsc.c -o qjsc.o
gcc -DCONFIG_VERSION=\"0.0.1\" -o qjsc.exe  quickjs.o libregexp.o libunicode.o cutils.o quickjs-libc.o qjsc.o