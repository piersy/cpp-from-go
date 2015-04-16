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

/*
Untill the fix for https://github.com/golang/go/issues/5428 is released we may have to
just put the shard library in the root directory of the project
 the fix will allow use of relative paths in the #cgo LDFLAGS variable via the use of ${SRCDIR}
*/

package main
import "fmt"
/*
#cgo CFLAGS: -I./c -I.
#cgo LDFLAGS: -L. -lMyMain_c
#include <MyMain_c.h>
*/
import "C"
//This worked only after i copied libMyMain.so into /usr/lib ??
func main() {
    fmt.Println(C.pain(C.int(5), nil))
    fmt.Println(5)
}
