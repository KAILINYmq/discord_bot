package helper

import (
	"fmt"
	"reflect"
)

// 并集
func SliceUnion(slice1, slice2 interface{}) []interface{} {
	s1 := toSLice(slice1)
	s2 := toSLice(slice2)
	if s1 == nil || s2 == nil {
		return nil
	}
	m := make(map[string]int)

	for _, v := range s1 {
		k := fmt.Sprintf("m_%v", v)
		m[k]++
	}

	for _, v := range s2 {
		k := fmt.Sprintf("m_%v", v)
		times, _ := m[k]
		if times == 0 {
			s1 = append(s1, v)
		}
	}
	return s1
}

// 交集
func SliceIntersect(slice1, slice2 interface{}) []interface{} {
	s1 := toSLice(slice1)
	s2 := toSLice(slice2)
	if s1 == nil || s2 == nil {
		return nil
	}
	m := make(map[string]int)
	result := make([]interface{}, 0)

	for _, v := range s1 {
		k := fmt.Sprintf("m_%v", v)
		m[k]++
	}

	for _, v := range s2 {
		k := fmt.Sprintf("m_%v", v)
		times, _ := m[k]
		if times == 1 {
			result = append(result, v)
		}
	}
	return result
}

// 差集
func SliceDifference(slice1, slice2 interface{}) []interface{} {
	s1 := toSLice(slice1)
	s2 := toSLice(slice2)
	if s1 == nil || s2 == nil {
		return nil
	}

	m := make(map[string]int)

	result := make([]interface{}, 0)

	inter := SliceIntersect(slice1, slice2)

	for _, v := range inter {
		k := fmt.Sprintf("m_%v", v)
		m[k]++
	}

	for _, v := range s1 {
		k := fmt.Sprintf("m_%v", v)
		times, _ := m[k]
		if times == 0 {
			result = append(result, v)
		}
	}
	return result
}

func toSLice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return nil
	}

	l := v.Len()
	result := make([]interface{}, l)

	for i := 0; i < l; i++ {
		result[i] = v.Index(i).Interface()
	}
	return result
}
