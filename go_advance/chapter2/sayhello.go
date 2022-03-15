// sayhello.go
package main

//#include <sayhello.h>
import "C"

func main() {
	C.SayHello(C.CString("Hello, World\n"))
}
