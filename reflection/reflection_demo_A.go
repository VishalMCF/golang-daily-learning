package reflection

import (
	"errors"
	"fmt"
	"reflect"
)

type Student struct {
	name string
	age  int
}

type LevelOneGod string

type Coordinate struct {
	pointX float32
	pointY float32
	pointZ float32
}

func Print() {
	var value LevelOneGod
	value = "hello mega god"
	fmt.Println(value)
}

func TestReflection() {
	student := Student{
		"Vishal",
		27,
	}

	coordinate := Coordinate{
		17.75,
		77.59,
		952.0,
	}

	helloToGod := LevelOneGod("meowG")

	helloToGodPtr := &helloToGod

	fmt.Println(getNames(student))
	fmt.Println(getNames(coordinate))
	fmt.Println(getNames(helloToGod))
	fmt.Println(getNames(helloToGodPtr))
}

func getNames(value interface{}) (name string, err error) {
	typeOf := reflect.TypeOf(value)
	switch typeOf.Kind() {
	case reflect.Struct:
		name := typeOf.Name()
		if name != "" {
			return name, err
		}
		return "", errors.New("notFoundName: name was not found")
	case reflect.String:
		if typeOf.Name() == typeOf.Kind().String() {
			return value.(string), err
		}
		return typeOf.Name(), err
	case reflect.Ptr:
		name := typeOf.Elem().Name()
		if name != "" {
			return name, err
		}
		return "", errors.New("notFoundName : name was not found")
	default:
		if typeOf.Name() == typeOf.Kind().String() {
			return "", errors.New("Not found any name")
		}
		return typeOf.Name(), err
	}
}
