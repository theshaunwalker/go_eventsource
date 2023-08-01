package util

import (
	"fmt"
	"reflect"
)

// Copied from https://stackoverflow.com/questions/8103617/call-a-struct-and-its-method-by-name-in-go
func CallDynamicFunctionOnStructInstance(object interface{}, functionName string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	//ListMethodsOnStruct(object)

	method := reflect.ValueOf(object).MethodByName(functionName)

	if method.IsValid() == false {
		panic("Method " + functionName + " does not exist on type " + reflect.TypeOf(object).Name())
	}

	method.Call(inputs)
}

func ListMethodsOnStruct(foo interface{}) {
	fooType := reflect.TypeOf(foo)

	fmt.Println("fooType --")
	fmt.Printf("%v\n", fooType)

	for i := 0; i < fooType.NumMethod(); i++ {
		method := fooType.Method(i)
		fmt.Println("Method: " + method.Name)
	}
}

func GetTypeName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
