package model

import (
	"encoding/json"
	"fmt"
	"regexp"
)

// SchemaNameRE is a regular expression matching valid label names. Note that the
// IsValid method of SchemaName performs the same check but faster than a match
// with this regular expression.
var SchemaNameRE = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$")

// A SchemaName is a key for a Resource Schema.  It has a value associated therewith.
type SchemaName string

// IsValid is true iff the label name matches the pattern of SchemaNameRE. This
// method, however, does not use SchemaNameRE for the check but a much faster
// hardcoded implementation.
func (sn SchemaName) IsValid() bool {
	if len(sn) == 0 {
		return false
	}
	for i, b := range sn {
		if !((b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || ((b >= '0' && b <= '9') || b == '_' && i > 0)) {
			return false
		}
	}
	return true
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (sn *SchemaName) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}
	if !SchemaName(s).IsValid() {
		return fmt.Errorf("%q is not a valid resource schema name", s)
	}
	*sn = SchemaName(s)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (sn *SchemaName) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if !SchemaName(s).IsValid() {
		return fmt.Errorf("%q is not a valid resource schema name", s)
	}
	*sn = SchemaName(s)
	return nil
}
