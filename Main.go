//package main
//
//import "fmt"
///*
//#include <stdlib.h>
//*/
//import  "C"
//
//func main(){
//    fmt.Println(int(C.random()))
//}

//package main
//import "fmt"
//// int fortytwo()
//// {
////	    return 42;
//// }
//import "C"
//
//func main() {
//fmt.Println(int(C.fortytwo()))
//// Output: 42
//}


package main
import "fmt"
/*
#cgo CFLAGS: -I./c
#cgo LDFLAGS: -L./c -lMyMain_c
#include <MyMain_c.h>
*/
import "C"
//This worked only after i copied libMyMain.so into /usr/lib ??
func main() {
    C.pain(C.int(5), nil)
    fmt.Println(5)
}
