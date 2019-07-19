rem TDM-GCC
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/quickjs.c -o quickjs.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/libregexp.c -o libregexp.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/libunicode.c -o libunicode.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/cutils.c -o cutils.o
gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/quickjs-libc.c -o quickjs-libc.o

gcc -DCONFIG_VERSION=\"0.0.1\" -c quickjs-2019-07-09/qjsc.c -o qjsc.o
gcc -DCONFIG_VERSION=\"0.0.1\" -o qjsc.exe  quickjs.o libregexp.o libunicode.o cutils.o quickjs-libc.o qjsc.o