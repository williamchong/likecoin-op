package commaseparatedstring

import "strings"

type CommaSeparatedString string

func (s CommaSeparatedString) ToSlice() []string {
	return strings.Split(string(s), ",")
}

func FromSlice(slice []string) CommaSeparatedString {
	return CommaSeparatedString(strings.Join(slice, ","))
}
