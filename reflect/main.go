package main

import (
	"Go-playground/reflect/structure"
	"fmt"
	"reflect"
)

type TestInt int

func main() {
	var a TestInt
	var b int
	a = TestInt(3)
	b = 3

	fmt.Println(reflect.TypeOf(a)) // main.TestInt
	fmt.Println(reflect.TypeOf(b)) // int

	fmt.Println(reflect.ValueOf(a)) // 3
	fmt.Println(reflect.ValueOf(b)) // 3

	fmt.Println(reflect.ValueOf(a).Kind()) // int
	fmt.Println(reflect.ValueOf(b).Kind()) // int

	var testInt TestInt
	testInt = reflect.ValueOf(b).Convert(reflect.TypeOf(testInt)).Interface().(TestInt)
	// testInt = b
	// testInt에 b를 직접 할당하면 에러 발생. 이를 위해 reflect.Convert를 사용하여 타입 변환

	fmt.Println(reflect.TypeOf(b))       // int
	fmt.Println(b)                       // 3
	fmt.Println(reflect.TypeOf(testInt)) // main.TestInt
	fmt.Println(testInt)                 // 3

	fmt.Println("--------------------------------------------------")

	t := &structure.TestStruct{TestStructValue: 1}
	rt := reflect.TypeOf(t)
	numMethod := rt.NumMethod()
	fmt.Println("NumMethod", numMethod)

	l := 0
	for l < numMethod {
		m := rt.Method(l)
		fmt.Println("Method:", m.Name, "Type:", m.Type)
		fmt.Println("Func:", m.Func, "Index:", m.Index)
		fmt.Println("PkgPath:", m.PkgPath)

		l += 1
	}

	fmt.Println("--------------------------------------------------")

	m, _ := rt.MethodByName("Increment")
	m2, _ := rt.MethodByName("String")

	m.Func.Call([]reflect.Value{reflect.ValueOf(t)})
	m.Func.Call([]reflect.Value{reflect.ValueOf(t)})
	fmt.Println("call from reflect.TypeOf: ", m2.Func.Call([]reflect.Value{reflect.ValueOf(t)}))
}

/*
reflect kind는 미리 정의된 자료형 중 어떤 것인지 확인하는 용
- https://pkg.go.dev/reflect#Kind

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)
*/
