package myutils

import "reflect"

func IsDeepEqual2dSlice(aa, bb interface{}) bool {
	switch reflect.TypeOf(bb).String() {
	case "[][]string":
		return isEqualStrSlices(aa.([][]string), bb.([][]string))
	case "[][]int":
		return isEqualIntSlices(aa.([][]int), bb.([][]int))
	}
	return true
}

func isEqualIntSlices(aa, bb [][]int) bool {
	eqCtr := 0
	for _, a := range aa {
		for _, b := range bb {
			if reflect.DeepEqual(a, b) {
				eqCtr++
			}
		}
	}
	if eqCtr != len(bb) || len(aa) != len(bb) {
		return false
	}
	return true
}

func isEqualStrSlices(aa, bb [][]string) bool {
	eqCtr := 0
	for _, a := range aa {
		for _, b := range bb {
			if reflect.DeepEqual(a, b) {
				eqCtr++
			}
		}
	}
	if eqCtr != len(bb) || len(aa) != len(bb) {
		return false
	}
	return true
}
