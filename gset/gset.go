package gset

import (
	"errors"
)

// package gset implements the GSet CRDT data type along with the functionality to 
// append, list & lookup values in a GSet. It also provides the functionality to 
// merge multiple GSets together and a utility function to clear a GSet used in tests

// GSet is the GSet CRDT data type
type GSet struct {
	// Set contains all the string 
	// values stored in the GSet
	Set []string `json:"set"`
}

// Initialize returns a new empty GSet
func Initialize() GSet {
	return GSet{Set: make([]string, 0)}
}

// Append adds a new unique value to the GSet using the 
// union operation for each value on the existing GSet
func (gset GSet) Append(value string) ([]string, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return []string{}, errors.New("empty value provided")
	}

	// Set = Set U value
	gset.Set = Union(gset.Set, value)
	
	// Return the new GSet followed by nil error
	return gset.Set, nil
}

// Lookup returns either boolean true/false indicating
// if a given value is present in the GSet or not 
func (gset GSet) Lookup(value string) (bool, error) {
	// Return an error if the value passed is nil
	if value == "" {
		return false, errors.New("empty value provided")
	}

	// Iterative over the GSet and check if the 
	// value is the one we're seraching
	// return true if the value exists
	for _, element := range gset.Set {
		if element == value {
			return true, nil
		}
	}

	// If the value isn't found after iterating 
	// over the entire GSet we return false
	return false, nil
}

// List returns all the elements present in the GSet
func (gset GSet) List() []string {
	return gset.Set
}

// Merge conbines multiple GSets together using Union 
// and returns a single merged GSet
func Merge(GSets ...GSet) (GSet, error) {
	var gsetMerged GSet
	var err error

	// GSetMerged = GSetMerged U GSetToMergeWith
	for _, gset := range GSets {
		for _, value := range gset.Set {
			gsetMerged.Set, err = gsetMerged.Append(value)
			if err != nil {
				return GSet{}, err
			}
		}
	}

	
	// Return the merged GSet followed by nil error
	return gsetMerged, nil
}

// Clear is utility function used only for tests 
// to empty the contents of a given GSet
func (gset GSet) Clear() []string {
	gset.Set = []string{}
	return gset.Set
}

// Union performs the Union operation for a set and a value
func Union(set []string, value string) []string {
	for _, element := range set {
		if element == value {
			return set
		}
	}
	return append(set, value)
}
