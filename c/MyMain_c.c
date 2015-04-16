#include "MyWrapper.hpp"
#include <stdio.h>

int main(int argc, char* argv[]) {
        struct MyClass* c = newMyClass();
        MyClass_int_set(c, 3);
        printf("%i\n", MyClass_int_get(c));
        deleteMyClass(c);
}