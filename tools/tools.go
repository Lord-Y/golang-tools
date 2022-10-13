// Package tools assemble useful functions used by other packages
package tools

import "encoding/json"

// StringInSlice find value exist in array
func StringInSlice(val string, inSlice []string) (b bool) {
	for i := range inSlice {
		if val == inSlice[i] {
			return true
		}
	}
	return false
}

// SliceDifference will compare 2 slices and return the difference
func SliceDifference(source, dest []string) (diff []string) {
	mb := make(map[string]struct{}, len(dest))
	for _, x := range dest {
		mb[x] = struct{}{}
	}
	for _, x := range source {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return
}

// ConvertMapStringInterfaceToMapStringString will convert map[string]interface{} into map[string]string
func ConvertMapStringInterfaceToMapStringString(m map[string]interface{}) (z map[string]string, err error) {
	j, err := json.Marshal(m)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, &z)
	if err != nil {
		return
	}
	return
}
