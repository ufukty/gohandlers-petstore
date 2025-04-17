package types

import (
	"fmt"
	"regexp"
)

// this is just to save space in types.go file
type reggy struct {
	Compiled *regexp.Regexp
}

func (r reggy) Validate(s string) error {
	if !r.Compiled.MatchString(string(s)) {
		return fmt.Errorf("invalid format")
	}
	return nil
}
