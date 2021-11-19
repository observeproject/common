package model

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// SchemaNameRE is a regular expression matching valid label names. Note that the
// IsValid method of SchemaName performs the same check but faster than a match
// with this regular expression.
var TypeNameRE = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$")

// A TypeName is a key for a Resource Schema.  It has a value associated therewith.
type TypeName string

// IsValid is true iff the label name matches the pattern of SchemaNameRE. This
// method, however, does not use SchemaNameRE for the check but a much faster
// hardcoded implementation.
func (tn TypeName) IsValid() bool {
	if len(tn) == 0 {
		return false
	}
	for i, b := range tn {
		if !((b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || ((b >= '0' && b <= '9') || b == '_' && i > 0)) {
			return false
		}
	}
	return true
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (tn *TypeName) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	if !TypeName(s).IsValid() {
		return fmt.Errorf("%q is not a valid resource schema name", s)
	}
	*tn = TypeName(s)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (tn *TypeName) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if !TypeName(s).IsValid() {
		return fmt.Errorf("%q is not a valid resource schema name", s)
	}
	*tn = TypeName(s)
	return nil
}
