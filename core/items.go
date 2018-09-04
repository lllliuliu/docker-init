package core

import "fmt"

// Items is alias of map[string]string, represet multiple of key/value
type Items struct {
	maps map[string]string
}

// NewItems build the new items use []map[string]string or nil
func NewItems(ms map[string]string) *Items {
	if ms == nil {
		ms = map[string]string{}
	}

	return &Items{maps: ms}
}

// AddTwoSlice is append slice of item
func (s *Items) AddTwoSlice(key []string, value []string) *Items {
	var v string
	for i, k := range key {
		if len(value) <= i {
			v = ""
		} else {
			v = value[i]
		}
		s.AddTwoString(k, v)
	}

	return s
}

// AddTwoString is append item
func (s *Items) AddTwoString(key string, value string) *Items {
	s.maps[key] = value
	return s
}

// ToString formats the output string using the
// connection string, one line per item
func (s *Items) ToString(connection string) string {
	var c string
	for k, v := range s.maps {
		c = fmt.Sprintf("%s%s%s%s\n", c, k, connection, v)
	}
	return c
}
