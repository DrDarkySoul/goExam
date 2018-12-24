package main

import (
	"fmt"
	"reflect"
)

func changeValue(i interface{}) {
	if reflect.ValueOf(i).Kind() != reflect.Ptr{
		fmt.Println("It is not a pointer!")
	} else {
		switch reflect.TypeOf(i).String() {
		case "*int":
			reflect.ValueOf(i).Elem().SetInt(69)
		case "*float64":
			reflect.ValueOf(i).Elem().SetFloat(96.666)
		case "*string":
			reflect.ValueOf(i).Elem().SetString("World")
		default:
			fmt.Println("Unknown type")
		}
	}
}

func main() {
	a := 10
	b := 1.0
	c := "Hello"
	e := true

	changeValue(a)
	changeValue(&a)
	changeValue(&b)
	changeValue(&c)
	changeValue(&e)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
