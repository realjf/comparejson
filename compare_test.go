// #############################################################################
// # File: compare_test.go                                                     #
// # Project: comparejson                                                      #
// # Created Date: 2023/12/05 10:59:19                                         #
// # Author: realjf                                                            #
// # -----                                                                     #
// # Last Modified: 2023/12/05 11:00:04                                        #
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

func TestCompareJson(t *testing.T) {
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
	fmt.Println(comparejson.CompareJson(oldMap, newMap))

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
	fmt.Println(comparejson.CompareJson(oldMap1, newMap1))

}
