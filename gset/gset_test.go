package gset

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	gset GSet
)

func init() {
	gset = Initialize()
}

// TestList checks the basic functionality of GSet List()
// List() should return all unique values appended to the GSet
func TestList(t *testing.T) {
	gset.Set, _ = gset.Append("xx")

	expectedValue := []string{"xx"}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestList_UpdatedValue checks the functionality of GSet List() when 
// multiple values are appended to GSet it should return 
// all the unique values appended to the GSet
func TestList_UpdatedValue(t *testing.T) {
	gset.Set, _ = gset.Append("xx")
	gset.Set, _ = gset.Append("yy")
	gset.Set, _ = gset.Append("zz")

	expectedValue := []string{"xx", "yy", "zz"}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestList_NoValue checks the functionality of GSet List() when 
// no values are appended to GSet, it should return 
// an empty string slice when the GSet is empty
func TestList_NoValue(t *testing.T) {
	expectedValue := []string{}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestAppend checks the basic functionality of GSet Append() 
// it should return the GSet back when the append is successfull
func TestAppend(t *testing.T) {
	expectedValue := []string{"xx"}
	actualValue, actualError := gset.Append("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestAppend_NoValue checks the functionality of GSet Append() 
// when a nil value is passed to it, it should return 
// the an empty string slice back along with an error
func TestAppend_NoValue(t *testing.T) {
	expectedValue := []string{}
	expectedError := errors.New("empty value provided")
	actualValue, actualError := gset.Append("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestAppend_Duplicate checks the functionality of GSet Append() 
// when a duplicate value is passed to it, it should return 
// only the unique GSet values
func TestAppend_Duplicate(t *testing.T) {
	gset.Set, _ = gset.Append("xx")
	gset.Set, _ = gset.Append("yy")
	gset.Set, _ = gset.Append("xx")

	expectedValue := []string{"xx", "yy"}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestClear checks the basic functionality of GSet Clear()
// utility function it clears all the values in a GSet set
func TestClear(t *testing.T) {
	gset.Set, _ = gset.Append("xx1")
	gset.Set, _ = gset.Append("xx2")
	gset.Set = gset.Clear()

	expectedValue := []string{}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestClear_EmptyStore checks the functionality of GSet Clear() utility function
// when no values are in it, it clears all the values in a GSet set
func TestClear_EmptyStore(t *testing.T) {
	gset.Set = gset.Clear()

	expectedValue := []string{}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestLookup checks the basic functionality of GSet Lookup() function
// it returns a boolean if a value passed is present in the GSet set or not 
func TestLookup(t *testing.T) {
	gset.Set, _ = gset.Append("xx")

	expectedValue := true
	actualValue, actualError := gset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestLookup_NotPresent checks the functionality of GSet Lookup() function
// it returns false if a value passed is not present in the GSet
func TestLookup_NotPresent(t *testing.T) {
	gset.Set, _ = gset.Append("xx")

	expectedValue := false
	actualValue, actualError := gset.Lookup("yy")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestLookup_EmptySet checks the functionality of GSet Lookup() function
// it returns false if the GSet is empty irrespective of the value passed 
func TestLookup_EmptySet(t *testing.T) {
	expectedValue := false
	actualValue, actualError := gset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestLookup_EmptyLookup checks the functionality of GSet Lookup() function
// it returns an error if the value passed is nil irrespective of the GSet
func TestLookup_EmptyLookup(t *testing.T) {
	expectedValue := false
	expectedError := errors.New("empty value provided")

	actualValue, actualError := gset.Lookup("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestMerge checks the basic functionality of the Merge() function on multiple GSets
// it returns all the GSets merged together with unique elements as one single GSet
func TestMerge(t *testing.T) {
	gset1 := GSet{[]string{"xx"}}
	gset2 := GSet{[]string{"yy"}}
	gset3 := GSet{[]string{"zz"}}

	expectedValue := GSet{[]string{"xx", "yy", "zz"}}
	actualValue, actualError := Merge(gset1, gset2, gset3)

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestMerge_Empty checks the functionality of the Merge() function on multiple GSets
// when one GSet is empty, it returns an empty GSet followed by an error
func TestMerge_Empty(t *testing.T) {
	gset1 := GSet{[]string{"xx"}}
	gset2 := GSet{[]string{""}}
	gset3 := GSet{[]string{"zz"}}

	expectedValue := GSet{}
	expectedError := errors.New("empty value provided")

	actualValue, actualError := Merge(gset1, gset2, gset3)

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

// TestMerge_Duplicate checks the functionality of the Merge() function on multiple GSets
// when duplicate values are passed with the GSet it returns all the GSets
// merged together with unique elements as one single GSet
func TestMerge_Duplicate(t *testing.T) {
	gset1 := GSet{[]string{"xx"}}
	gset2 := GSet{[]string{"zz"}}
	gset3 := GSet{[]string{"zz"}}

	expectedValue := GSet{[]string{"xx", "zz"}}
	actualValue, actualError := Merge(gset1, gset2, gset3)

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}
