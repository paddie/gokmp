// String-matching in Golang using the Knuth–Morris–Pratt algorithm (KMP)
package gokmp

import (
	"errors"
	"fmt"
)

type KMP struct {
	pattern string
	prefix  []int
	size    int
}

// For debugging
func (kmp *KMP) String() string {
	return fmt.Sprintf("pattern: %v\nprefix: %v", kmp.pattern, kmp.prefix)
}

// compile new prefix-array given argument
func NewKMP(pattern string) (*KMP, error) {
	prefix, err := computePrefix(pattern)
	if err != nil {
		return nil, err
	}
	return &KMP{
			pattern: pattern,
			prefix:  prefix,
			size:    len(pattern)},
		nil
}

// returns an array containing indexes of matches
// - error if pattern argument is less than 1 char 
func computePrefix(pattern string) ([]int, error) {
	// sanity check
	len_p := len(pattern)
	if len_p < 2 {
		if len_p == 0 {
			return nil, errors.New("'pattern' must contain at least one character")
		}
		return []int{-1}, nil
	}
	t := make([]int, len_p)
	t[0], t[1] = -1, 0

	pos, count := 2, 0
	for pos < len_p {

		if pattern[pos-1] == pattern[count] {
			count++
			t[pos] = count
			pos++
		} else {
			if count > 0 {
				count = t[count]
			} else {
				t[pos] = 0
				pos++
			}
		}
	}
	return t, nil
}

// return index of first occurence of kmp.pattern in argument 's'
// - if not found, returns -1
func (kmp *KMP) FindStringIndex(s string) int {
	// sanity check
	if len(s) < kmp.size {
		return -1
	}
	m, i := 0, 0
	for m+i < len(s) {
		if kmp.pattern[i] == s[m+i] {
			if i == kmp.size-1 {
				return m
			}
			i++
		} else {
			m = m + i - kmp.prefix[i]
			if kmp.prefix[i] > -1 {
				i = kmp.prefix[i]
			} else {
				i = 0
			}
		}
	}
	return -1
}

// returns true if pattern i matched at least once
func (kmp *KMP) ContainedIn(s string) bool {
	return kmp.FindStringIndex(s) >= 0
}

// returns the number of occurences of pattern in argument
func (kmp *KMP) Occurrences(s string) int {
	return len(kmp.FindAllStringIndex(s))
}

// for effeciency, define default array-size
const startSize = 10

// find every occurence of the kmp.pattern in 's'
func (kmp *KMP) FindAllStringIndex(s string) []int {
	// precompute
	len_s := len(s)

	if len_s < kmp.size {
		return []int{}
	}

	match := make([]int, 0, startSize)
	m, i := 0, 0
	for m+i < len_s {
		if kmp.pattern[i] == s[m+i] {
			if i == kmp.size-1 {
				// the word was matched
				match = append(match, m)
				// simulate miss, and keep running
				m = m + i - kmp.prefix[i]
				if kmp.prefix[i] > -1 {
					i = kmp.prefix[i]
				} else {
					i = 0
				}
			} else {
				i++
			}
		} else {
			m = m + i - kmp.prefix[i]
			if kmp.prefix[i] > -1 {
				i = kmp.prefix[i]
			} else {
				i = 0
			}
		}
	}
	return match
}
