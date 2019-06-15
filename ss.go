// package ss provides translation from spreadsheet column references
// to zero-based indexes.
//
// Eg. A => 0, B => 1, Z => 25, AA => 26 etc.
package ss

import "strings"

func CtoI(col string) int {
	col = strings.ToLower(col)
	l := len(col)
	idx := 0
	for i := l - 1; i >= 0; i-- {
		idx = idx + int(col[i]-'a'+1)*pow26(l-i-1)

	}
	return idx - 1
}
func pow26(l int) int {
	m := 1
	if l == 0 {
		return m
	}
	for i := 0; i < l; i++ {
		m = m * 26
	}
	return m
}
