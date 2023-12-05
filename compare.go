// #############################################################################
// # File: compare.go                                                          #
// # Project: comparejson                                                      #
// # Created Date: 2023/12/05 10:58:01                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/12/05 11:31:03                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package comparejson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func CompareJson(oldJson, newJson []byte) (changed []string) {
	changed = make([]string, 0)
	var oldMap, newMap map[string]interface{}
	err := json.Unmarshal(oldJson, &oldMap)
	if err != nil {
		return
	}
	err = json.Unmarshal(newJson, &newMap)
	if err != nil {
		return
	}

	return CompareMap(oldMap, newMap)
}

func CompareMap(oldMap, newMap map[string]interface{}) (changed []string) {
	changed = make([]string, 0)
	changedMap := make(map[string]string, 0)
	for key, val := range oldMap {
		if val2, ok := newMap[key]; ok {
			// 对比是否相等
			equal, kk := compareJson(val, val2)
			if !equal {
				if len(kk) == 0 {
					changedMap[key] = key
				} else {
					for _, ks := range kk {
						newKey := fmt.Sprintf("%s.%s", key, ks)
						changedMap[newKey] = newKey
					}
				}
			}
		} else {
			// 不存在
			fmt.Printf("不存在：%s", key)
			changedMap[key] = key
		}
	}

	for key, val := range newMap {
		if val2, ok := oldMap[key]; ok {
			// 对比是否相等
			equal, kk := compareJson(val, val2)
			if !equal {
				if len(kk) == 0 {
					changedMap[key] = key
				} else {
					for _, ks := range kk {
						newKey := fmt.Sprintf("%s.%s", key, ks)
						changedMap[newKey] = newKey
					}
				}
			}
		} else {
			// 不存在
			changedMap[key] = key
		}
	}

	for _, v := range changedMap {
		changed = append(changed, v)
	}

	return
}

func compareJson(val interface{}, val2 interface{}) (equal bool, keys []string) {
	keys = make([]string, 0)
	valRef := reflect.ValueOf(val)
	valRef2 := reflect.ValueOf(val2)
	if val2 == nil {
		return
	}
	if val == nil {
		return
	}
	switch valRef2.Kind() {
	case reflect.Map, reflect.Struct:
		if valRef.Kind() == valRef2.Kind() {
			isEqual := true
			for _, k := range valRef2.MapKeys() {
				if !valRef.MapIndex(k).IsValid() {
					isEqual = false
					keys = append(keys, k.String())
					continue
				}
				if valRef2.IsValid() && valRef2.MapIndex(k).IsValid() && valRef2.MapIndex(k).CanInterface() {
					// 存在
					v, kk := compareJson(valRef.MapIndex(k).Interface(), valRef2.MapIndex(k).Interface())
					if !v {
						if len(kk) == 0 {
							keys = append(keys, k.String())
						} else {
							for _, ks := range kk {
								keys = append(keys, fmt.Sprintf("%s.%s", k.String(), ks))
							}
						}
						isEqual = false
					}
				} else {
					// 不存在
					keys = append(keys, k.String())
				}
			}
			return isEqual, keys
		} else {
			return
		}
	case reflect.Array, reflect.Slice:
		if valRef.Kind() != valRef2.Kind() {
			return
		}
		equal = true
		for i := 0; i < valRef2.Len(); i++ {
			if i > valRef.Len()-1 {
				equal = false
				keys = append(keys, fmt.Sprintf("[%d]", i))
				continue
			}
			v1 := valRef.Index(i).Interface()
			v2 := valRef2.Index(i).Interface()
			isEqual, kk := compareJson(v1, v2)
			if !isEqual {
				if len(kk) == 0 {
					keys = append(keys, fmt.Sprintf("[%d]", i))
				} else {
					for _, ks := range kk {
						keys = append(keys, fmt.Sprintf("[%d].%s", i, ks))
					}
				}
				equal = false
			}
		}
		return equal, keys
	default:
		if valRef.Kind() != valRef2.Kind() {
			return
		}
		if val == val2 {
			return true, keys
		}
		return
	}
}
