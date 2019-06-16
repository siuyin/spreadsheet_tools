// Package ss provides translation from spreadsheet column references
// to zero-based indexes.
//
// Eg. A => 0, B => 1, Z => 25, AA => 26 etc.
package ss

import "strings"

// CtoI converts spreadsheet column references to zero-based indexes.
// Eg CtoI("A") returns 0 and CtoI("Aa") returns 26. The column reference
// string is case-insensitive.
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
	for i := 0; i < l; i++ {
		m = m * 26
	}
	return m
}
