package kmp

import (
    "testing"
    "strings"
)

// TESTS

// pretty much the worst case string for strings.Index
// wrt. string and pattern
const str = "aabaabaaaabbaabaabaaabbaabaabb"
const pattern = "aabb"

func TestFindAllStringIndex(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	// fmt.Println(kmp)
	ints := kmp.FindAllStringIndex(str)
	test := []int{8,19,26}
	for i, v := range ints {
		if test[i] != v {
			t.Errorf("FindAllStringIndex:\t%v != %v, (exp: %v != act: %v)", test[i], v, ints, test)
		}
	}
}

func TestFindStringIndex(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	ints := kmp.FindStringIndex(str)
	test := 8
	if ints != test {
		t.Errorf("FindStringIndex:\t%v != %v",ints, test)
	}
}

func TestContainedIn(t * testing.T) {
	kmp, _ := NewKMP(pattern)
	if !kmp.ContainedIn(str) {
		t.Errorf("ContainedIn:\tExpected: True !=  actual: False")
	}	
}

func TestOccurrences(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	nr := kmp.Occurrences(str)
	if nr != 3 {
		t.Errorf("Occurences:\texp: %v != act: %v)", 3, nr)
	}
}

func TestOccurrencesFail(t *testing.T) {
	kmp, _ := NewKMP(pattern)
	nr := kmp.Occurrences("pebble")
	if nr != 0 {
		t.Errorf("Occurences:\texp: %v != act: %v)", 0, nr)
	}
}

// BENCHMARKS

func BenchmarkKMPIndexComparison(b *testing.B ) {
    kmp,_ := NewKMP(pattern)
    for i := 0; i < b.N; i++ {        
        _ = kmp.FindStringIndex(str)
    }
}

func BenchmarkStringsIndexComparison(b *testing.B ) {
    for i := 0; i < b.N; i++ {
        _ = strings.Index(str, pattern)
    }
}