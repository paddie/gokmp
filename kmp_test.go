package kmp

import (
	// "fmt"
    "testing"
    "strings"
)

const pattern = "aabb"
const str = "aabaabaaaabbaabaabaaabbaabaabb"
const bad_str = "pepper"

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