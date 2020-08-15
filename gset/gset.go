package gset

import (
	"errors"
)

// package gset ...

// GSet ...
type GSet struct {
	Set []string
}

// Initialize ...
func Initialize() GSet {
	return GSet{Set: make([]string, 0)}
}

// Append ...
func (gset GSet) Append(value string) ([]string, error) {
	if value == "" {
		return []string{}, errors.New("empty value provided")
	}

	gset.Set = Union(gset.Set, value)
	return gset.Set, nil
}

// Lookup ...
func (gset GSet) Lookup(value string) (bool, error) {
	if value == "" {
		return false, errors.New("empty value provided")
	}

	for _, element := range gset.Set {
		if element == value {
			return true, nil
		}
	}

	return false, nil
}

// Merge ...
func Merge(GSets ...GSet) (GSet, error) {
	var gsetMerged GSet
	var err error

	for _, gset := range GSets {
		for _, value := range gset.Set {
			gsetMerged.Set, err = gsetMerged.Append(value)
			if err != nil {
				return GSet{}, err
			}
		}
	}

	return gsetMerged, nil
}

// List ...
func (gset GSet) List() []string {
	return gset.Set
}

// Clear ...
// TODO: Indicate that it is utility
// func used only for tests
func (gset GSet) Clear() []string {
	gset.Set = []string{}
	return gset.Set
}

// Union ...
func Union(set []string, value string) []string {
	for _, element := range set {
		if element == value {
			return set
		}
	}
	return append(set, value)
}
