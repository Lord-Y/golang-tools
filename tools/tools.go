// Package tools assemble useful functions used by other packages
package tools

// StringInSlice find value exist in array
func StringInSlice(val string, inSlice []string) (b bool) {
	for i := range inSlice {
		if val == inSlice[i] {
			return true
		}
	}
	return false
}
