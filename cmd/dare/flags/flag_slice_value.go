package flags

import (
	"flag"
	"strings"
)

// AppendSliceValue implements the flag.Value interface and allows multiple
// calls to the same variable to append a list.
type AppendSliceValue []string

var _ flag.Value = new(AppendSliceValue)

// String ...
func (s *AppendSliceValue) String() string {
	return strings.Join(*s, ",")
}

// Set ...
func (s *AppendSliceValue) Set(value string) error {
	if *s == nil {
		*s = make([]string, 0, 1)
	}

	*s = append(*s, value)
	return nil
}
