package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

type Vector []float64

func sum(array []float64) (result float64) {
	for _, v := range array {
		result += v
	}
	return
}

func (v *Vector) Add(u interface{}) (newVector Vector, err error) {
	switch u.(type) {
	case float32, float64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]+reflect.ValueOf(u).Float())
		}
		return
	case int, int8, int16, int32, int64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]+float64(reflect.ValueOf(u).Int()))
		}
		return
	case uint, uint8, uint16, uint32, uint64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]+float64(reflect.ValueOf(u).Uint()))
		}
		return
	case Vector:
		s := reflect.ValueOf(u)
		if len(*v) != s.Len() {
			return *v, errors.New("Error: Vectors are not of the same size")
		}
		for i := range *v {
			newVector = append(newVector, (*v)[i]+s.Index(i).Float())
		}
		return
	default:
		return *v, fmt.Errorf("Error: Can't use %v to manipulate a Vector", reflect.TypeOf(u).String())
	}
}

func (v *Vector) Sub(u interface{}) (newVector Vector, err error) {
	switch u.(type) {
	case float32, float64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]-reflect.ValueOf(u).Float())
		}
		return
	case int, int8, int16, int32, int64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]-float64(reflect.ValueOf(u).Int()))
		}
		return
	case uint, uint8, uint16, uint32, uint64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]-float64(reflect.ValueOf(u).Uint()))
		}
		return
	case Vector:
		s := reflect.ValueOf(u)
		if len(*v) != s.Len() {
			return *v, errors.New("Error: Vectors are not of the same size")
		}
		for i := range *v {
			newVector = append(newVector, (*v)[i]-s.Index(i).Float())
		}
		return
	default:
		return *v, fmt.Errorf("Error: Can't use %v to manipulate a Vector", reflect.TypeOf(u).String())
	}
}

func (v *Vector) Multiply(u interface{}) (newVector Vector, err error) {
	switch u.(type) {
	case float32, float64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]*reflect.ValueOf(u).Float())
		}
		return
	case int, int8, int16, int32, int64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]*float64(reflect.ValueOf(u).Int()))
		}
		return
	case uint, uint8, uint16, uint32, uint64:
		for i := range *v {
			newVector = append(newVector, (*v)[i]*float64(reflect.ValueOf(u).Uint()))
		}
		return
	default:
		return *v, fmt.Errorf("Error: Can't use %v to manipulate a Vector", reflect.TypeOf(u).String())
	}
}

func (v *Vector) Divide(u interface{}) (newVector Vector, err error) {
	switch u.(type) {
	case float32, float64:
		if reflect.ValueOf(u).Float() == 0 {
			return *v, errors.New("Error: Cannot divide by zero")
		}
		for i := range *v {
			newVector = append(newVector, (*v)[i]/reflect.ValueOf(u).Float())
		}
		return
	case int, int8, int16, int32, int64:
		if reflect.ValueOf(u).Int() == 0 {
			return *v, errors.New("Error: Cannot divide by zero")
		}
		for i := range *v {
			newVector = append(newVector, (*v)[i]/float64(reflect.ValueOf(u).Int()))
		}
		return
	case uint, uint8, uint16, uint32, uint64:
		if reflect.ValueOf(u).Uint() == 0 {
			return *v, errors.New("Error: Cannot divide by zero")
		}
		for i := range *v {
			newVector = append(newVector, (*v)[i]/float64(reflect.ValueOf(u).Uint()))
		}
		return
	default:
		return *v, fmt.Errorf("Error: Can't use a %v to manipulate a Vector", reflect.TypeOf(u).String())
	}
	return
}

func (v *Vector) Len() float64 {
	var summed []float64
	for i := 0; i < len(*v); i++ {
		summed = append(summed, math.Pow((*v)[i], 2))
	}
	return math.Sqrt(sum(summed))
}

func (v *Vector) Normalize() (newVector Vector, err error) {
	newVector, err = (*v).Divide((*v).Len())
	return
}
