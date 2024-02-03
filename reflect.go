package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name    string
	Address string
	Email   string
}

func readField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("TypeName", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
	fmt.Println("=========================================")
}
func main() {

	sample := Sample{"adib"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)

	fmt.Println(structField.Name)
	adib := Person{"Adib", "Bandung", "adibhauzan48@gmail.com"}

	readField(Sample{"Adib"})
	readField(sample)
	readField(adib)
}
