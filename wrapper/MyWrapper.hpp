#ifndef __MYWRAPPER_HPP
#define __MYWRAPPER_HPP


//This is both c++ and c compatible
#ifdef __cplusplus
extern "C" {
#endif

typedef struct MyClass MyClass;

MyClass* newMyClass();

void MyClass_int_set(MyClass* v, int i);

int MyClass_int_get(MyClass* v);

void deleteMyClass(MyClass* v);

#ifdef __cplusplus
}
#endif
#endif