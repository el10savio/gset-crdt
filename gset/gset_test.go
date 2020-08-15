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

func TestList(t *testing.T) {
	gset.Set, _ = gset.Append("xx")

	expectedValue := []string{"xx"}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestList_UpdatedValue(t *testing.T) {
	gset.Set, _ = gset.Append("xx")
	gset.Set, _ = gset.Append("yy")
	gset.Set, _ = gset.Append("zz")

	expectedValue := []string{"xx", "yy", "zz"}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestList_NoValue(t *testing.T) {
	expectedValue := []string{}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestAppend(t *testing.T) {
	expectedValue := []string{"xx"}
	actualValue, actualError := gset.Append("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestAppend_NoValue(t *testing.T) {
	expectedValue := []string{}
	expectedError := errors.New("empty value provided")
	actualValue, actualError := gset.Append("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestAppend_Duplicate(t *testing.T) {
	gset.Set, _ = gset.Append("xx")
	gset.Set, _ = gset.Append("yy")
	gset.Set, _ = gset.Append("xx")

	expectedValue := []string{"xx", "yy"}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestClear(t *testing.T) {
	gset.Set, _ = gset.Append("xx1")
	gset.Set, _ = gset.Append("xx2")
	gset.Set = gset.Clear()

	expectedValue := []string{}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestClear_EmptyStore(t *testing.T) {
	gset.Set = gset.Clear()

	expectedValue := []string{}
	actualValue := gset.List()

	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestLookup(t *testing.T) {
	gset.Set, _ = gset.Append("xx")

	expectedValue := true
	actualValue, actualError := gset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestLookup_NotPresent(t *testing.T) {
	gset.Set, _ = gset.Append("xx")

	expectedValue := false
	actualValue, actualError := gset.Lookup("yy")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestLookup_EmptySet(t *testing.T) {
	expectedValue := false
	actualValue, actualError := gset.Lookup("xx")

	assert.Nil(t, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

func TestLookup_EmptyLookup(t *testing.T) {
	expectedValue := false
	expectedError := errors.New("empty value provided")

	actualValue, actualError := gset.Lookup("")

	assert.Equal(t, expectedError, actualError)
	assert.Equal(t, expectedValue, actualValue)

	gset.Set = gset.Clear()
}

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
