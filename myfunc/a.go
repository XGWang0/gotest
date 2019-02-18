package myfunc

import "fmt"

func init() {
	fmt.Println("This is init func of a.go")

}

func Afunc(str string) {
    fmt.Println("a.go is package mufunc.")
    fmt.Println(str)
}
