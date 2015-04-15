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

package main
import "fmt"
// int fortytwo()
// {
//	    return 42;
// }
import "C"

func main() {
fmt.Println(int(C.fortytwo()))
// Output: 42
}


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
