// #############################################################################
// # File: compare_test.go                                                     #
// # Project: comparejson                                                      #
// # Created Date: 2023/12/05 10:59:19                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/12/05 11:15:36                                        #
// # Modified By: realjf                                                       #
// # -----                                                                     #
// #                                                                           #
// #############################################################################
package comparejson_test

import (
	"fmt"
	"testing"

	"github.com/realjf/comparejson"
)

func TestCompareMap(t *testing.T) {
	oldMap := map[string]interface{}{
		"a": []interface{}{
			"a", "b", 34, 23.23,
		},
	}
	newMap := map[string]interface{}{
		"a": []interface{}{
			"a", "d", 34, 23.13,
		},
	}
	fmt.Println(comparejson.CompareMap(oldMap, newMap))

	oldMap1 := map[string]interface{}{
		"a": []interface{}{
			"a", "b", 34, 23.23,
			[]string{"c", "d"},
		},
		"map": map[string]interface{}{
			"a": []interface{}{
				"e", "f", 22,
			},
			"b": "test",
			"c": 12,
			"d": "dddd",
		},
	}
	newMap1 := map[string]interface{}{
		"a": []interface{}{
			"a", "b", 34, 23.13,
			[]string{"c", "f"},
		},
		"map": map[string]interface{}{
			"a": []interface{}{
				"e", "g", 22,
			},
			"b": "test1",
			"c": 12,
			"f": 32,
		},
	}
	fmt.Println(comparejson.CompareMap(oldMap1, newMap1))

}

func TestCompareJson(t *testing.T) {
	oldJson := []byte(`{"a": "test", "b": "test23", "c": { "a": "test", "b": "test"}, "d":[23, 232, 23]}`)
	newJson := []byte(`{"a": "test1", "b": "test23", "c": { "a1": "test1", "f": 131}, "d":[213, 232, 23, 323]}`)
	fmt.Println(comparejson.CompareJson(oldJson, newJson))
}
