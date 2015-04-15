package main

import "fmt"
/*
#include <stdlib.h>
*/
import  "C"

func main(){
    fmt.Println(int(C.random()))
}
//package main
//// int fortytwo()
//// {
////	    return 42;
//// }
//import "C"
//import "fmt"
//
//func main() {
//f := C.intFunc(C.fortytwo)
//fmt.Println(int(C.bridge_int_func(f)))
//// Output: 42
//}


//package main
//import "fmt"
///*
//#cgo CFLAGS: -I.
//#include<add.h>
//*/
//import "C"
//
//func main() {
//    fmt.Println(C.add(5, 4))
//}
