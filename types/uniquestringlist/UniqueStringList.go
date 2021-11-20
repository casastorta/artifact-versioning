package uniquestringlist

import (
	"fmt"
	"strings"
)

type UniqueStringList []string

// Contains Return true if UniqueStringList contains (any of) the item(s)
func (d *UniqueStringList) Contains(elems ...string) bool {
	for _, l := range *d {
		for _, e := range elems {
			if l == e {
				return true
			}
		}
	}
	return false
}

// Append item(s) to the UniqueStringList
func (d *UniqueStringList) Append(elems ...string) error {
	for _, e := range elems {
		if d.Contains(e) {
			return fmt.Errorf("cannot add existing item: '%s'", e)
		}
	}
	*d = append(*d, elems...)
	return nil
}

// FromString Parse UniqueStringList elements from the string of values separated by the `sep` separator
func (d *UniqueStringList) FromString(s, sep string) error {
	var (
		ss = strings.Split(s, sep)
		us UniqueStringList
	)
	for _, ts := range ss {
		err := us.Append(strings.TrimSpace(ts))
		if err != nil {
			return err
		}
	}
	*d = us
	return nil
}
